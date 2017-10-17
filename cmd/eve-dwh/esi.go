package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/PuerkitoBio/rehttp"
	log "github.com/Sirupsen/logrus"
	"github.com/antihax/goesi"
	"github.com/curzonj/eve-dwh-golang/types"
	"github.com/curzonj/eve-dwh-golang/utils/dbcache"
	"github.com/gregjones/httpcache"
	"github.com/rubyist/circuitbreaker"
)

type loggingTransport struct {
	Transport http.RoundTripper
}

type ESIError struct {
	Error    string      `json:"error"`
	Response interface{} `json:"response"`
}

func extractErrorMessage(res *http.Response) (string, error) {
	bodyBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	var bodyStruct ESIError
	if err = json.Unmarshal(bodyBytes, &bodyStruct); err != nil {
		return "", err
	}

	return bodyStruct.Error, nil
}

func (t *loggingTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	start := time.Now()
	res, err := t.Transport.RoundTrip(req)

	finished := time.Now()

	l, ok := req.Context().Value(types.ContextLoggerKey).(log.FieldLogger)
	if !ok {
		l = clients.Logger
	}

	l = l.WithFields(log.Fields{
		"at":   "httpClient",
		"host": req.URL.Hostname(),
		"path": req.URL.RequestURI(),
	})

	var status int

	if res != nil {
		s := res.Header.Get("X-Esi-Error-Limit-Remain")
		if s != "" && s != "100" {
			l = l.WithField("errorsRemaining", s)
		}

		status = res.StatusCode
		if status > 399 {
			msg, err := extractErrorMessage(res)
			if err != nil {
				l.Error(err)
			} else {
				l = l.WithField("error_message", msg)
			}
		}
	}

	l = l.WithFields(log.Fields{
		"elapsed": finished.Sub(start).Seconds(),
		"status":  status,
	})

	if err != nil {
		l.Error(err)
	} else {
		l.Info()
	}

	return res, err
}

// This is used for manual testing of error handling by
// using instead of the http default transport.
type errorTransport struct{}

func (t *errorTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	return nil, errors.New("injected error")
}

type staleTransport struct {
	Transport http.RoundTripper
}

func (t *staleTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Add("Cache-Control", "stale-if-error=86400") // 24 hrs
	return t.Transport.RoundTrip(req)
}

type breakerTransport struct {
	Transport http.RoundTripper

	rate       float64
	minSamples int64
	timeout    time.Duration
	panel      *circuit.Panel
}

func (t *breakerTransport) breakerLookup(url *url.URL) *circuit.Breaker {
	cb, ok := t.panel.Get(url.Host)
	if !ok {
		cb = circuit.NewBreakerWithOptions(&circuit.Options{
			ShouldTrip: circuit.RateTripFunc(t.rate, t.minSamples),
		})
		t.panel.Add(url.Host, cb)
	}

	return cb
}

func (t *breakerTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	var resp *http.Response
	var err error

	breaker := t.breakerLookup(req.URL)
	err = breaker.Call(func() error {
		resp, err = t.Transport.RoundTrip(req)
		return err
	}, t.timeout)
	return resp, err
}

func buildESIClient() {
	cache := dbcache.New(clients.DB, clients.Logger)

	// Add retries, backoff and logging in the transport
	rt := rehttp.NewTransport(
		&loggingTransport{
			Transport: http.DefaultTransport,
		},
		rehttp.RetryAll(
			rehttp.RetryMaxRetries(6),
			rehttp.RetryHTTPMethods("GET"),
			rehttp.RetryAny(
				rehttp.RetryTemporaryErr(),
				rehttp.RetryStatusInterval(500, 600),
			),
		),
		rehttp.ExpJitterDelay(time.Second, time.Minute),
	)

	retriesClient := &http.Client{
		Transport: &httpcache.Transport{
			Transport:           rt,
			Cache:               cache,
			MarkCachedResponses: true,
		},
	}

	scopes := strings.Split(cfg.ESI.ScopesString, " ")

	clients.ESIAuthenticator = goesi.NewSSOAuthenticator(
		retriesClient,
		cfg.ESI.OauthClientID,
		cfg.ESI.OauthSecretKey,
		cfg.ESI.OauthRedirectURL,
		scopes)

	clients.EVERetryClient = goesi.NewAPIClient(retriesClient, cfg.ESI.UserAgent)

	breakerClient := &http.Client{
		Transport: &staleTransport{
			Transport: &httpcache.Transport{
				Transport: &loggingTransport{
					Transport: &breakerTransport{
						Transport: http.DefaultTransport,
						//Transport:  &errorTransport{},
						rate:       0.9,
						minSamples: 10,
						timeout:    time.Second * 10,
						panel:      circuit.NewPanel(),
					},
				},
				Cache:               cache,
				MarkCachedResponses: true,
			},
		},
	}

	clients.HTTPBreakerClient = breakerClient
	clients.EVEBreakerClient = goesi.NewAPIClient(breakerClient, cfg.ESI.UserAgent)
}

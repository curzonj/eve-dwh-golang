package main

import (
	"net/http"
	"strings"
	"time"

	"github.com/PuerkitoBio/rehttp"
	log "github.com/Sirupsen/logrus"
	"github.com/antihax/goesi"
	"github.com/curzonj/eve-dwh-golang/poller"
	"github.com/curzonj/eve-dwh-golang/types"
	"github.com/curzonj/eve-dwh-golang/utils/dbcache"
	"github.com/curzonj/eve-dwh-golang/web"
	"github.com/gregjones/httpcache"
	"github.com/jmoiron/sqlx"
	"github.com/joeshaw/envdecode"
	_ "github.com/lib/pq"
)

var cfg struct {
	Poller poller.Cfg
	Web    web.Cfg

	DatabaseURL string `env:"DATABASE_URL,required"`

	ESI struct {
		UserAgent        string `env:"USER_AGENT,required"`
		OauthClientID    string `env:"OAUTH_CLIENT_ID,required"`
		OauthSecretKey   string `env:"OAUTH_SECRET_KEY,required"`
		OauthRedirectURL string `env:"OAUTH_REDIRECT_URL,required"`
		ScopesString     string `env:"ESI_SCOPES,required"`
	}
}

var clients types.Clients

func connectToDatabase() {
	db, err := sqlx.Connect("postgres", cfg.DatabaseURL)
	if err != nil {
		// TODO could add something like https://github.com/heroku/x/blob/master/hredis/redigo/redigo.go#L15
		log.Fatal(err)
	}

	clients.DB = db
}

type loggingTransport struct{}

func (t *loggingTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	start := time.Now()
	client := http.DefaultClient
	res, err := client.Do(req)

	finished := time.Now()

	l, ok := req.Context().Value(types.ContextLoggerKey).(log.FieldLogger)
	if !ok {
		l = clients.Logger
	}

	s := res.Header.Get("X-Esi-Error-Limit-Remain")
	if s != "" && s != "100" {
		l = l.WithField("errorsRemaining", s)
	}

	l.WithFields(log.Fields{
		"at":      "httpClient",
		"elapsed": finished.Sub(start).Seconds(),
		"status":  res.StatusCode,
		"host":    req.URL.Hostname(),
		"path":    req.URL.RequestURI(),
	}).Info()

	return res, err
}

func buildESIClient() {
	// Add retries, backoff and logging in the transport
	rt := rehttp.NewTransport(
		&loggingTransport{},
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

	httpClient := &http.Client{
		Transport: &httpcache.Transport{
			Transport:           rt,
			Cache:               dbcache.New(clients.DB, clients.Logger),
			MarkCachedResponses: true,
		},
	}

	scopes := strings.Split(cfg.ESI.ScopesString, " ")

	clients.ESIAuthenticator = goesi.NewSSOAuthenticator(
		httpClient,
		cfg.ESI.OauthClientID,
		cfg.ESI.OauthSecretKey,
		cfg.ESI.OauthRedirectURL,
		scopes)

	clients.ESIClient = goesi.NewAPIClient(httpClient, cfg.ESI.UserAgent).ESI
}

func buildClients() {
	connectToDatabase()
	buildESIClient()
}

func loadEnvironment() {
	err := envdecode.Decode(&cfg)
	if err != nil {
		log.Fatal(err)
	}

	clients.Logger = log.New()
}

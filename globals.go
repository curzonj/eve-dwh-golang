package main

import (
	"net/http"
	"time"

	"github.com/PuerkitoBio/rehttp"
	log "github.com/Sirupsen/logrus"
	"github.com/antihax/goesi"
	"github.com/antihax/goesi/esi"
	"github.com/gregjones/httpcache"
	"github.com/gregjones/httpcache/diskcache"
	"github.com/jmoiron/sqlx"
	"github.com/joeshaw/envdecode"
	_ "github.com/lib/pq"
)

var cfg struct {
	DatabaseURL  string `env:"DATABASE_URL,required"`
	UserAgent    string `env:"USER_AGENT,required"`
	RegionID     int32  `env:"REGION_ID,default=10000002"`
	RetryLimit   int32  `env:"RETRY_LIMIT,default=10"`
	MarketGroups []int  `env:"MARKET_GROUPS,required"`
}

var globals struct {
	esiClient *esi.APIClient
	db        *sqlx.DB
	logger    log.FieldLogger
}

func connectToDatabase() {
	db, err := sqlx.Connect("postgres", cfg.DatabaseURL)
	if err != nil {
		log.Fatal(err)
	}

	globals.db = db
}

type loggingTransport struct{}

func (t *loggingTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	start := time.Now()
	client := http.DefaultClient
	res, err := client.Do(req)

	finished := time.Now()

	l := globals.logger
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

func buildEsiClient() {
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
			Cache:               diskcache.New("./cache"),
			MarkCachedResponses: true,
		},
	}

	globals.esiClient = goesi.NewAPIClient(httpClient, cfg.UserAgent).ESI
}

func loadEnvironment() {
	err := envdecode.Decode(&cfg)
	if err != nil {
		log.Fatal(err)
	}

	globals.logger = log.New()
	globals.logger.WithFields(log.Fields{
		"at": "app-boot",
	}).Info()
}

package main

import (
	"net/http"

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

func buildEsiClient() {
	// Add retries, backoff and logging in the transport
	transport := httpcache.NewTransport(diskcache.New("./cache"))
	httpClient := &http.Client{Transport: transport}

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

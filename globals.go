package main

import (
	"log"
	"net/http"

	"database/sql"

	"github.com/Sirupsen/logrus"
	"github.com/antihax/goesi"
	"github.com/antihax/goesi/esi"
	"github.com/gregjones/httpcache"
	"github.com/gregjones/httpcache/diskcache"
	"github.com/joeshaw/envdecode"
	_ "github.com/lib/pq"
)

var cfg struct {
	DatabaseURL string `env:"DATABASE_URL,required"`
	UserAgent   string `env:"USER_AGENT,required"`
	RegionID    int32  `env:"REGION_ID,default=10000002"`
	RetryLimit  int32  `env:"RETRY_LIMIT,default=10"`
}

var globals struct {
	esiClient *esi.APIClient
	db        *sql.DB
	logger    logrus.FieldLogger
}

func connectToDatabase() {
	db, err := sql.Open("postgres", cfg.DatabaseURL)
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

	globals.logger = logrus.New()
}

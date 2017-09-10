package main

import (
	"log"
	"net/http"

	"github.com/antihax/goesi"
	"github.com/antihax/goesi/esi"
	"github.com/gregjones/httpcache"
	"github.com/gregjones/httpcache/diskcache"
	"github.com/joeshaw/envdecode"
)

var cfg struct {
	UserAgent  string `env:"USER_AGENT,required"`
	RegionID   int32  `env:"REGION_ID,default=10000002"`
	RetryLimit int32  `env:"RETRY_LIMIT,default=10"`
}

var globals struct {
	esiClient *esi.APIClient
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
}

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/antihax/goesi"
	"github.com/gregjones/httpcache"
	"github.com/gregjones/httpcache/diskcache"
	"github.com/joeshaw/envdecode"
	"gopkg.in/urfave/cli.v1"
)

var cfg struct {
	UserAgent string `env:"USER_AGENT,required"`
}

func cliServerAction(c *cli.Context) error {
	transport := httpcache.NewTransport(diskcache.New("./cache"))
	httpClient := &http.Client{Transport: transport}

	esiClient := goesi.NewAPIClient(httpClient, cfg.UserAgent).ESI
	regionIDs, _, err := esiClient.UniverseApi.GetUniverseRegions(nil)

	fmt.Printf("%+v %+v", err, regionIDs)

	history, _, err := esiClient.MarketApi.GetMarketsRegionIdHistory(10000002, 34, nil)
	fmt.Printf("%+v %+v", err, history)

	return nil
}

func main() {
	err := envdecode.Decode(&cfg)
	if err != nil {
		log.Fatal(err)
	}

	app := cli.NewApp()
	app.Usage = "main entry point for all operations"

	app.Action = cliServerAction
	app.Commands = []cli.Command{}

	app.Run(os.Args)
}

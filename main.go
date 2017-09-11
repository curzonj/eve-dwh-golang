package main

import (
	"os"

	"gopkg.in/urfave/cli.v1"
)

func cliServerAction(c *cli.Context) error {
	connectToDatabase()

	fetcher := &marketDataFetcher{}
	_, err := fetcher.GetOrderDataset(cfg.RegionID)

	return err
}

func main() {
	loadEnvironment()
	buildEsiClient()

	app := cli.NewApp()
	app.Usage = "main entry point for all operations"

	app.Action = cliServerAction
	app.Commands = []cli.Command{}

	app.Run(os.Args)
}

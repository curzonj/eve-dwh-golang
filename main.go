package main

import (
	"os"

	"github.com/pkg/errors"
	"gopkg.in/urfave/cli.v1"
)

func cliServerAction(c *cli.Context) error {
	connectToDatabase()

	fetcher := &marketDataFetcher{}
	data, err := fetcher.GetOrderDataset(cfg.RegionID)
	if err != nil {
		return errors.Wrap(err, "fetching order data")
	}

	return importBulkOrderStats(data)
}

func main() {
	loadEnvironment()
	buildEsiClient()

	app := cli.NewApp()
	app.Usage = "main entry point for all operations"

	app.Action = cliServerAction
	app.Commands = []cli.Command{
		{
			Name:    "error",
			Aliases: []string{"e"},
			Usage:   "fake an error",
			Action: func(c *cli.Context) error {
				return errors.New("test")
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		globals.logger.Fatal(err)
	}
}

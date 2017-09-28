package main

import (
	"os"
	"time"

	"github.com/curzonj/eve-dwh-golang/data"
	"github.com/curzonj/eve-dwh-golang/web"
	"gopkg.in/urfave/cli.v1"
)

func cliServerAction(c *cli.Context) error {
	atBoot()

	go data.MarketStatisticsPoller(clients, 5*time.Minute)

	select {}
}

func developmentAction(c *cli.Context) error {
	atBoot()

	go web.RunHandler(clients, cfg.Port)

	select {}
}

func atBoot() {
	utils.LoadEnvironment()
	utils.BuildESIClient()
	utils.ConnectToDatabase()
}

func main() {
	app := cli.NewApp()
	app.Usage = "main entry point for all operations"

	app.Commands = []cli.Command{
		{
			Name:   "dev",
			Usage:  "run the development process for whatever I'm working on",
			Action: developmentAction,
		},
		{
			Name:   "serve",
			Usage:  "run the server processes",
			Action: cliServerAction,
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		globals.logger.Fatal(err)
	}
}

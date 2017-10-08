package main

import (
	"os"

	"github.com/curzonj/eve-dwh-golang/web"
	"gopkg.in/urfave/cli.v1"
)

func cliServerAction(c *cli.Context) error {
	buildClients()

	//go poller.MarketStatisticsPoller(clients, cfg.Poller)
	go web.RunHandler(clients, cfg.Web)

	select {}
}

func developmentAction(c *cli.Context) error {
	buildClients()

	go web.RunHandler(clients, cfg.Web)

	select {}
}

func main() {
	loadEnvironment()

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
		clients.Logger.Fatal(err)
	}
}

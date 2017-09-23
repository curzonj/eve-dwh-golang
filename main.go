package main

import (
	"os"
	"time"

	"gopkg.in/urfave/cli.v1"
)

func cliServerAction(c *cli.Context) error {
	atBoot()

	go marketStatisticsPoller(5 * time.Minute)

	select {}
}

func developmentAction(c *cli.Context) error {
	atBoot()

	go runWebHandler()

	select {}
}

func atBoot() {
	loadEnvironment()
	buildEsiClient()
	connectToDatabase()
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

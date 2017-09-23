package main

import (
	"os"
	"time"

	"gopkg.in/urfave/cli.v1"
)

func cliServerAction(c *cli.Context) error {
	atBoot()

	go runWebHandler()

	return nil
}

func developmentAction(c *cli.Context) error {
	atBoot()

	marketStatisticsPoller(5 * time.Minute)

	return nil
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

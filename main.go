package main

import (
	"os"

	"gopkg.in/urfave/cli.v1"
)

func cliServerAction(c *cli.Context) error {
	return nil
}

func main() {
	loadEnvironment()
	buildEsiClient()

	app := cli.NewApp()
	app.Usage = "main entry point for all operations"

	app.Action = importAction //cliServerAction
	app.Commands = []cli.Command{
		{
			Name:   "import",
			Usage:  "import order data",
			Action: importAction,
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		globals.logger.Fatal(err)
	}
}

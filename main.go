package main

import (
	"os"

	"gopkg.in/urfave/cli.v1"
)

func cliServerAction(c *cli.Context) error {
	return nil
}

func main() {
	app := cli.NewApp()
	app.Usage = "main entry point for all operations"

	app.Action = cliServerAction
	app.Commands = []cli.Command{}

	app.Run(os.Args)
}

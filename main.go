package main

import (
	"os"
	"strconv"

	"gopkg.in/urfave/cli.v1"
)

func cliServerAction(c *cli.Context) error {
	connectToDatabase()

	for _, id := range cfg.MarketGroups {
		var typeIDs []int
		err := globals.db.Select(&typeIDs, "select \"typeID\" from \"invTypes\" where \"marketGroupID\" in (select market_group_id from market_group_arrays where id_list && '{"+strconv.Itoa(id)+"}')")
		if err != nil {
			return err
		}

		globals.logger.Info(typeIDs)
	}

	return nil
}

func main() {
	loadEnvironment()
	buildEsiClient()

	app := cli.NewApp()
	app.Usage = "main entry point for all operations"

	app.Action = cliServerAction
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

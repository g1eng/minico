package subcommand

import (
	"github.com/urfave/cli/v2"
)

var List = &cli.Command{
	Name:        "list",
	Usage:       "list and filter containers with fuzzy name matching",
	Description: "list and filter containers with fuzzy name matching",
	Action: func(cli *cli.Context) error {
		var filter string
		if cli.NArg() == 0 {
			filter = "."
		} else {
			filter = cli.Args().Get(0)
		}
		return Platform.ContainerList(
			cli.Bool("all"),
			cli.Bool("with-id"),
			filter,
		)
	},
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:    "all",
			Aliases: []string{"a"},
			Usage:   "list all containers",
			Value:   false,
		},
		&cli.BoolFlag{
			Name:    "with-id",
			Aliases: []string{"id"},
			Usage:   "show id with the list",
			Value:   false,
		},
	},
}

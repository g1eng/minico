package subcommand

import (
	"github.com/urfave/cli/v2"
)

var Images = &cli.Command{
	Name:        "images",
	Usage:       "list and filter images with fuzzy name matching",
	Description: "list and filter images with fuzzy name matching",
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:    "with-id",
			Aliases: []string{"id"},
			Value:   false,
		},
	},
	Action: func(cli *cli.Context) error {
		var filter string
		if cli.NArg() == 0 {
			filter = "."
		} else {
			filter = cli.Args().Get(0)
		}
		return Platform.ImageList(cli.Bool("with-id"), filter)
	},
}

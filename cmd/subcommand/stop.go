package subcommand

import (
	"errors"

	"github.com/urfave/cli/v2"
)

var Stop = &cli.Command{
	Name:        "stop",
	Usage:       "stop containers with fuzzy name matching",
	Description: "stop containers with fuzzy name matching",
	Action: func(cli *cli.Context) error {
		if cli.NArg() == 0 {
			return errors.New("no arg")
		} else {
			args := cli.Args().Slice()
			for i := range args {
				if err := Platform.ContainerStop(args[i]); err != nil {
					return err
				}
			}
			return nil
		}
	},
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name: "all",
		},
	},
}

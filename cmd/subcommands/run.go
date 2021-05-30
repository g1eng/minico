package subcommand

import (
	"errors"
	"github.com/urfave/cli/v2"
)

var Run = &cli.Command{
	Name:  "run",
	Usage: "run a container with the specified image",
	Action: func(cli *cli.Context) error {
		if cli.NArg() == 0 {
			//fixme: simple attaching to the conatainer shell without any argument
			return errors.New(cli.App.Name + " run needs image name")
		} else {
			cmd := cli.Args().Slice()[1:]
			if len(cmd) == 0 {
				cmd = []string{"/bin/sh"}
			}
			return Platform.ContainerRun(
				cli.Args().Get(0),
				cli.Bool("persistent"),
				cmd,
			)
		}
	},
	ArgsUsage: "<image> [command] [args...]",
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:    "persistent",
			Aliases: []string{"P"},
			Usage:   "make a non-temporary container",
			Value:   false,
		},
		&cli.StringFlag{
			Name:   "v",
			Usage:  "[disabled] specify the `volume:path` to attach to the container",
			Hidden: true,
		},
		&cli.StringFlag{
			Name:   "device",
			Usage:  "[disabled] specify device `path` to attach to the container",
			Hidden: true,
		},
	},
}

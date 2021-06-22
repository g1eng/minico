package cmd

import (
	"os"

	"github.com/g1eng/minico/cmd/subcommand"
	"github.com/urfave/cli/v2"
)

func Cmd() error {

	app := &cli.App{
		Name:    "minico",
		Version: "v0.8.1",
		//Usage:  "[cosh-ng] launch, attach, connect containers for your GUI applications",
		Usage:     "cli syntax sugar for container host operators",
		ArgsUsage: "[command on container]",
		Flags:     globalOptions,
		Commands: []*cli.Command{
			subcommand.Run,
			subcommand.List,
			subcommand.Stop,
			subcommand.Images,
			subcommand.Copy,
			subcommand.Cdi,
			subcommand.Config,
			subcommand.Version,
		},
		HideVersion: true,
	}
	return app.Run(os.Args)
}

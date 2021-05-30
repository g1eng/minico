package subcommand

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"strings"
)

var Version = &cli.Command{
	Name:  "version",
	Usage: "show version",
	Action: func(cli *cli.Context) error {
		fmt.Printf(
			"%s-%s\n",
			cli.App.Name,
			strings.TrimLeft(cli.App.Version, "v"),
		)
		return nil
	},
	ArgsUsage: " ",
}

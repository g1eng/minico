package subcommand

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

func unimplemented(_ *cli.Context) error {
	return fmt.Errorf("unimplemented")
}

var Copy = &cli.Command{
	Name:        "copy",
	Usage:       "[unimplemented] copy resources between host and containers",
	Description: "copy resources between host and containers, or between containers",
	Action:      unimplemented,
	Hidden:      true,
}

var Cdi = &cli.Command{
	Name:   "cdi",
	Usage:  "[unimplemented] manage opencdi stack",
	Action: unimplemented,
	Hidden: true,
}

var Config = &cli.Command{
	Name:   "config",
	Usage:  "[unimplemented] manage running configuration",
	Action: unimplemented,
	Hidden: true,
}

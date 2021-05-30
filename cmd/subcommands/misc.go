package subcommand

import (
	"github.com/g1eng/minico/functions"
	"github.com/urfave/cli/v2"
)

var Copy = &cli.Command{
	Name:        "copy",
	Usage:       "[unimplemented] copy resources between host and containers",
	Description: "copy resources between host and containers, or between containers",
	Action:      functions.Unimplemented,
	Hidden:      true,
}

var Cdi = &cli.Command{
	Name:   "cdi",
	Usage:  "[unimplemented] manage opencdi stack",
	Action: functions.Unimplemented,
	Hidden: true,
}

var Config = &cli.Command{
	Name:   "config",
	Usage:  "[unimplemented] manage running configuration",
	Action: functions.Unimplemented,
	Hidden: true,
}

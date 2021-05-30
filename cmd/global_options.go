package cmd

import "github.com/urfave/cli/v2"

var globalOptions = []cli.Flag{
	&cli.StringFlag{
		Name:    "u",
		Aliases: []string{"user"},
		Usage:   "[disabled] specify user `name` for execution in the container",
		Hidden:  true,
	},
	&cli.StringFlag{
		Name:    "c",
		Aliases: []string{"con"},
		Usage:   "[disabled] specify container `name` where the command executes",
		Hidden:  true,
	},
	&cli.StringFlag{
		Name:    "n",
		Aliases: []string{"net"},
		Usage:   "[disabled] set network bridge `name` attached to the container",
		Hidden:  true,
	},
	&cli.StringFlag{
		Name:   "e",
		Usage:  "[disabled] specify environmental variable (`key=val`)",
		Hidden: true,
	},
	&cli.StringFlag{
		Name:    "H",
		Aliases: []string{"host"},
		Usage:   "[disabled] set socket `URI` to switch endpoint",
		Hidden:  true,
	},
	&cli.StringFlag{
		Name:   "display",
		Usage:  "[disabled] set DISPLAY (`:X`) attached to the container",
		Hidden: true,
	},
	&cli.StringFlag{
		Name:   "locale",
		Hidden: true,
		Usage:  "[disabled] set locale `name` to use",
	},
	&cli.StringFlag{
		Name:    "I",
		Aliases: []string{"ime"},
		Usage:   "[disabled] set IME `name` to use",
		Hidden:  true,
	},
}

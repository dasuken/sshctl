package actions

import (
	"github.com/urfave/cli/v2"
)

func MakeApp(name, usage, version string) *cli.App {
	app := cli.NewApp()
	app.Name = name
	app.Usage = usage
	app.Version = version
	app.Commands = commands()
	return app
}

func commands() []*cli.Command {
	return []*cli.Command {
		{
			Name:   "add",
			Usage:  "add new config",
			Action: add,
		},
		{
			Name:   "list",
			Usage:  "list all config",
			Action: list,
		},
	}
}

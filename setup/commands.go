package setup

import (
	"vallese-cli/commands"

	"github.com/urfave/cli"
)

func Commands(app *cli.App) {

	app.Commands = []cli.Command{
		{
			Name:    "ask",
			Aliases: []string{"a"},
			Usage:   "ask a question to open ai",
			Action: func(c *cli.Context) {
				commands.ExecuteBash(c)
			},
		},
		{
			Name:    "listen",
			Aliases: []string{"la"},
			Usage:   "ask a question to open ai",
			Action: func(c *cli.Context) {
				commands.Listen(c)
			},
		},
	}
}

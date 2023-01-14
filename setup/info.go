package setup

import (
	"github.com/urfave/cli"
)

func Info(app *cli.App) {
	app.Name = "OpenAI CLI"
	app.Usage = "CLI that helps integrate with openai interface to use chatgpt features"
	app.Author = "vallese"
	app.Version = "0.0.1"
}

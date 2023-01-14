package main

import (
	"log"
	"os"
	"vallese-cli/setup"

	"github.com/urfave/cli"
)

func main() {
	var app = cli.NewApp()
	setup.Info(app)
	setup.Commands(app)
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

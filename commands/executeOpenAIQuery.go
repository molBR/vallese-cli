package commands

import (
	"fmt"
	"vallese-cli/out"

	"github.com/urfave/cli"
)

type RequestOpenAi struct {
	Model       string  `json:"model"`
	Prompt      string  `json:"prompt"`
	MaxToken    int32   `json:"max_tokens"`
	Temperature float32 `json:"temperature"`
}

func ExecuteBash(c *cli.Context) {
	prompt := ""
	for _, s := range c.Args() {
		prompt = prompt + " " + s
	}
	fmt.Println(prompt)
	out.SendRequestToOpenAI(prompt)

}

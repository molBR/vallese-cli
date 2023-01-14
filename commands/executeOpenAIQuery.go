package commands

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

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
	size := int32(4090 - len(prompt)/4)
	body := &RequestOpenAi{
		"text-davinci-003",
		prompt,
		size,
		0.9,
	}

	payloadBuf := new(bytes.Buffer)
	url := "https://api.openai.com/v1/completions"
	json.NewEncoder(payloadBuf).Encode(body)

	req, _ := http.NewRequest("POST", url, payloadBuf)
	//AUTHORIZATION HEADER
	token := os.Getenv("TOKEN_OPENAI")
	bearerToken := "Bearer " + token
	req.Header.Set("Authorization", bearerToken)
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	res, e := client.Do(req)
	if e != nil {
		return
	}

	defer res.Body.Close()
	fmt.Println("response Status:", res.Status)
	// Print the body to the stdout
	io.Copy(os.Stdout, res.Body)

}

package out

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type RequestOpenAi struct {
	Model       string  `json:"model"`
	Prompt      string  `json:"prompt"`
	MaxToken    int32   `json:"max_tokens"`
	Temperature float32 `json:"temperature"`
}

type ResponseOpenAi struct {
	Choices []Choices `json:"choices"`
}

type Choices struct {
	Text string `json:"text"`
}

func SendRequestToOpenAI(prompt string) {

	size := int32(4080 - len(prompt)/4)
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
	var roa ResponseOpenAi
	responseBytes, _ := ioutil.ReadAll(res.Body)
	json.Unmarshal(responseBytes, &roa)
	// Print the body to the stdout
	fmt.Println(roa.Choices[0].Text)
}

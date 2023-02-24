package out

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type ResponseFileUpload struct {
	Url string `json:"upload_url"`
}

type ResponseTranscribe struct {
	Id string `json:"id"`
}

type ResponseTrascribeResult struct {
	Status string `json:"status"`
	Text   string `json:"text"`
}

func SendRequestFileUpload() ResponseFileUpload {
	file, err := os.Open("audio.aiff")
	if err != nil {
		fmt.Println("Error opening file:", err)

	}
	defer file.Close()
	// Create a new HTTP request
	request, err := http.NewRequest("POST", "https://api.assemblyai.com/v2/upload", file)
	if err != nil {
		fmt.Println("Error creating request:", err)

	}
	apiToken := os.Getenv("TOKEN_TRANSCRIBE")
	// Set the request headers
	request.Header.Set("Content-Type", "audio/aiff")
	request.Header.Set("Authorization", apiToken)
	request.Header.Set("Transfer-Encoding", "chunked")

	// Send the request
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		fmt.Println("Error sending request:", err)

	}
	defer response.Body.Close()

	// Read the response body
	responseBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)

	}

	// Print the response body
	var rfu ResponseFileUpload
	json.Unmarshal(responseBytes, &rfu)
	return rfu
}

func SendRequestTranscribe(url string) ResponseTranscribe {

	apiToken := os.Getenv("TOKEN_TRANSCRIBE")
	audioURL := url

	// Create a request body
	requestBody, err := json.Marshal(map[string]string{
		"audio_url":     audioURL,
		"language_code": "pt",
	})
	if err != nil {
		fmt.Println("Error creating request body:", err)
	}

	// Create a new HTTP request
	request, err := http.NewRequest("POST", "https://api.assemblyai.com/v2/transcript", bytes.NewBuffer(requestBody))
	if err != nil {
		fmt.Println("Error creating request:", err)
	}

	// Set the request headers
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", apiToken)

	// Send the request
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		fmt.Println("Error sending request:", err)
	}
	defer response.Body.Close()

	// Read the response body
	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
	}

	var rft ResponseTranscribe
	json.Unmarshal(responseBody, &rft)

	return rft
}

func SendRequestTranscribeId(id string) ResponseTrascribeResult {

	apiToken := os.Getenv("TOKEN_TRANSCRIBE")
	idUrl := "https://api.assemblyai.com/v2/transcript/" + id

	// Create a request body

	// Create a new HTTP request
	request, err := http.NewRequest("GET", idUrl, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)

	}

	// Set the request headers
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", apiToken)

	// Send the request
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		fmt.Println("Error sending request:", err)

	}
	defer response.Body.Close()

	// Read the response body
	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
	}
	var rftr ResponseTrascribeResult
	json.Unmarshal(responseBody, &rftr)
	return rftr
}

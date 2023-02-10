package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	url := ""

	// Define the JSON-RPC request as a Go struct
	request := struct {
		JSONRPC string        `json:"jsonrpc"`
		Method  string        `json:"method"`
		Params  []interface{} `json:"params"`
		ID      int           `json:"id"`
	}{
		JSONRPC: "2.0",
		Method:  "Filecoin.ChainHead",
		Params:  []interface{}{},
		ID:      1,
	}

	// Marshal the Go struct into a JSON-RPC request message
	requestMessage, err := json.Marshal(request)
	if err != nil {
		fmt.Println("Error marshaling JSON-RPC request:", err)
		return
	}

	// Send the HTTP POST request with the JSON-RPC request message
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(requestMessage))
	if err != nil {
		fmt.Println("Error sending HTTP request:", err)
		return
	}

	// Read the HTTP response body
	defer resp.Body.Close()
	responseMessage, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading HTTP response:", err)
		return
	}
	type ChainHeadResult struct {
		Cids []map[string]string `json:"Cids"`
	}

	// Unmarshal the JSON-RPC response message into a Go struct
	var response ChainHeadResult

	// print the response	message
	fmt.Println("JSON-RPC response:", string(responseMessage))

	err = json.Unmarshal([]byte(responseMessage), &response)
	if err != nil {
		fmt.Println("Error unmarshaling JSON-RPC response:", err)
		return
	}

	// Log the response
	fmt.Println("JSON-RPC response:", response)
}

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// baseUrl := "wss://solana-mainnet.g.alchemy.com/v2/"
	// baseUrl := "https://solana-mainnet.g.alchemy.com/v2/"
	// apiKey := os.Getenv("ALCHEMY_SOLANA_API_KEY")
	// url := baseUrl + apiKey
	url := "https://api.mainnet-beta.solana.com"

	// health_request := struct {
	// 	JSONRPC string        `json:"jsonrpc"`
	// 	Method  string        `json:"method"`
	// 	Params  []interface{} `json:"params"`
	// 	ID      int           `json:"id"`
	// }{
	// 	JSONRPC: "2.0",
	// 	Method:  "getHealth",
	// 	ID:      1,
	// }

	getEpochInfoRequest := struct {
		JSONRPC string        `json:"jsonrpc"`
		Method  string        `json:"method"`
		Params  []interface{} `json:"params"`
		ID      int           `json:"id"`
	}{
		JSONRPC: "2.0",
		Method:  "getEpochInfo",
		ID:      1,
	}

	getEpochInfoRequestMessage, err := json.Marshal(getEpochInfoRequest)
	if err != nil {
		log.Fatal(err)
	}

	getEpochInfoRequestResponse, err := http.Post(url, "application/json", bytes.NewBuffer(getEpochInfoRequestMessage))
	if err != nil {
		log.Fatal(err)
	}
	getEpochInfoRequestResponseBytes, err := ioutil.ReadAll(getEpochInfoRequestResponse.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("getEpochInfoRequestResponseBytes:", string(getEpochInfoRequestResponseBytes))

	// type HealthResponse struct {
	// 	Jsonrpc string `json:"jsonrpc"`
	// 	Error   struct {
	// 		Code    int    `json:"code"`
	// 		Message string `json:"message"`
	// 		Data    struct {
	// 		} `json:"data"`
	// 	} `json:"error"`
	// 	ID int `json:"id"`
	// }

	// var healthResponse HealthResponse

	// Decode the JSON-RPC response message into a Go struct

}

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

func createRequestMessage(methodName string, params []interface{}) []byte {
	request := struct {
		JSONRPC string        `json:"jsonrpc"`
		Method  string        `json:"method"`
		Params  []interface{} `json:"params"`
		ID      int           `json:"id"`
	}{
		JSONRPC: "2.0",
		Method:  methodName,
		Params:  params,
		ID:      1,
	}

	requestMessage, err := json.Marshal(request)
	if err != nil {
		fmt.Println("Error marshalling request message:", err)
	}

	return requestMessage
}
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

	// Get Epoch Info

	requestMessage := createRequestMessage("getEpochInfo", []interface{}{})

	response, err := http.Post(url, "application/json", bytes.NewBuffer(requestMessage))

	responseBytes, err := ioutil.ReadAll(response.Body)

	fmt.Println("getEpochInfoRequestResponseBytes:", string(responseBytes))

}

package rpcMethods

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"solana_data/types"
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

func GetBlock(blockNumber int, url string) (types.Block, error) {
	var blockResponse types.Block

	methodName := "getBlock"
	params := []interface{}{blockNumber, map[string]interface{}{
		"encoding":                       "json",
		"maxSupportedTransactionVersion": 0,
		"transactionDetails":             "full",
		"rewards":                        false,
	}}

	requestMessage := createRequestMessage(methodName, params)
	response, err := http.Post(url, "application/json", bytes.NewBuffer(requestMessage))
	responseBytes, err := ioutil.ReadAll(response.Body)

	// read into blockResponse
	err = json.Unmarshal(responseBytes, &blockResponse)
	if err != nil {
		fmt.Println("Error unmarshalling block response:", err)
	}

	return blockResponse, err
}

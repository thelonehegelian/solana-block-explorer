// @todo handle errors properly

package rpcMethods

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"solana_data/types"
)

// creates the request message to send to the RPC endpoint in json format
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

// returns the block data for the given block number
func GetBlock(blockNumber int, nodeApi string) (types.Block, error) {
	var blockResponse types.Block

	methodName := "getBlock"
	params := []interface{}{blockNumber, map[string]interface{}{
		"encoding":                       "json",
		"maxSupportedTransactionVersion": 0,
		"transactionDetails":             "full",
		"rewards":                        false,
	}}

	requestMessage := createRequestMessage(methodName, params)
	response, err := http.Post(nodeApi, "application/json", bytes.NewBuffer(requestMessage))
	responseBytes, err := ioutil.ReadAll(response.Body)

	// read into blockResponse
	err = json.Unmarshal(responseBytes, &blockResponse)
	if err != nil {
		fmt.Println("Error unmarshalling block response:", err)
	}

	return blockResponse, err
}

// returns the current epoch data
func GetCurrentEpoch(nodeApi string) (types.CurrentEpoch, error) {
	var currentEpochResponse types.CurrentEpoch

	methodName := "getEpochInfo"
	params := []interface{}{}

	requestMessage := createRequestMessage(methodName, params)

	response, err := http.Post(nodeApi, "application/json", bytes.NewBuffer(requestMessage))
	// @todo see how to return errors
	if response.StatusCode != 200 {
		return currentEpochResponse, fmt.Errorf("Error getting current epoch: %v", response.StatusCode)
	}

	responseBytes, err := ioutil.ReadAll(response.Body)

	err = json.Unmarshal(responseBytes, &currentEpochResponse)
	if err != nil {
		fmt.Println("Error unmarshalling current epoch response:", err)
	}

	return currentEpochResponse, err
}

func GetTransactionBySignature(signature string, nodeApi string) (types.Transaction, error) {
	var transactionResponse types.Transaction

	methodName := "getTransaction"
	params := []interface{}{signature, map[string]interface{}{
		"encoding": "json",
	}}

	requestMessage := createRequestMessage(methodName, params)
	response, err := http.Post(nodeApi, "application/json", bytes.NewBuffer(requestMessage))

	responseBytes, err := ioutil.ReadAll(response.Body)
	err = json.Unmarshal(responseBytes, &transactionResponse)
	if err != nil {
		fmt.Println("Error unmarshalling transaction response:", err)
	}
	return transactionResponse, err
}

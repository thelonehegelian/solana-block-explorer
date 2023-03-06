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

type SolRequest struct {
	JSONRPC string        `json:"jsonrpc"`
	ID      uint64        `json:"id"`
	Method  string        `json:"method"`
	Params  []interface{} `json:"params"`
}

type SolGetBlockConfig struct {
	Encoding                       string `json:"encoding"`
	MaxSupportedTransactionVersion int    `json:"maxSupportedTransactionVersion"`
	TransactionDetails             string `json:"transactionDetails"`
	Rewards                        bool   `json:"rewards"`
}

type SolGetBlockParams struct {
	SlotNumber uint64
	Object     SolGetBlockConfig
}

// creates the request message to send to the RPC endpoint in json format
func createRequestMessage(method string, params []interface{}) []byte {
	request := SolRequest{
		JSONRPC: "2.0",
		ID:      1,
		Method:  method,
		Params:  params,
	}

	requestMessage, err := json.Marshal(request)
	if err != nil {
		fmt.Println("Error marshalling request message:", err)
	}

	return requestMessage
}

// returns the block data for the given block number
func GetBlock(blockNumber int, nodeApi string) (*types.Block, error) {
	getBlockParams := SolGetBlockParams{
		SlotNumber: uint64(blockNumber),
		Object: SolGetBlockConfig{
			Encoding:                       "json",
			MaxSupportedTransactionVersion: 0,
			TransactionDetails:             "full",
			Rewards:                        false,
		},
	}
	params := []interface{}{getBlockParams.SlotNumber, getBlockParams.Object}

	requestMessage := createRequestMessage("getBlock", params)
	response, err := http.Post(nodeApi, "application/json", bytes.NewBuffer(requestMessage))
	if err != nil {
		return nil, err
	}
	responseBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	// read into blockResponse
	var blockResponse *types.Block
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

func GetTransactionBySignature(signature string, nodeApi string) (*types.Transaction, error) {
	var transactionResponse types.Transaction
	// var transactionResponse map[string]interface{}

	methodName := "getTransaction"
	params := []interface{}{signature, map[string]interface{}{
		"encoding": "json",
	}}

	requestMessage := createRequestMessage(methodName, params)
	response, err := http.Post(nodeApi, "application/json", bytes.NewBuffer(requestMessage))
	_ = err
	responseBytes, err := ioutil.ReadAll(response.Body)
	_ = err
	err = json.Unmarshal(responseBytes, &transactionResponse)
	if err != nil {
		fmt.Println("Error unmarshalling transaction response:", err)
	}
	return &transactionResponse, err
}

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

type Block struct {
	Jsonrpc string `json:"jsonrpc"`

	Result struct {
		BlockHeight       int         `json:"blockHeight"`
		BlockTime         interface{} `json:"blockTime"`
		Blockhash         string      `json:"blockhash"`
		ParentSlot        int         `json:"parentSlot"`
		PreviousBlockhash string      `json:"previousBlockhash"`
		Transactions      []struct {
			Meta struct {
				Err               interface{}   `json:"err"`
				Fee               int           `json:"fee"`
				InnerInstructions []interface{} `json:"innerInstructions"`
				LogMessages       []interface{} `json:"logMessages"`
				PostBalances      []interface{} `json:"postBalances"`
				PostTokenBalances []interface{} `json:"postTokenBalances"`
				PreBalances       []interface{} `json:"preBalances"`
				PreTokenBalances  []interface{} `json:"preTokenBalances"`
				Rewards           interface{}   `json:"rewards"`
				Status            struct {
					Ok interface{} `json:"Ok"`
				} `json:"status"`
			} `json:"meta"`
			Transaction struct {
				Message struct {
					AccountKeys []string `json:"accountKeys"`
					Header      struct {
						NumReadonlySignedAccounts   int `json:"numReadonlySignedAccounts"`
						NumReadonlyUnsignedAccounts int `json:"numReadonlyUnsignedAccounts"`
						NumRequiredSignatures       int `json:"numRequiredSignatures"`
					} `json:"header"`
					Instructions []struct {
						Accounts       []int  `json:"accounts"`
						Data           string `json:"data"`
						ProgramIDIndex int    `json:"programIdIndex"`
					} `json:"instructions"`
					RecentBlockhash string `json:"recentBlockhash"`
				} `json:"message"`
				Signatures []string `json:"signatures"`
			} `json:"transaction"`
		} `json:"transactions"`
	} `json:"result"`
	ID int `json:"id"`
}

// ////////
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

func getBlock(blockNumber int, url string) (Block, error) {
	var blockResponse Block

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

	blockData, err := getBlock(1, url)

	// requestMessage := createRequestMessage(methodName, params)
	// response, err := http.Post(url, "application/json", bytes.NewBuffer(requestMessage))
	// responseBytes, err := ioutil.ReadAll(response.Body)
	// fmt.Println("response:", string(responseBytes))

	fmt.Println("blockData:", blockData.Result.Transactions[0].Meta.Fee)

}

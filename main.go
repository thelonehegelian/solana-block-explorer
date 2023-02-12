package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"solana_data/rpcMethods"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// @note can't use websockets because the functions are set up to send http requests
	// @todo add websocket support
	// baseUrl := "wss://solana-mainnet.g.alchemy.com/v2/"
	baseUrl := "https://solana-mainnet.g.alchemy.com/v2/"
	apiKey := os.Getenv("ALCHEMY_SOLANA_API_KEY")
	url := baseUrl + apiKey
	// url := "https://api.mainnet-beta.solana.com"

	// blockData, err := rpcMethods.GetBlock(1, url)
	// data, _ := json.MarshalIndent(blockData, "", " ")
	// print(string(data))

	// Epoch 409
	start_block := 176688000
	end_block := 176689000

	fmt.Println("blockNumber", "|", "blockHeight", "|", "blockTime", "|", "blockHash", "|", "prevBlockHash", "|", "txCount")
	fmt.Println("------------------------------------------------------------------------------------------------------------------")

	// loop through all the blocks in the epoch
	for i := start_block; i <= end_block; i++ {
		// sleep for 1 second
		time.Sleep(1 * time.Second)
		block, _ := rpcMethods.GetBlock(i, url)
		blockTime, _ := json.Marshal(block.Result.BlockTime)
		blockTimeInt, _ := strconv.ParseInt(string(blockTime), 10, 64)

		// @note block numbers are not always sequential so we need to check if the blockTime is 0
		if blockTimeInt != 0 {
			blockHeight, _ := json.Marshal(block.Result.BlockHeight)
			blockHash, _ := json.Marshal(block.Result.Blockhash)
			prevBlockHash, _ := json.Marshal(block.Result.PreviousBlockhash)
			txCount, _ := json.Marshal(len(block.Result.Transactions))

			fmt.Println(i, string(blockHeight), string(blockTime), string(blockHash), string(prevBlockHash), string(txCount))
		}

	}
}

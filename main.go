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
	// url := "https://api.mainnet-beta.solana.com" // there is a rate limit for this so doesn't work after some requests

	// blockData, err := rpcMethods.GetBlock(1, url)
	// data, _ := json.MarshalIndent(blockData, "", " ")
	// print(string(data))

	// Epoch 409, first
	start_block := 176688000
	// end_block := 176689000
	blockCount := 1
	end_block := 1000

	fmt.Println("blockNumber", "|", "blockHeight", "|", "blockTime", "|", "blockHash", "|", "prevBlockHash", "|", "txCount")
	fmt.Println("------------------------------------------------------------------------------------------------------------------")

	// loop through all the blocks in the epoch
	// @todo write to CSV
	// @todo refactor
	for {
		// sleep for 1 second
		time.Sleep(1 * time.Second)
		block, _ := rpcMethods.GetBlock(start_block, url)
		blockTime, _ := json.Marshal(block.Result.BlockTime)
		blockTimeInt, _ := strconv.ParseInt(string(blockTime), 10, 64)

		// @note block numbers are not always sequential so we need to check if the blockTime is 0
		if blockTimeInt != 0 {
			blockHeight, _ := json.Marshal(block.Result.BlockHeight)
			blockHash, _ := json.Marshal(block.Result.Blockhash)
			prevBlockHash, _ := json.Marshal(block.Result.PreviousBlockhash)

			// get transaction data from the blockchain here
			txCount, _ := json.Marshal(len(block.Result.Transactions))
			tx := block.Result.Transactions
			// txLen := len(tx)
			if len(tx) > 0 {
				for i := 0; i < len(tx); i++ {
					txSig := tx[i].Transaction.Signatures[0]

					fmt.Println(txSig)
					time.Sleep(5 * time.Second)
					rpcMethods.GetTransactionBySignature(txSig, url)

					// // id := transaction.Result.ID
					// blockTime, _ := json.Marshal(transaction.BlockTime)
					// fmt.Println(string(blockTime))

				}
			}
			fmt.Println(start_block, string(blockHeight), string(blockTime), string(blockHash), string(prevBlockHash), string(txCount))
			blockCount += 1
		}
		if blockCount > end_block {
			break
		}
		start_block += 1
	}
}

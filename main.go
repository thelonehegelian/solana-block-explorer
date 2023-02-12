package main

import (
	"encoding/csv"
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
	end_block := 4

	// @todo Header to the blocks.csv file
	fmt.Println("blockNumber", "|", "blockHeight", "|", "blockTime", "|", "blockHash", "|", "prevBlockHash", "|", "txCount")
	fmt.Println("------------------------------------------------------------------------------------------------------------------")

	file, err := os.Create("transactions.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Create a new CSV writer
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write the header row
	err = writer.Write([]string{"Timestamp", "Transaction Signature", "Transaction Slot", "Block Hash", "Recent Hash"})
	if err != nil {
		panic(err)
	}

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
			tx := block.Result.Transactions
			// @todo write to blocks.csv
			fmt.Println(start_block, string(blockHeight), string(blockTime), string(blockHash), string(prevBlockHash), len(tx))
			/**
			* get the transaction details for each transaction in the block and write to CSV
			 */
			if len(tx) > 0 {
				// @todo i < len(tx) testing with 2 transactions
				for i := 0; i < 2; i++ {
					txSig := tx[i].Transaction.Signatures[0]
					time.Sleep(1 * time.Second) // to avoid overloading the node
					txDetails, _ := rpcMethods.GetTransactionBySignature(txSig, url)
					txSlot := txDetails.Result.Slot
					txFee := strconv.FormatInt(int64(txDetails.Result.Meta.Fee), 10)
					recentBlockHash := txDetails.Result.Transaction.Message.RecentBlockhash
					// @todo add the two below to the CSV
					// timestamp := txDetails.BlockTime
					// txError := txDetails.Result.Meta.Err

					// write row to CSV
					row := []string{txSig, strconv.FormatInt(int64(txSlot), 10), string(blockHash), txFee, recentBlockHash}
					err = writer.Write(row)
					if err != nil {
						panic(err)
					}

				}
			}

			blockCount += 1
		}
		if blockCount > end_block {
			break
		}
		start_block += 1
	}
}

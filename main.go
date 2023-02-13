package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"solana_data/rpcMethods"
	"time"

	"github.com/joho/godotenv"
)

func toCSV(filename string, data [][]string) {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, row := range data {
		err = writer.Write(row)
		if err != nil {
			panic(err)
		}
	}
}

func toString(d interface{}) string {
	return fmt.Sprintf("%v", d)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

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

	// @todo Header to the blocks.csv file
	fmt.Println("blockNumber", "|", "blockHeight", "|", "blockTime", "|", "blockHash", "|", "prevBlockHash", "|", "txCount")
	fmt.Println("------------------------------------------------------------------------------------------------------------------")

	transactions := [][]string{}
	blocks := [][]string{}
	// Header of transactions.csv
	transactions = append(transactions,
		[]string{"Timestamp", "Transaction Signature", "Transaction Slot", "Block Hash", "Recent Hash", "txFee"},
	)

	// loop through all the blocks in the epoch
	// @todo write to CSV
	// @todo refactor
	startBlock := 176688000
	endBlock := startBlock + 4
	for currBlockNum := startBlock; currBlockNum < endBlock; currBlockNum++ {
		// sleep for 1 second to avoid overloading the node
		// time.Sleep(1 * time.Second)
		block, err := rpcMethods.GetBlock(currBlockNum, url)
		checkErr(err)
		if block.Result.BlockTime == 0 {
			fmt.Printf("blockTime %d is nil\n", currBlockNum)
			continue
		}

		// @note block numbers are not always sequential so we need to check if the blockTime is 0
		if block.Result.BlockTime != 0 {
			blockHeight := toString(block.Result.BlockHeight)
			blockHash := block.Result.Blockhash
			prevBlockHash := block.Result.PreviousBlockhash
			tx := block.Result.Transactions
			// @todo append to list and at the end write to blocks.csv
			row := []string{toString(currBlockNum), blockHeight, toString(block.Result.BlockTime), blockHash, prevBlockHash, toString(len(tx))}
			blocks = append(blocks, row)
			fmt.Println(currBlockNum, blockHeight, block.Result.BlockTime, blockHash, prevBlockHash, len(tx))
			/**
			* get the transaction details for each transaction in the block and write to CSV
			 */
			if len(tx) > 0 {
				// @todo i < len(tx) testing with 2 transactions
				for i := 0; i < 10; i++ {
					txSig := tx[i].Transaction.Signatures[0]
					time.Sleep(1 * time.Second) // to avoid overloading the node
					txDetails, _ := rpcMethods.GetTransactionBySignature(txSig, url)
					txSlot := toString(txDetails.Result.Slot)
					txFee := toString(txDetails.Result.Meta.Fee)
					recentBlockHash := txDetails.Result.Transaction.Message.RecentBlockhash
					// @todo add the two below to the CSV
					timestamp := toString(txDetails.BlockTime)
					fmt.Println("timestamp: ", timestamp)

					// txError := txDetails.Result.Meta.Err

					// write row to CSV
					row := []string{timestamp, txSig, txSlot, blockHash, recentBlockHash, txFee}
					transactions = append(transactions, row)
				}
			}
		}
	}
	fmt.Println("Transactions: ", len(transactions)-1)
	toCSV("transactions.csv", transactions)
	toCSV("blocks.csv", blocks)
}

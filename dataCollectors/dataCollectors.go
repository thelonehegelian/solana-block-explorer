package dataCollectors

import (
	"fmt"
	"solana_data/helpers"
	"solana_data/rpcMethods"
	"time"
)

// @todo return an error if there is one
// @todo can be refactored to be simpler?
func CollectBlocksToCsv(startBlock int, endBlock int, nodeApi string) {
	blocks := [][]string{}
	// add headers to the blocks.csv file
	blocks = append(blocks, []string{"blockNumber", "blockHeight", "blockTime", "blockHash", "prevBlockHash", "txCount"})

	for currBlockNum := startBlock; currBlockNum < endBlock; currBlockNum++ {
		// sleep for 1 second to avoid overloading the node
		time.Sleep(1 * time.Second)
		block, err := rpcMethods.GetBlock(currBlockNum, nodeApi)
		helpers.CheckErr(err)
		if block.Result.BlockTime == 0 {
			fmt.Printf("blockTime %d is nil\n", currBlockNum)
			continue
		}
		// @note block numbers are not always sequential so we need to check if the blockTime is 0
		if block.Result.BlockTime != 0 {
			blockHeight := helpers.ToString(block.Result.BlockHeight)
			blockHash := block.Result.Blockhash
			prevBlockHash := block.Result.PreviousBlockhash
			tx := block.Result.Transactions
			// @todo append to list and at the end write to blocks.csv
			row := []string{helpers.ToString(currBlockNum), blockHeight, helpers.ToString(block.Result.BlockTime), blockHash, prevBlockHash, helpers.ToString(len(tx))}
			blocks = append(blocks, row)
			fmt.Println(currBlockNum, blockHeight, block.Result.BlockTime, blockHash, prevBlockHash, len(tx))
		}
	}
	helpers.ToCSV("blocks.csv", blocks)

}

// @todo refactor
// @todo return error if there is one
func CollectTransactionsToCSV(startBlock int, endBlock int, nodeApi string) {
	transactions := [][]string{}
	// add headers to the transactions.csv
	transactions = append(transactions, []string{"timestamp", "txSig", "txSlot", "blockHash", "recentBlockHash", "txFee"})
	for currBlockNum := startBlock; currBlockNum < endBlock; currBlockNum++ {
		// sleep for 1 second to avoid overloading the node
		time.Sleep(1 * time.Second)
		block, err := rpcMethods.GetBlock(currBlockNum, nodeApi)
		helpers.CheckErr(err)
		if block.Result.BlockTime == 0 {
			fmt.Printf("blockTime %d is nil\n", currBlockNum)
			continue
		}

		// @note block numbers are not always sequential so we need to check if the blockTime is 0
		if block.Result.BlockTime != 0 {
			blockHash := block.Result.Blockhash
			tx := block.Result.Transactions
			// @todo append to list and at the end write to blocks.csv
			/**
			* get the transaction details for each transaction in the block and write to CSV
			 */
			if len(tx) > 0 {
				// @todo i < len(tx) testing with 2 transactions
				for i := 0; i < 2; i++ {
					txSig := tx[i].Transaction.Signatures[0]
					time.Sleep(1 * time.Second) // to avoid overloading the node
					txDetails, _ := rpcMethods.GetTransactionBySignature(txSig, nodeApi)
					txSlot := helpers.ToString(txDetails.Result.Slot)
					txFee := helpers.ToString(txDetails.Result.Meta.Fee)
					recentBlockHash := txDetails.Result.Transaction.Message.RecentBlockhash
					// @todo add the two below to the CSV
					timestamp := helpers.ToString(txDetails.BlockTime)
					fmt.Println("txSig: ", txSig)
					// txError := txDetails.Result.Meta.Err
					// write row to CSV
					row := []string{timestamp, txSig, txSlot, blockHash, recentBlockHash, txFee}
					transactions = append(transactions, row)
				}
			}
		}
	}
	helpers.ToCSV("transactions.csv", transactions)
}

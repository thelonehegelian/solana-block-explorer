package dataCollectors

import (
	"encoding/json"
	"fmt"
	"solana_data/helpers"
	"solana_data/rpcMethods"
	"solana_data/types"
	"time"
)

func transactionAsCsvRow(txSignature, blockHash string, tx *types.Transaction) []string {
	return []string{
		txSignature,
		blockHash,
		helpers.ToString(tx.BlockTime),
		helpers.ToString(tx.Result.Slot),
		tx.Result.Transaction.Message.RecentBlockhash,
		helpers.ToString(tx.Result.Meta.Fee),
	}
}

// @todo refactor
// @todo return error if there is one
func CollectTransactionsToCSV(startSlot int, endSlot int, nodeApi string) {
	transactions := [][]string{}
	// add headers to the transactions.csv
	transactions = append(transactions, []string{"txSig", "blockHash", "blockTime", "txSlot", "recentBlockHash", "txFee"})
	for currSlotNum := startSlot; currSlotNum < endSlot; currSlotNum++ {
		block, err := rpcMethods.GetBlock(currSlotNum, nodeApi)
		helpers.CheckErr(err)
		if block.Result.BlockTime == 0 {
			// @note block numbers are not always sequential so we need to check if the blockTime is 0
			fmt.Printf("blockTime %d is nil\n", currSlotNum)
			continue
		}

		// blockHash := block.Result.Blockhash
		tx := block.Result.Transactions
		// get the transaction details for each transaction in the block and write to CSV
		if len(tx) > 0 {
			// @todo i < len(tx) testing with 2 transactions
			for i := 0; i < 2 && i < len(tx); i++ {
				txSig := tx[i].Transaction.Signatures[0]
				// fmt.Println("txSig: ", txSig)
				txDetails, _ := rpcMethods.GetTransactionBySignature(txSig, nodeApi)
				bts, _ := json.MarshalIndent(txDetails, "", "  ")
				fmt.Println(string(bts))
				// row := transactionAsCsvRow(txSig, blockHash, &txDetails)
				// transactions = append(transactions, row)
				time.Sleep(1000 * time.Millisecond) // to avoid overloading the node
			}
		}
		// sleep for 0.1 second to avoid overloading the node
		time.Sleep(100 * time.Millisecond)
	}
	// helpers.ToCSV("transactions.csv", transactions)
}

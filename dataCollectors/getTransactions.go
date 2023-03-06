package dataCollectors

import (
	"encoding/json"
	"fmt"
	"solana_data/helpers"
	"solana_data/rpcMethods"
	"solana_data/types"
	"time"
)

var TRANSACTIONS_HEADER = []string{
	"block_BlockHeight",
	"block_BlockTime",
	"block_Blockhash",
	"block_ParentSlot",
	"block_PreviousBlockhash",
	"block_lenTransactions",

	"tx_Signature",
	"tx_Slot",
	"tx_BlockTime",
	"tx_Fee",
	"tx_Status",
	"tx_Err",
	"tx_RecentBlockhash",
	"tx_NumRequiredSignatures",
	"tx_lenSignatures",
	"tx_AccountKey0",
	"tx_AccountKey1",
	"tx_AccountKeys",
}

func transactionAsCsvRow(block *types.Block, txSignature string, tx *types.Transaction) []string {
	return []string{
		helpers.ToString(block.Result.BlockHeight),
		helpers.ToString(block.Result.BlockTime),
		helpers.ToString(block.Result.Blockhash),
		helpers.ToString(block.Result.ParentSlot),
		helpers.ToString(block.Result.PreviousBlockhash),
		helpers.ToString(len(block.Result.Transactions)),

		helpers.ToString(txSignature),
		helpers.ToString(tx.Result.Slot),
		helpers.ToString(tx.Result.BlockTime),
		helpers.ToString(tx.Result.Meta.Fee),
		helpers.ToString(tx.Result.Meta.Status.Ok),
		helpers.ToString(tx.Result.Meta.Err),
		helpers.ToString(tx.Result.Transaction.Message.RecentBlockhash),
		helpers.ToString(tx.Result.Transaction.Message.Header.NumRequiredSignatures),
		helpers.ToString(len(tx.Result.Transaction.Signatures)),
		helpers.ToString(tx.Result.Transaction.Message.AccountKeys[0]),
		helpers.ToString(tx.Result.Transaction.Message.AccountKeys[1]),
		helpers.ToString(tx.Result.Transaction.Message.AccountKeys[2:]),
	}
}

func CollectTransactionsToCSV(startSlot int, endSlot int, nodeApi string) {
	txTable := [][]string{}
	txTable = append(txTable, TRANSACTIONS_HEADER)
	// Per each splot's block, get the transactions
	for currSlotNum := startSlot; currSlotNum < endSlot; currSlotNum++ {
		block, err := rpcMethods.GetBlock(currSlotNum, nodeApi)
		helpers.CheckErr(err)
		if block.Result.BlockTime == 0 {
			// @note block numbers are not always sequential so we need to check if the blockTime is 0
			fmt.Printf("blockTime %d is nil\n", currSlotNum)
			continue
		}

		// get the transaction details for each transaction in the block and write to CSV
		transactions := block.Result.Transactions
		// txs2Get := 10
		for i := 0; i < len(transactions); i++ {
			// each transaction may have multiple signatures (?)
			for _, txSig := range transactions[i].Transaction.Signatures {
				txDetails, err := rpcMethods.GetTransactionBySignature(txSig, nodeApi)
				helpers.CheckErr(err)
				bts, _ := json.MarshalIndent(txDetails, "", "  ")
				fmt.Println(string(bts))
				if txDetails.Result.BlockTime == 0 {
					fmt.Printf("txDetails blockTime is nil\n")
					continue
				}

				row := transactionAsCsvRow(block, txSig, txDetails)
				txTable = append(txTable, row)
				time.Sleep(100 * time.Millisecond) // to avoid overloading the node
			}
		}
	}
	helpers.ToCSV("data/transactions.csv", txTable)
}

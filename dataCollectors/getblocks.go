package dataCollectors

import (
	"fmt"
	"solana_data/helpers"
	"solana_data/rpcMethods"
	"solana_data/types"
	"time"
)

func blockAsCSVRow(slotNumber string, block *types.Block) []string {
	return []string{
		slotNumber,
		helpers.ToString(block.Result.BlockHeight),
		helpers.ToString(block.Result.BlockTime),
		block.Result.Blockhash,
		block.Result.PreviousBlockhash,
		helpers.ToString(len(block.Result.Transactions)),
	}
}

// @todo return an error if there is one
// @todo can be refactored to be simpler?
func CollectBlocksToCsv(startSlot int, endSlot int, nodeApi string) {
	blocks := [][]string{}
	// add headers to the blocks.csv file
	blocks = append(blocks, []string{"slotNumber", "blockHeight", "blockTime", "blockHash", "prevBlockHash", "txCount"})

	for currSlotNum := startSlot; currSlotNum < endSlot; currSlotNum++ {
		// sleep for 0.1 second to avoid overloading the node
		time.Sleep(100 * time.Millisecond)
		block, err := rpcMethods.GetBlock(currSlotNum, nodeApi)
		helpers.CheckErr(err)
		if block.Result.BlockTime == 0 {
			fmt.Printf("blockTime %d is nil\n", currSlotNum)
			continue
		}
		// @note block numbers are not always sequential so we need to check if the blockTime is 0
		if block.Result.BlockTime != 0 {
			row := blockAsCSVRow(helpers.ToString(currSlotNum), block)
			blocks = append(blocks, row)
			fmt.Println(row)
		}
	}
	helpers.ToCSV("blocks.csv", blocks)
}

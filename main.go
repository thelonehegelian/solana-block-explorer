package main

import (
	"solana_data/dataCollectors"
)

func main() {

	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// }

	// @note can't use websockets because the functions are set up to send http requests
	// @todo add websocket support
	// baseUrl := "wss://solana-mainnet.g.alchemy.com/v2/"
	// baseUrl := "https://solana-mainnet.g.alchemy.com/v2/"
	// apiKey := os.Getenv("ALCHEMY_SOLANA_API_KEY")
	// url := baseUrl + apiKey
	url := "https://api.mainnet-beta.solana.com" // there is a rate limit for this so doesn't work after some requests

	startBlock := 176688000
	endBlock := startBlock + 4

	// dataCollectors.CollectBlocksToCsv(startBlock, endBlock, url)
	dataCollectors.CollectTransactionsToCSV(startBlock, endBlock, url)

}

package main

import (
	"fmt"
	"log"
	"os"
	"solana_data/dataCollectors"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	apiKey := os.Getenv("ALCHEMY_SOLANA_API_KEY")
	var url string
	if apiKey != "" {
		// @note can't use websockets because the functions are set up to send http requests
		// @todo add websocket support
		// baseUrl := "wss://solana-mainnet.g.alchemy.com/v2/"
		baseUrl := "https://solana-mainnet.g.alchemy.com/v2/"
		url = baseUrl + apiKey
		fmt.Println("Using Alchemy API with api key", baseUrl)
	} else {
		// there is a rate limit for this so doesn't work after some requests
		url = "https://api.mainnet-beta.solana.com"
		fmt.Println("Using Solana API:", url)
	}
	startBlock := 176688000
	endBlock := startBlock + 20

	// dataCollectors.CollectBlocksToCsv(startBlock, endBlock, url)
	dataCollectors.CollectTransactionsToCSV(startBlock, endBlock, url)

}

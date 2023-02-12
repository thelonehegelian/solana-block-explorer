package main

import (
	"encoding/json"
	"log"
	"solana_data/rpcMethods"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// baseUrl := "wss://solana-mainnet.g.alchemy.com/v2/"
	// baseUrl := "https://solana-mainnet.g.alchemy.com/v2/"
	// apiKey := os.Getenv("ALCHEMY_SOLANA_API_KEY")
	// url := baseUrl + apiKey
	url := "https://api.mainnet-beta.solana.com"

	// blockData, err := rpcMethods.GetBlock(1, url)
	// data, _ := json.MarshalIndent(blockData, "", " ")
	// print(string(data))

	epochData, err := rpcMethods.GetCurrentEpoch(url)
	data, _ := json.MarshalIndent(epochData, "", " ")
	print(string(data))

}

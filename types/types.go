package types

type Block struct {
	Jsonrpc string `json:"jsonrpc"`
	Result  struct {
		BlockHeight       int         `json:"blockHeight"`
		BlockTime         interface{} `json:"blockTime"`
		Blockhash         string      `json:"blockhash"`
		ParentSlot        int         `json:"parentSlot"`
		PreviousBlockhash string      `json:"previousBlockhash"`
		Transactions      []struct {
			Meta struct {
				Err               interface{}   `json:"err"`
				Fee               int           `json:"fee"`
				InnerInstructions []interface{} `json:"innerInstructions"`
				LogMessages       []interface{} `json:"logMessages"`
				PostBalances      []interface{} `json:"postBalances"`
				PostTokenBalances []interface{} `json:"postTokenBalances"`
				PreBalances       []interface{} `json:"preBalances"`
				PreTokenBalances  []interface{} `json:"preTokenBalances"`
				Rewards           interface{}   `json:"rewards"`
				Status            struct {
					Ok interface{} `json:"Ok"`
				} `json:"status"`
			} `json:"meta"`
			Transaction struct {
				Message struct {
					AccountKeys []string `json:"accountKeys"`
					Header      struct {
						NumReadonlySignedAccounts   int `json:"numReadonlySignedAccounts"`
						NumReadonlyUnsignedAccounts int `json:"numReadonlyUnsignedAccounts"`
						NumRequiredSignatures       int `json:"numRequiredSignatures"`
					} `json:"header"`
					Instructions []struct {
						Accounts       []int  `json:"accounts"`
						Data           string `json:"data"`
						ProgramIDIndex int    `json:"programIdIndex"`
					} `json:"instructions"`
					RecentBlockhash string `json:"recentBlockhash"`
				} `json:"message"`
				Signatures []string `json:"signatures"`
			} `json:"transaction"`
		} `json:"transactions"`
	} `json:"result"`
	ID int `json:"id"`
}

type CurrentEpoch struct {
	Jsonrpc string `json:"jsonrpc"`
	Result  struct {
		AbsoluteSlot     int `json:"absoluteSlot"`
		BlockHeight      int `json:"blockHeight"`
		Epoch            int `json:"epoch"`
		SlotIndex        int `json:"slotIndex"`
		SlotsInEpoch     int `json:"slotsInEpoch"`
		TransactionCount int `json:"transactionCount"`
	} `json:"result"`
	ID int `json:"id"`
}

type Transaction struct {
	Jsonrpc string `json:"jsonrpc"`
	Result  struct {
		Meta struct {
			Err               interface{}   `json:"err"`
			Fee               int           `json:"fee"`
			InnerInstructions []interface{} `json:"innerInstructions"`
			PostBalances      []interface{} `json:"postBalances"`
			PostTokenBalances []interface{} `json:"postTokenBalances"`
			PreBalances       []interface{} `json:"preBalances"`
			PreTokenBalances  []interface{} `json:"preTokenBalances"`
			Rewards           []interface{} `json:"rewards"`
			Status            struct {
				Ok interface{} `json:"Ok"`
			} `json:"status"`
		} `json:"meta"`
		Slot        int `json:"slot"`
		Transaction struct {
			Message struct {
				AccountKeys []string `json:"accountKeys"`
				Header      struct {
					NumReadonlySignedAccounts   int `json:"numReadonlySignedAccounts"`
					NumReadonlyUnsignedAccounts int `json:"numReadonlyUnsignedAccounts"`
					NumRequiredSignatures       int `json:"numRequiredSignatures"`
				} `json:"header"`
				Instructions []struct {
					Accounts       []int  `json:"accounts"`
					Data           string `json:"data"`
					ProgramIDIndex int    `json:"programIdIndex"`
				} `json:"instructions"`
				RecentBlockhash string `json:"recentBlockhash"`
			} `json:"message"`
			Signatures []string `json:"signatures"`
		} `json:"transaction"`
	} `json:"result"`
	BlockTime interface{} `json:"blockTime"`
	ID        int         `json:"id"`
}

package models

type Whale struct {
	Address  string `json:"address"`
	Amount   int64  `json:"amount"`
	Decimals int    `json:"decimals"`
	Owner    string `json:"owner"`
	Rank     int    `json:"rank"`
}

type TokenHeld struct {
	TokenAddress string      `json:"tokenAddress"`
	TokenAmount  TokenAmount `json:"tokenAmount"`
}

type TokenAmount struct {
	Amount         string  `json:"amount"`
	Decimals       int     `json:"decimals"`
	UiAmount       float64 `json:"uiAmount"`
	UiAmountString string  `json:"uiAmountString"`
}

package models

type Whale struct {
	Address  string `json:"address"`
	Amount   int64  `json:"amount"`
	Decimals int    `json:"decimals"`
	Owner    string `json:"owner"`
	Rank     int    `json:"rank"`
}

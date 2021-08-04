package database

type Wallet struct {
	ID    string `json:"id"`
	Asset *Asset `json:"asset"`
}
type Asset struct {
	Hash    string  `json:"hashid"`
	Balance float64 `json:"balance"`
}

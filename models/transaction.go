package models

type Transaction struct {
	TransactionID string  `json:"transactionId"`
	CardNumber    string  `json:"cardNumber"`
	Type          string  `json:"type"`
	Amount        float64 `json:"amount"`
	Status        string  `json:"status"`
	Timestamp     string  `json:"timestamp"`
}

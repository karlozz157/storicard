package entity

import "time"

type Summary struct {
	Balance              float64           `json:"balance"`
	AverageDebit         float64           `json:"average_debit"`
	AverageCredit        float64           `json:"average_credit"`
	NumberOfTransactions map[time.Time]int `json:"number_of_transaction"`
}

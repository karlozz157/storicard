package entity

import "time"

type Transaction struct {
	Date   time.Time `json:"date"`
	Amount float64   `json:"amount"`
}

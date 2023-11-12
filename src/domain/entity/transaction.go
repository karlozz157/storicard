package entity

import "time"

type Transaction struct {
	Date   time.Time `json:"date" bson:"data"`
	Amount float64   `json:"amount" bson:"amount"`
	Email  string    `json:"email" bson:"email"`
}

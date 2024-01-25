package entity

import (
	"time"
)

type Transaction struct {
	ID         int       `json:"id"`
	UserID     string    `json:"customer_id"`
	MerchantID string    `json:"merchant_id"`
	Amount     uint64    `json:"amount"`
	CreatedAt  time.Time `json:"created_at"`
}

package dto

import (
	"time"
)

type TransactionDto struct {
	ID        int         `json:"id"`
	Customer  UserDto     `json:"customer"`
	Merchant  MerchantDto `json:"merchant"`
	Amount    uint64      `json:"amount"`
	CreatedAt time.Time   `json:"created_at"`
}

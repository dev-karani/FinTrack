package transactions

import (
	"time"
)

type createTransactionRequest struct {
	Amount      int64  `json:"amount"`
	Label       string `json:"label"`
	Category    string `json:"category_type"`
	Source      string `json:"origin_account"`
	Destination string `json:"destination_account"`
}

type createTransactionResponse struct {
	Amount      int64     `json:"amount"`
	Label       string    `json:"label"`
	Category    string    `json:"category_type"`
	Source      string    `json:"origin_account"`
	Destination string    `json:"destination_account"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type updateTransactionRequest struct {
	Amount      int64  `json:"amount"`
	Label       string `json:"label"`
	Category    string `json:"category"`
	Source      string `json:"source"`
	Destination string `json:"destination"`
}

type updateTransactionResponse struct {
	Amount      int64     `json:"amount"`
	Label       string    `json:"label"`
	Category    string    `json:"category"`
	Source      string    `json:"source"`
	Destination string    `json:"destination"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

package transactions

import (
	"time"

	"github.com/google/uuid"
)

type createTransactionRequest struct {
	Amount      int64  `json:"amount"`
	Label       string `json:"label"`
	Category    string `json:"category_type"`
	Source      string `json:"origin_account"`
	Destination string `json:"destination_account"`
}

type createTransactionResponse struct {
	ID          uuid.UUID `json:"id"`
	Amount      int64     `json:"amount"`
	Label       string    `json:"label"`
	Category    string    `json:"category_type"`
	Source      string    `json:"origin_account"`
	Destination string    `json:"destination_account"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type updateTransactionRequest struct {
	ID          uuid.UUID `json:"id"`
	Amount      int64     `json:"amount"`
	Label       string    `json:"label"`
	Category    string    `json:"category"`
	Source      string    `json:"source"`
	Destination string    `json:"destination"`
}

type updateTransactionResponse struct {
	ID          uuid.UUID `json:"id"`
	Amount      int64     `json:"amount"`
	Label       string    `json:"label"`
	Category    string    `json:"category"`
	Source      string    `json:"source"`
	Destination string    `json:"destination"`

	UpdatedAt time.Time `json:"updated_at"`
}

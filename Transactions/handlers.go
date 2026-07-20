package transactions

import (
	"encoding/json"
	"net/http"

	"github.com/dev-karani/FinTrack/internal/database"
	httpx "github.com/dev-karani/FinTrack/internal/httpX"
)

type Handler struct {
	service *Service
}

func NewHandler(db database.Queries, jwtSecret string) *Handler{
	service := NewService(db)

	return &Handler{
		service: service,
	}
}

func (h Handler) CreateTransaction (w http.ResponseWriter, r *http.Request) {
	decoder := json.Decoder(r.Body)
	req := createTransactionRequest{}
	if err := decoder.Decode(&req); err != nil {
		httpx.RespondWithError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	transaction, err := h.service.CreateTransaction(req.Amount, req.Category, req.Label, req.From, req.To)
	if err != nil {
		httpx.RespondWithError(w, http.StatusInternalServerError, "couldnt create user")
		return
	}

	httpx.RespondWithJSON(w, http.StatusCreated, createTransactionResponse{
		Amount: transaction.Amount,
		Label: transaction.Label,
		Category:,
		From:,
		To:,
		CreatedAAt:
		UpdatedAt:
	})
}

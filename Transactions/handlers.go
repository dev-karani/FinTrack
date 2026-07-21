package transactions

import (
	"encoding/json"
	"net/http"

	"github.com/dev-karani/FinTrack/internal/auth"
	"github.com/dev-karani/FinTrack/internal/database"
	httpx "github.com/dev-karani/FinTrack/internal/httpX"
)

type Handler struct {
	service *Service
}

func NewHandler(db *database.Queries, jwtSecret string) *Handler {
	service := NewService(db, jwtSecret)

	return &Handler{
		service: service,
	}
}

func (h Handler) CreateTransaction(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	req := createTransactionRequest{}
	if err := decoder.Decode(&req); err != nil {
		httpx.RespondWithError(w, http.StatusBadRequest, "invalid request body")
		return
	}
	token, err := auth.GetBearerToken(r.Header)
	if err != nil {
		httpx.RespondWithError(w, http.StatusBadRequest, "invalid token")
		return
	}

	transaction, err := h.service.CreateTransaction(token, req.Amount, req.Category, req.Label, req.Source, req.Destination)
	if err != nil {
		httpx.RespondWithError(w, http.StatusInternalServerError, "couldnt create user")
		return
	}

	httpx.RespondWithJSON(w, http.StatusCreated, createTransactionResponse{
		Amount:      int32(transaction.Amount),
		Label:       transaction.Label,
		Category:    transaction.Category,
		Source:      transaction.Source,
		Destination: transaction.Destination,
		CreatedAt:   transaction.CreatedAt,
		UpdatedAt:   transaction.UpdatedAt,
	})
}

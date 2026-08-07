package transactions

import (
	"encoding/json"
	"net/http"

	"github.com/dev-karani/FinTrack/internal/auth"
	"github.com/dev-karani/FinTrack/internal/database"
	httpx "github.com/dev-karani/FinTrack/internal/httpX"
	"github.com/google/uuid"
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

func (h *Handler) CreateTransaction(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	req := createTransactionRequest{}
	if err := decoder.Decode(&req); err != nil {
		httpx.RespondWithError(w, http.StatusBadRequest, "invalid req body")
		return
	}
	token, err := auth.GetBearerToken(r.Header)
	if err != nil {
		httpx.RespondWithError(w, http.StatusBadRequest, "invalid token")
		return
	}

	transaction, err := h.service.CreateTransaction(r.Context(), token, req.Amount, req.Category, req.Label, req.Source, req.Destination)
	if err != nil {
		httpx.RespondWithError(w, http.StatusInternalServerError, "couldnt create user")
		return
	}

	httpx.RespondWithJSON(w, http.StatusCreated, createTransactionResponse{
		ID:          transaction.ID,
		Amount:      transaction.Amount,
		Label:       transaction.Label,
		Category:    transaction.Category,
		Source:      transaction.Source,
		Destination: transaction.Destination,
		CreatedAt:   transaction.CreatedAt,
		UpdatedAt:   transaction.UpdatedAt,
	})
}

func (h *Handler) GetUserBalance(w http.ResponseWriter, r http.Request) {
	token, err := auth.GetBearerToken(r.Header)
	if err != nil {
		httpx.RespondWithError(w, http.StatusBadRequest, "invalid auth header")
		return
	}

	balance, err := h.service.GetUserBalance(r.Context(), token)
	if err != nil {
		httpx.RespondWithError(w, http.StatusInternalServerError, "failed to get user balance")
		return
	}

	httpx.RespondWithJSON(w, http.StatusOK, BalanceResponse{
		Balance: balance,
	})
}

func (h *Handler) GetUserIncome(w http.ResponseWriter, r *http.Request) {
	token, err := auth.GetBearerToken(r.Header)
	if err != nil {
		httpx.RespondWithError(w, http.StatusBadRequest, "invalid auth header")
		return
	}

	income, err := h.service.GetUserIncome(r.Context(), token)
	if err != nil {
		httpx.RespondWithError(w, http.StatusInternalServerError, "failed to get user income")
		return
	}
	httpx.RespondWithJSON(w, http.StatusOK, IncomeResponse{
		Income: income,
	})
}

func (h *Handler) GetUserExpenses(w http.ResponseWriter, r *http.Request) {
	token, err := auth.GetBearerToken(r.Header)
	if err != nil {
		httpx.RespondWithError(w, http.StatusBadRequest, "invalid auth header")
		return
	}

	expenses, err := h.service.GetUserExpenses(r.Context(), token)
	if err != nil {
		httpx.RespondWithError(w, http.StatusInternalServerError, "failed to get expenses")
		return
	}

	httpx.RespondWithJSON(w, http.StatusOK, ExpenseResponse{
		Expense: expenses,
	})
}

func (h *Handler) GetAllUserTransactions(w http.ResponseWriter, r *http.Request) {
	token, err := auth.GetBearerToken(r.Header)
	if err != nil {
		httpx.RespondWithError(w, http.StatusBadRequest, "invalid auth header")
		return
	}

	transactions, err := h.service.GetUserTransactions(r.Context(), token)
	if err != nil {
		httpx.RespondWithError(w, http.StatusInternalServerError, "failed to get transactions")
		return
	}

	transactionResponse := make([]createTransactionResponse, 0, len(transactions))
	for _, transaction := range transactions {
		transactionResponse = append(transactionResponse, createTransactionResponse{
			ID:          transaction.ID,
			Amount:      transaction.Amount,
			Label:       transaction.Label,
			Category:    transaction.Category,
			Source:      transaction.Source,
			Destination: transaction.Destination,
			CreatedAt:   transaction.CreatedAt,
			UpdatedAt:   transaction.UpdatedAt,
		})
	}
	httpx.RespondWithJSON(w, http.StatusOK, transactionResponse)
}

func (h *Handler) GetTransactionByID(w http.ResponseWriter, r *http.Request) {
	token, err := auth.GetBearerToken(r.Header)
	if err != nil {
		httpx.RespondWithError(w, http.StatusBadRequest, "invalid auth header")
		return
	}

	transactionID, err := uuid.Parse(r.PathValue("transactionID"))
	if err != nil {
		httpx.RespondWithError(w, http.StatusBadRequest, "invalid transactionID")
		return
	}

	transaction, err := h.service.GetTransactionByID(r.Context(), token, transactionID)
	if err != nil {
		httpx.RespondWithError(w, http.StatusInternalServerError, "failed to get single transaction")
		return
	}
	httpx.RespondWithJSON(w, http.StatusOK, createTransactionResponse{
		ID:          transaction.ID,
		Amount:      transaction.Amount,
		Label:       transaction.Label,
		Category:    transaction.Category,
		Source:      transaction.Destination,
		Destination: transaction.Source,
		CreatedAt:   transaction.CreatedAt,
		UpdatedAt:   transaction.UpdatedAt,
	})
}
func (h *Handler) UpdateTransaction(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	req := updateTransactionRequest{}
	if err := decoder.Decode(&req); err != nil {
		httpx.RespondWithError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	transactionID, err := uuid.Parse(r.PathValue("transactionID"))
	if err != nil {
		httpx.RespondWithError(w, http.StatusBadRequest, "invalid transaction ID")
		return
	}

	token, err := auth.GetBearerToken(r.Header)
	if err != nil {
		httpx.RespondWithError(w, http.StatusBadRequest, "missing auth header")
		return
	}

	transaction, err := h.service.UpdateTransactionByID(
		r.Context(),
		token,
		transactionID,
		req.Amount,
		req.Label,
		req.Category,
		req.Source,
		req.Destination,
	)

	if err != nil {
		httpx.RespondWithError(w, http.StatusInternalServerError, "failed to update transaction")
		return
	}

	httpx.RespondWithJSON(w, http.StatusOK, updateTransactionResponse{
		ID:          transaction.ID,
		Amount:      transaction.Amount,
		Label:       transaction.Label,
		Category:    transaction.Category,
		Source:      transaction.Source,
		Destination: transaction.Destination,
		UpdatedAt:   transaction.UpdatedAt,
	})
}

func (h *Handler) DeleteTransactionByID(w http.ResponseWriter, r *http.Request) {
	//get id from url PathValue
	transactionID, err := uuid.Parse(r.PathValue("transactionID"))
	if err != nil {
		httpx.RespondWithError(w, http.StatusBadRequest, "invalid transaction ID")
		return
	}
	// get token
	token, err := auth.GetBearerToken(r.Header)
	if err != nil {
		httpx.RespondWithError(w, http.StatusBadRequest, "missing auth header")
		return
	}

	//pass id to transaction service
	err = h.service.DeleteTransactionByID(r.Context(), token, transactionID)
	if err != nil {
		httpx.RespondWithError(w, http.StatusInternalServerError, "failed to delete transaction")
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

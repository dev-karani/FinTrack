package users

import (
	"encoding/json"
	"net/http"

	"github.com/dev-karani/FinTrack/internal/database"
	httpx "github.com/dev-karani/FinTrack/internal/httpX"
)

type Handler struct {
	service *Service
}

func NewHandler(db *database.Queries, jwtSecret string) *Handler {
	service := NewService(db)

	return &Handler{
		service: service,
	}
}

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	req := createUserRequest{}
	if err := decoder.Decode(&req); err != nil {
		httpx.RespondWithError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	user, err := h.service.CreateUser(r.Context(), req.Email, req.Password)
	if err != nil {
		httpx.RespondWithError(w, http.StatusInternalServerError, "could not create user")
		return
	}

	httpx.RespondWithJSON(w, http.StatusCreated, UserResponse{
		ID:        user.ID,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	})
}

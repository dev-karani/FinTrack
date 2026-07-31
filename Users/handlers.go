package users

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

func (h *Handler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	token, err := auth.GetBearerToken(r.Header)
	if err != nil {
		httpx.RespondWithError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	err = h.service.DeleteUser(r.Context(), token)
	if err != nil {
		httpx.RespondWithError(w, http.StatusInternalServerError, "")
	}
	w.WriteHeader(http.StatusNoContent)

}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	req := loginRequest{}
	if err := decoder.Decode(&req); err != nil {
		httpx.RespondWithError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	userLoginDetails, err := h.service.Login(r.Context(), req.Email, req.Password)
	if err != nil {
		httpx.RespondWithError(w, http.StatusUnauthorized, "invalid email or passwrod")
		return
	}

	httpx.RespondWithJSON(w, http.StatusOK, LoginResponse{
		ID:           userLoginDetails.User.ID,
		CreatedAt:    userLoginDetails.User.CreatedAt,
		UpdatedAt:    userLoginDetails.User.UpdatedAt,
		Email:        userLoginDetails.User.Email,
		JWTToken:     userLoginDetails.JWTToken,
		RefreshToken: userLoginDetails.RefreshToken,
	})
}

func (h *Handler) Revoke(w http.ResponseWriter, r *http.Request) {
	token, err := auth.GetBearerToken(r.Header)
	if err != nil {
		httpx.RespondWithError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	err = h.service.Revoke(r.Context(), token)
	if err != nil {
		httpx.RespondWithError(w, http.StatusInternalServerError, "could not revoke refrsh token")

	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *Handler) Refresh(w http.ResponseWriter, r *http.Request) {
	refreshToken, err := auth.GetBearerToken(r.Header)
	if err != nil {
		httpx.RespondWithError(w, http.StatusUnauthorized, "missing refresh token")
		return
	}

	verifiedToken, err := h.service.RefreshToken(r.Context(), refreshToken)
	if err != nil {
		httpx.RespondWithError(w, http.StatusUnauthorized, "invalid refresh token")
		return
	}

	newJwt := verifiedToken.JWTToken
	httpx.RespondWithJSON(w, http.StatusOK, RefreshTokenResponse{
		JWTToken: newJwt,
	})
}

func (h *Handler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	req := updateUserRequest{}
	if err := decoder.Decode(&req); err != nil {
		httpx.RespondWithError(w, http.StatusBadRequest, "invalid request body")
		return
	}
	if len(req.Password) < 8 {
		httpx.RespondWithError(w, http.StatusBadRequest, "invalid password length")
		return
	}

	token, err := auth.GetBearerToken(r.Header)
	if err != nil {
		httpx.RespondWithError(w, http.StatusUnauthorized, "missing token")
		return
	}

	user, err := h.service.UpdateUser(r.Context(), token, req.Email, req.Password)
	if err != nil {
		httpx.RespondWithError(w, http.StatusInternalServerError, "failed to update user")
		return
	}

	httpx.RespondWithJSON(w, http.StatusOK, updatedUserResponse{
		ID:        user.ID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Email:     user.Email,
	})
}

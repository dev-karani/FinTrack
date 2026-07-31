package users

import (
	"time"

	"github.com/google/uuid"
)

// create|login|update
type createUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type loginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	ID           uuid.UUID `json:"id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	Email        string    `json:"email"`
	JWTToken     string    `json:"jwt_token"`
	RefreshToken string    `json:"refresh_token"`
}
type updateUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type updatedUserResponse struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Email     string    `json:"email"`
}

// user response
type UserResponse struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Email     string    `json:"email"`
}

type RefreshTokenResponse struct {
	JWTToken string `json:"jwt_token"`
}

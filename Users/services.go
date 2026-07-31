package users

import (
	"context"
	"time"

	"github.com/dev-karani/FinTrack/internal/auth"
	"github.com/dev-karani/FinTrack/internal/database"
)

type Service struct {
	db        *database.Queries
	jwtSecret string
}

func NewService(db *database.Queries, jwtSecret string) *Service {
	return &Service{
		db:        db,
		jwtSecret: jwtSecret,
	}
}

func (s *Service) CreateUser(ctx context.Context, email string, password string) (database.User, error) {
	hashedPassword, err := auth.HashPassword(password)
	if err != nil {
		return database.User{}, err
	}

	dbUser, err := s.db.CreateUser(ctx, database.CreateUserParams{
		Email:          email,
		HashedPassword: hashedPassword,
	})
	if err != nil {
		return database.User{}, err
	}

	return dbUser, nil
}

func (s *Service) DeleteUser(ctx context.Context, token string) error {
	userID, err := auth.ValidateJWT(token, s.jwtSecret)
	if err != nil {
		return err
	}

	err = s.db.RevokeAllRefreshTokensForUser(ctx, userID)
	if err != nil {
		return err
	}
	err = s.db.DeleteUserByID(ctx, userID)
	if err != nil {
		return err
	}

	return nil

}

type LoginResult struct {
	User         database.User
	JWTToken     string
	RefreshToken string
}

func (s *Service) Login(ctx context.Context, email, password string) (LoginResult, error) {
	user, err := s.db.GetUserByEmail(ctx, email)
	if err != nil {
		return LoginResult{}, err
	}

	match, err := auth.CheckPasswordHash(password, user.HashedPassword)
	if err != nil || !match {
		return LoginResult{}, err
	}

	jwtToken, err := auth.MakeJWT(user.ID, s.jwtSecret, time.Hour)
	if err != nil {
		return LoginResult{}, err
	}

	refreshToken, err := auth.MakeRefreshToken()
	if err != nil {
		return LoginResult{}, err
	}

	now := time.Now().UTC()

	_, err = s.db.CreateRefreshToken(ctx, database.CreateRefreshTokenParams{
		Token:     refreshToken,
		CreatedAt: now,
		UpdatedAt: now,
		UserID:    user.ID,
		ExpiresAt: now.Add(60 * 24 * time.Hour),
	})
	if err != nil {
		return LoginResult{}, err
	}
	userLogin := LoginResult{
		User:         user,
		JWTToken:     jwtToken,
		RefreshToken: refreshToken,
	}

	return userLogin, nil
}

func (s *Service) Revoke(ctx context.Context, token string) error {
	err := s.db.RevokeRefreshToken(ctx, token)
	if err != nil {
		return err
	}
	return nil
}

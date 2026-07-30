package users

import (
	"context"

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
	err = s.db.DeleteUserByID(ctx, userID)
	if err != nil {
		return err
	}

	refreshToken, err := s.db.Get

	return nil

}

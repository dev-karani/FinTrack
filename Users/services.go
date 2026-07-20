package users

import (
	"context"

	"github.com/dev-karani/FinTrack/internal/auth"
	"github.com/dev-karani/FinTrack/internal/database"
)

type Service struct {
	db *database.Queries
}

func NewService(db *database.Queries) *Service {
	return &Service{
		db: db,
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

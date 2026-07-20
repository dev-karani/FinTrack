package transactions

import (
	"context"

	"github.com/dev-karani/FinTrack/internal/database"
)

type Service struct {
	db *database.Queries
}

func NewService(db database.Queries) *Service {
	return &Service{
		db: db,
	}
}

func (s *Service) CreateTransaction(ctx context.Context, amount int32, category, label, source, destination string) (database.User, error) {
	//get user id from validate jwt
	//pass post to create post fuunction
	//return
	return

}

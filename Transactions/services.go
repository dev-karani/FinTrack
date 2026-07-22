package transactions

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

func (s *Service) CreateTransaction(ctx context.Context, token string, amount int64, category, label, source, destination string) (database.Transaction, error) {
	userID, err := auth.ValidateJWT(token, s.jwtSecret)
	if err != nil {
		return database.Transaction{}, err
	}
	transaction, err := s.db.CreateTransaction(ctx, database.CreateTransactionParams{
		UserID:      userID,
		Amount:      amount,
		Label:       label,
		Category:    category,
		Source:      source,
		Destination: destination,
	})

	if err != nil {
		return database.Transaction{}, err
	}
	return transaction, nil

}

func (s *Service) GetUserTransactions(ctx context.Context, token string) (database.Transaction, error) {
	var transactions []database.Transaction

	userID, err := auth.ValidateJWT(token, s.jwtSecret)
	if err != nil {
		return database.Transaction{}, err
	}

	transactions, err ;= s.db.GetAllTransactionsByUserID(ctx, userID)
	if err !=nil {
		return database.Transaction{}, err
	}
	return transactions[], nil
} 

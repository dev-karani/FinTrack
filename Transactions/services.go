package transactions

import (
	"context"

	"github.com/dev-karani/FinTrack/internal/auth"
	"github.com/dev-karani/FinTrack/internal/database"

	"github.com/google/uuid"
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

// Get all user transactions
func (s *Service) GetUserTransactions(ctx context.Context, token string) ([]database.Transaction, error) {
	userID, err := auth.ValidateJWT(token, s.jwtSecret)
	if err != nil {
		return nil, err
	}

	transactions, err := s.db.GetAllTransactionsByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}
	return transactions, nil
}

// get transaction
func (s *Service) GetTransactionByID(ctx context.Context, token string) (database.Transaction, error) {

}
func (s *Service) UpdateTransactionByID(ctx context.Context, token string, transactionID uuid.UUID, amount int64, label, category, source, destination string) (database.Transaction, error) {
	userID, err := auth.ValidateJWT(token, s.jwtSecret)
	if err != nil {
		return database.Transaction{}, err
	}

	transaction, err := s.db.UpdateTransactionByID(ctx, database.UpdateTransactionByIDParams{
		UserID:      userID,
		ID:          transactionID,
		Amount:      amount,
		Category:    category,
		Label:       label,
		Source:      source,
		Destination: destination,
	})

	if err != nil {
		return database.Transaction{}, err
	}

	return transaction, nil
}

func (s *Service) DeleteTransactionByID(ctx context.Context, token string, transactionID uuid.UUID) error {
	userID, err := auth.ValidateJWT(token, s.jwtSecret)
	if err != nil {
		return err
	}

	err = s.db.DeleteTransactionByID(ctx, database.DeleteTransactionByIDParams{
		ID:     transactionID,
		UserID: userID,
	})
	if err != nil {
		return err
	}

	return nil
}

package transactions

import "fmt"

const (
	CategoryDebit  = "DEBIT"
	CategoryCredit = "CREDIT"
)

func validateCategory(category string) error {
	switch category {
	case CategoryCredit, CategoryDebit:
		return nil
	default:
		return fmt.Errorf("invalid category: %s", category)
	}
}

package services

import (
	"bread/backend/models"
	"bread/backend/persistence"
	"fmt"
)

type TransactionService struct{}

func (s *TransactionService) CreateTransaction(
	projectID int64,
	categoryID *int64,
	desc string,
	amount int64,
	date string,
	expenseType bool,
	notes string,
) (models.Transaction, error) {

	t := models.Transaction{
		Description: desc,
		ProjectID:     projectID,
		CategoryID:  categoryID,
		Date:        date,
		Amount:      amount,
		ExpenseType: expenseType,
		Notes: notes,
	}

	id, err := persistence.InsertTransaction(t, nil)
	if err != nil {
		return t, fmt.Errorf("CreateTransaction: %w", err)
	}
	t.ID = id
	return t, nil
}

func (s *TransactionService) GetTransaction(id int64) (models.Transaction, error) {
	t, err := persistence.GetTransaction(id, nil)
	if err != nil {
		return t, fmt.Errorf("GetTransaction: %w", err)
	}
	return t, nil
}

// ListTransactions returns all transactions optionally filtered by categoryID
func (s *TransactionService) ListTransactions(projectID int64, groupID, categoryID *int64) ([]models.Transaction, error) {
	result, err := persistence.ListTransactions(projectID, groupID, categoryID, nil)
	if err != nil {
		return nil, fmt.Errorf("ListTransactions: %w", err)
	}

	return result, err

}

// UpdateTransaction updates an existing transaction
func (s *TransactionService) UpdateTransaction(t models.Transaction) error {
	if err := persistence.UpdateTransaction(t, nil); err != nil {
		return fmt.Errorf("UpdateTransaction: %w", err)
	}
	return nil
}

// DeleteTransaction removes a transaction by ID
func (s *TransactionService) DeleteTransaction(id int64) error {
	if err := persistence.DeleteTransaction(id, nil); err != nil {
		return fmt.Errorf("DeleteTransaction: %w", err)
	}
	return nil
}


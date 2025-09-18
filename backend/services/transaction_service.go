package services

import (
	"bread/backend/models"
	"bread/backend/persistence"
	"fmt"
	"time"
)

type TransactionService struct{}

func (s *TransactionService) CreateTransaction(
	desc string,
	amount int64,
	date string,
	notes string,
	tags string,
	budgetID int64,
	groupID *int64,
	categoryID *int64,
) (models.Transaction, error) {
	currentTS := time.Now().Format("2006-01-02 15:04:05")

	t := models.Transaction{
		Description: desc,
		Amount:      amount,
		Date:        date,
		Notes:       notes,
		Tags:        tags,
		BudgetID:    budgetID,
		GroupID:     groupID,
		CategoryID:  categoryID,
		CreatedAt:   currentTS,
		UpdatedAt:   currentTS,
	}

	id, err := persistence.InsertTransaction(t)
	if err != nil {
		return t, fmt.Errorf("CreateTransaction: %w", err)
	}
	t.ID = id
	return t, nil
}

func (s *TransactionService) GetTransaction(id int64) (models.Transaction, error) {
	t, err := persistence.GetTransaction(id)
	if err != nil {
		return t, fmt.Errorf("GetTransaction: %w", err)
	}
	return t, nil
}

// ListTransactions returns all transactions optionally filtered by categoryID
func (s *TransactionService) ListTransactions(categoryID *int64) ([]models.Transaction, error) {
	all, err := persistence.GetTransactions()
	if err != nil {
		return nil, fmt.Errorf("ListTransactions: %w", err)
	}

	if categoryID == nil {
		return all, nil
	}

	// Filter by categoryID
	var filtered []models.Transaction
	for _, t := range all {
		if t.CategoryID != nil && *t.CategoryID == *categoryID {
			filtered = append(filtered, t)
		}
	}
	return filtered, nil
}

// UpdateTransaction updates an existing transaction
func (s *TransactionService) UpdateTransaction(t models.Transaction) error {
	t.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
	if err := persistence.UpdateTransaction(t); err != nil {
		return fmt.Errorf("UpdateTransaction: %w", err)
	}
	return nil
}

// DeleteTransaction removes a transaction by ID
func (s *TransactionService) DeleteTransaction(id int64) error {
	if err := persistence.DeleteTransaction(id); err != nil {
		return fmt.Errorf("DeleteTransaction: %w", err)
	}
	return nil
}


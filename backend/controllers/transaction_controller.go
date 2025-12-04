package controllers

import (
	"bread/backend/models"
	"bread/backend/services"
)

type TransactionController struct {
	service *services.TransactionService
}

func NewTransactionController() *TransactionController {
	return &TransactionController{
		service: &services.TransactionService{},
	}
}

// CreateTransaction inserts a new transaction and returns it
func (c *TransactionController) CreateTransaction(
	projectID int64,
	categoryID *int64,
	desc string,
	amount int64,
	date string,
	expenseType bool,
	notes string,
) (models.Transaction, error) {
	return c.service.CreateTransaction(projectID, categoryID, desc, amount, date, expenseType, notes)
}

// GetTransaction retrieves a transaction by ID
func (c *TransactionController) GetTransaction(id int64) (models.Transaction, error) {
	return c.service.GetTransaction(id)
}

// ListTransactions returns all transactions (optionally filtered by group or category)
func (c *TransactionController) ListTransactions(projectID int64, groupID, categoryID *int64, startDate, endDate *string) ([]models.Transaction, error) {
	return c.service.ListTransactions(projectID, groupID, categoryID, startDate, endDate)
}

// UpdateTransaction updates an existing transaction
func (c *TransactionController) UpdateTransaction(t models.Transaction) error {
	return c.service.UpdateTransaction(t)
}

// DeleteTransaction removes a transaction by ID
func (c *TransactionController) DeleteTransaction(id int64) error {
	return c.service.DeleteTransaction(id)
}


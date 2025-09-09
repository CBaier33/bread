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

func (c *TransactionController) AddTransaction(
	desc string,
	amount int64,
	date string,
	notes string,
	budgetID int64,
	groupID *int64,
	categoryID *int64,
) (models.Transaction, error) {
	return c.service.CreateTransaction(desc, amount, date, notes, budgetID, groupID, categoryID)
}

func (c *TransactionController) GetTransaction(id int64) (models.Transaction, error) {
	return c.service.GetTransaction(id)
}

func (c *TransactionController) ListTransactions(categoryID *int64) ([]models.Transaction, error) {
	return c.service.ListTransactions(categoryID)
}

func (c *TransactionController) UpdateTransaction(t models.Transaction) error {
	return c.service.UpdateTransaction(t)
}

func (c *TransactionController) DeleteTransaction(id int64) error {
	return c.service.DeleteTransaction(id)
}


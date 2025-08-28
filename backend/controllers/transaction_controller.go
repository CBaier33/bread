package controllers

import (
	"bread/backend/models"
	"bread/backend/services"
	"context"
)

type TransactionVM struct{
	ctx context.Context
}

func NewTransactionVM() *TransactionVM {
    return &TransactionVM{}
}

// Exposed to frontend via Wails
func (vm *TransactionVM) AddTransaction(desc string, amount float64, direction string, category string) (models.Transaction, error) {
    return services.CreateTransaction(desc, amount, direction, category)
}

func (vm *TransactionVM) GetTransactions() ([]models.Transaction, error) {
    return services.ListTransactions()
}




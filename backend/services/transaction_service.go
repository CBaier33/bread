package services

import (
	"bread/backend/models"
	"bread/backend/persistence"
	"time"
)

func CreateTransaction(desc string, amount float64, direction string, category string) (models.Transaction, error) {
    t := models.Transaction{
        Description: desc,
        Amount:      amount,
				Direction: 	 direction,
        Category:    category,
        CreatedAt:   time.Now().Format(time.RFC3339), // convert to string
    }

    id, err := persistence.InsertTransaction(t)
    if err != nil {
        return t, err
    }
    t.ID = id
    return t, nil
}

func ListTransactions() ([]models.Transaction, error) {
    return persistence.GetTransactions()
}



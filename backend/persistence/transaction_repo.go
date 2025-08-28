package persistence

import (
	"bread/backend/models"
)

func InsertTransaction(t models.Transaction) (int64, error) {
    stmt, err := DB.Prepare("INSERT INTO transactions(description, amount, direction, category, created_at) VALUES (?, ?, ?, ?, ?)")
    if err != nil {
        return 0, err
    }
    res, err := stmt.Exec(t.Description, t.Amount, t.Direction, t.Category, t.CreatedAt)
    if err != nil {
        return 0, err
    }
    return res.LastInsertId()
}

func GetTransactions() ([]models.Transaction, error) {
    rows, err := DB.Query("SELECT id, description, amount, direction, category, created_at FROM transactions ORDER BY created_at DESC")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var transactions []models.Transaction
    for rows.Next() {
        var t models.Transaction
        err = rows.Scan(&t.ID, &t.Description, &t.Amount, &t.Direction, &t.Category, &t.CreatedAt)
        if err != nil {
            return nil, err
        }
        transactions = append(transactions, t)
    }
    return transactions, nil
}


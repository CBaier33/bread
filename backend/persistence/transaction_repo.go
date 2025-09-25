package persistence

import (
	"bread/backend/models"
	"fmt"
	"time"
)

// InsertTransaction inserts a new transaction and returns its ID
func InsertTransaction(t models.Transaction) (int64, error) {
	res, err := DB.Exec(`
        INSERT INTO transactions( description, project_id, category_id, date, amount, expense_type, notes) 
				VALUES (?, ?, ?, ?, ?, ?, ?)`,
		t.Description,
		t.ProjectID,
		t.CategoryID,
		t.Date,
		t.Amount,
		t.ExpenseType,
		t.Notes,
	)
	if err != nil {
		return 0, err
	}

	return res.LastInsertId()
}

// GetTransactions returns all transactions
func ListTransactions() ([]models.Transaction, error) {
	query := `
		SELECT id, description, project_id, category_id, date, amount, expense_type, notes, created_at, updated_at
		FROM transactions
		ORDER BY created_at DESC;
    `

	rows, err := DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var transactions []models.Transaction
	for rows.Next() {
		var t models.Transaction
		if err := rows.Scan(
			&t.ID,
			&t.Description,
			&t.ProjectID,
			&t.CategoryID,
			&t.Date,
			&t.Amount,
			&t.ExpenseType,
			&t.Notes,
			&t.CreatedAt,
			&t.UpdatedAt,
		); err != nil {
			return nil, err
		}
		transactions = append(transactions, t)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return transactions, nil
}


// GetTransaction returns a single transaction by ID
func GetTransaction(id int64) (models.Transaction, error) {
	var t models.Transaction
	query := `
		SELECT 
			id, description, project_id, category_id, date, amount, expense_type, notes, created_at, updated_at
		FROM transactions
		WHERE id = ?;
	`
	row := DB.QueryRow(query, id)
	if err := row.Scan(
		&t.ID,
		&t.Description,
		&t.ProjectID,
		&t.CategoryID,
		&t.Date,
		&t.Amount,
		&t.ExpenseType,
		&t.Notes,
		&t.CreatedAt,
		&t.UpdatedAt,
	); err != nil {
		return t, fmt.Errorf("GetTransaction: %w", err)
	}
	return t, nil
}

// UpdateTransaction updates an existing transaction
func UpdateTransaction(t models.Transaction) error {
	if t.UpdatedAt == "" {
		t.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
	}

	_, err := DB.Exec(`
        UPDATE transactions
        SET description = ?, project_id = ?,  category_id = ?, date = ?, amount = ?, expense_type = ?, notes = ?, updated_at = (datetime('now'))
        WHERE id = ?`,
		t.Description,
		t.ProjectID,
		t.CategoryID,
		t.Date,
		t.Amount,
		t.ExpenseType,
		t.Notes,
		t.ID,
	)
	return err
}

// DeleteTransaction deletes a transaction by ID
func DeleteTransaction(id int64) error {
	_, err := DB.Exec(`DELETE FROM transactions WHERE id = ?`, id)
	return err
}


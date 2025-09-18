package persistence

import (
	"bread/backend/models"
	"fmt"
	"time"
)

// InsertTransaction inserts a new transaction and returns its ID
func InsertTransaction(t models.Transaction) (int64, error) {
	res, err := DB.Exec(`
        INSERT INTO transactions(
            description,
						budget_id,
            category_id,
            date,
            amount,
            notes,
						tags,
            created_at,
            updated_at
        ) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		t.Description,
		t.BudgetID,
		t.CategoryID,
		t.Date,
		t.Amount,
		t.Notes,
		t.Tags,
		t.CreatedAt,
		t.UpdatedAt,
	)
	if err != nil {
		return 0, err
	}

	return res.LastInsertId()
}

// GetTransactions returns all transactions
func GetTransactions() ([]models.Transaction, error) {
	query := `
		SELECT 
			t.id,
			t.description,
	    t.budget_id,
			t.category_id,
			COALESCE(c.name, '') AS category_name,
			t.date,
			t.amount,
			t.notes,
			t.tags,
			t.created_at,
			t.updated_at
		FROM transactions t
		LEFT JOIN categories c ON t.category_id = c.id
		ORDER BY t.created_at DESC;
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
			&t.BudgetID,
			&t.CategoryID,
			&t.CategoryName, // <- populated from join
			&t.Date,
			&t.Amount,
			&t.Notes,
			&t.Tags,
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
			id, description, category_id, date, amount, notes, tags, created_at, updated_at
		FROM transactions
		WHERE id = ?;
	`
	row := DB.QueryRow(query, id)
	if err := row.Scan(
		&t.ID,
		&t.Description,
		&t.BudgetID,
		&t.CategoryID,
		&t.Date,
		&t.Amount,
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
        SET description = ?, budget_id = ? category_id = ?, date = ?, amount = ?, notes = ?, updated_at = ?
        WHERE id = ?`,
		t.Description,
		t.BudgetID,
		t.CategoryID,
		t.Date,
		t.Amount,
		t.Notes,
		t.UpdatedAt,
		t.ID,
	)
	return err
}

// DeleteTransaction deletes a transaction by ID
func DeleteTransaction(id int64) error {
	_, err := DB.Exec(`DELETE FROM transactions WHERE id = ?`, id)
	return err
}


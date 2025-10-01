package persistence

import (
	"bread/backend/models"
	"fmt"
	"time"
)

// InsertTransaction inserts a new transaction and returns its ID
func InsertTransaction(t models.Transaction, db runner) (int64, error) {


	if db == nil {
		db = DB
	}

	res, err := db.Exec(`
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
func ListTransactions(projectID int64, groupID, categoryID *int64, db runner) ([]models.Transaction, error) {

	if db == nil {
		db = DB
	}

	var (
		query string
		accountIDs  []interface{}
	)

	if categoryID != nil {
		query = `
			SELECT id, description, project_id, category_id, date, amount, expense_type, notes, created_at, updated_at
			FROM transactions
			WHERE project_id = ? AND category_id = ?
			ORDER BY created_at DESC;
		`
		accountIDs = []interface{}{projectID, *categoryID}

	}

	if groupID != nil {

		query = `
			SELECT t.id, t.description, t.project_id, t.category_id, t.date, t.amount, t.expense_type, t.notes, t.created_at, t.updated_at
			FROM transactions t
			JOIN categories c ON c.id = t.category_id
			WHERE t.project_id = ? AND c.group_id = ?
			ORDER BY created_at DESC;
		`
		accountIDs = []interface{}{projectID, groupID}

	} else {

		query = `
			SELECT id, description, project_id, category_id, date, amount, expense_type, notes, created_at, updated_at
			FROM transactions
			WHERE project_id = ?
			ORDER BY created_at DESC;
		`
		accountIDs = []interface{}{projectID}
	}

	rows, err := db.Query(query, accountIDs...)
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
func GetTransaction(id int64, db runner) (models.Transaction, error) {


	if db == nil {
		db = DB
	}

	var t models.Transaction
	query := `
		SELECT 
			id, description, project_id, category_id, date, amount, expense_type, notes, created_at, updated_at
		FROM transactions
		WHERE id = ?;
	`
	row := db.QueryRow(query, id)
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
func UpdateTransaction(t models.Transaction, db runner) error {


	if db == nil {
		db = DB
	}

	if t.UpdatedAt == "" {
		t.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
	}

	_, err := db.Exec(`
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
func DeleteTransaction(id int64, db runner) error {


	if db == nil {
		db = DB
	}

	_, err := db.Exec(`DELETE FROM transactions WHERE id = ?`, id)
	return err
}


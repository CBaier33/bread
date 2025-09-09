package persistence

import (
	"bread/backend/models"
)

// InsertBudget inserts a new budget and returns its ID
func InsertBudget(b models.Budget) (int64, error) {
	res, err := DB.Exec(`
        INSERT INTO budgets(name, period_start, period_end, created_at, updated_at)
        VALUES (?, ?, ?, ?, ?)`,
		b.Name,
		b.PeriodStart,
		b.PeriodEnd,
		b.CreatedAt,
		b.UpdatedAt,
	)
	if err != nil {
		return 0, err
	}

	return res.LastInsertId()
}

// GetBudget retrieves a budget by ID
func GetBudget(id int64) (models.Budget, error) {
	row := DB.QueryRow(`
        SELECT id, name, period_start, period_end, created_at, updated_at
        FROM budgets
        WHERE id = ?`, id)

	var b models.Budget
	if err := row.Scan(
		&b.ID,
		&b.Name,
		&b.PeriodStart,
		&b.PeriodEnd,
		&b.CreatedAt,
		&b.UpdatedAt,
	); err != nil {
		return b, err
	}
	return b, nil
}

// ListBudgets retrieves all budgets
func ListBudgets() ([]models.Budget, error) {
	rows, err := DB.Query(`
        SELECT id, name, period_start, period_end, created_at, updated_at
        FROM budgets
        ORDER BY period_start DESC
    `)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var budgets []models.Budget
	for rows.Next() {
		var b models.Budget
		if err := rows.Scan(
			&b.ID,
			&b.Name,
			&b.PeriodStart,
			&b.PeriodEnd,
			&b.CreatedAt,
			&b.UpdatedAt,
		); err != nil {
			return nil, err
		}
		budgets = append(budgets, b)
	}

	return budgets, nil
}

// UpdateBudget updates a budget
func UpdateBudget(b models.Budget) error {
	_, err := DB.Exec(`
        UPDATE budgets
        SET name = ?, period_start = ?, period_end = ?, updated_at = ?
        WHERE id = ?`,
		b.Name,
		b.PeriodStart,
		b.PeriodEnd,
		b.UpdatedAt,
		b.ID,
	)
	return err
}

// DeleteBudget deletes a budget by ID
func DeleteBudget(id int64) error {
	_, err := DB.Exec(`DELETE FROM budgets WHERE id = ?`, id)
	return err
}


package persistence

import (
	"bread/backend/models"
)

// Budget Logic

// InsertBudget inserts a new budget and returns its ID
func InsertBudget(b models.Budget) (int64, error) {
	res, err := DB.Exec(`
        INSERT INTO budgets(project_id, name, period_start, period_end, expected_income, starting_balance)
        VALUES (?, ?, ?, ?, ?, ?)`,
		b.ProjectID,
		b.Name,
		b.PeriodStart,
		b.PeriodEnd,
		b.ExpectedIncome,
		b.StartingBalance,
	)
	if err != nil {
		return 0, err
	}

	return res.LastInsertId()
}

// GetBudget retrieves a budget by ID
func GetBudget(id int64) (models.Budget, error) {
	row := DB.QueryRow(`
        SELECT id, project_id, name, period_start, period_end, expected_income, starting_balance, created_at, updated_at
        FROM budgets
        WHERE id = ?`, id)

	var b models.Budget
	if err := row.Scan(
		&b.ID,
		&b.ProjectID,
		&b.Name,
		&b.PeriodStart,
		&b.PeriodEnd,
		&b.ExpectedIncome,
		&b.StartingBalance,
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
        SELECT id, project_id, name, period_start, period_end, expected_income, starting_balance, created_at, updated_at
        FROM budgets
        ORDER BY id DESC
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
			&b.ProjectID,
			&b.Name,
			&b.PeriodStart,
			&b.PeriodEnd,
			&b.ExpectedIncome,
			&b.StartingBalance,
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
        SET project_id = ?, name = ?, period_start = ?, period_end = ?, expected_income = ?, starting_balance = ?, updated_at = (datetime('now'))
        WHERE id = ?`,
		b.ProjectID,
		b.Name,
		b.PeriodStart,
		b.PeriodEnd,
		b.ExpectedIncome,
		b.StartingBalance,
		b.ID,
	)
	return err
}

// DeleteBudget deletes a budget by ID
func DeleteBudget(id int64) error {
	_, err := DB.Exec(`DELETE FROM budgets WHERE id = ?`, id)
	return err
}

// Budget Allocation Logic

func InsertAllocation(b models.BudgetAllocation) (int64, error) {
	res, err := DB.Exec(`
        INSERT INTO budget_allocations(budget_id, category_id, expected_cost)
        VALUES (?, ?, ?)`,
		b.BudgetID,
		b.CategoryID,
		b.ExpectedCost,
	)
	if err != nil {
		return 0, err
	}

	return res.LastInsertId()
}

func GetAllocation(id int64) (models.BudgetAllocation, error) {
	row := DB.QueryRow(`
        SELECT id, budget_id, category_id, expected_cost, created_at, updated_at
        FROM budget_allocations
        WHERE id = ?`, id)

	var b models.BudgetAllocation
	if err := row.Scan(
		&b.ID,
		&b.BudgetID,
		&b.CategoryID,
		&b.ExpectedCost,
		&b.CreatedAt,
		&b.UpdatedAt,
	); err != nil {
		return b, err
	}
	return b, nil
}

func ListAllocations() ([]models.BudgetAllocation, error) {
	rows, err := DB.Query(`
        SELECT id, budget_id, category_id, expected_cost, created_at, updated_at
        FROM budget_allocations
        ORDER BY created_at DESC
    `)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var allocations []models.BudgetAllocation
	for rows.Next() {
		var b models.BudgetAllocation
		if err := rows.Scan(
			&b.ID,
			&b.BudgetID,
			&b.CategoryID,
			&b.ExpectedCost,
			&b.CreatedAt,
			&b.UpdatedAt,
		); err != nil {
			return nil, err
		}
		allocations = append(allocations, b)
	}

	return allocations, nil
}

func UpdateAllocation(b models.BudgetAllocation) error {
	_, err := DB.Exec(`
        UPDATE budget_allocations
        SET budget_id = ?, category_id = ?, expected_cost = ?, updated_at = (datetime('now'))
        WHERE id = ?`,
		b.BudgetID,
		b.CategoryID,
		b.ExpectedCost,
		b.ID,
	)
	return err
}

func DeleteAllocation(id int64) error {
	_, err := DB.Exec(`DELETE FROM budget_allocations WHERE id = ?`, id)
	return err
}

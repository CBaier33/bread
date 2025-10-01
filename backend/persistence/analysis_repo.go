package persistence

func BudgetTotalProjectedCost(budgetID int64, db runner) (int64, error) {

	if db == nil {
		db = DB
	}

	var total int64

	err := db.QueryRow(`
        SELECT COALESCE(SUM(expected_cost), 0)
        FROM budget_allocations
        WHERE budget_id = ?
    `, budgetID).Scan(&total)

	if err != nil {
		return 0, err
	}

	return total, nil
}

func BudgetActualCost(budgetID int64, db runner) (int64, error) {

	if db == nil {
		db = DB
	}

	budget, err := GetBudget(budgetID, db)
	
	if err != nil {
		return 0, err
	}

	var total int64

	err = db.QueryRow(`
		SELECT COALESCE(SUM(amount), 0)
		FROM transactions
		WHERE date BETWEEN ? AND ?; 
		`, budget.PeriodStart, budget.PeriodEnd).Scan(&total)

	if err != nil {
		return 0, err
	}

	return total, nil
}

func AllocationCost(budgetID int64, categoryID *int64, db runner) (int64, error) {

	var (
		query string
		accountIDs []interface{}

	)

	if db == nil {
		db = DB
	}


	budget, err := GetBudget(budgetID, db)
	if err != nil {
		return 0, err
	}

	// Logic allows for evaluating the cost of noncategorized transactions in the budget
	if categoryID == nil {
		accountIDs = []interface{}{budget.PeriodStart, budget.PeriodEnd, budget.ProjectID}
		query = `
		SELECT COALESCE(SUM(t.amount), 0)
		FROM transactions t
		WHERE t.date BETWEEN ? AND ? AND t.project_id = ? AND t.category_id IS NULL)
		`
	} else {
		accountIDs = []interface{}{budget.PeriodStart, budget.PeriodEnd, budget.ProjectID, categoryID}
		query = `
		SELECT COALESCE(SUM(t.amount), 0)
		FROM transactions t
		WHERE t.date BETWEEN ? AND ? AND t.project_id = ? AND t.category_id = ?
		`
	}

	var total int64

	err = db.QueryRow(query, accountIDs...).Scan(&total)

	if err != nil {
		return 0, err
	}

	return total, nil
	
}

package models

type Budget struct {
	ID              int64  `json:"id"`
	ProjectID       int64  `json:"project_id"`
	Name            string `json:"name"`
	PeriodStart     string `json:"period_start"`
	PeriodEnd       string `json:"period_end"`
	ExpectedIncome  int64  `json:"expected_income"`
	StartingBalance int64  `json:"starting_balance"`
	CreatedAt       string `json:"created_at"`
	UpdatedAt       string `json:"updated_at"`
}

type BudgetAllocation struct {
	ID           int64  `json:"id"`
	BudgetID     int64  `json:"budget_id"`
	CategoryID   int64  `json:"category_id"`
	ExpectedCost int64  `json:"expected_cost"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
}

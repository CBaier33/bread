package models

type Transaction struct {
	ID           int64     `json:"id"`
	Description  string    `json:"description"`
	BudgetID     int64     `json:"budget_id"`
	GroupID      *int64    `json:"group_id"`
	CategoryID   *int64    `json:"category_id"`
	CategoryName string    `json:"category_name"`
	Date         string    `json:"date"`
	Amount       int64     `json:"amount"`
	Tags 				 string    `json:"tags"`
	Notes        string    `json:"notes"`
	CreatedAt    string    `json:"created_at"`
	UpdatedAt    string    `json:"updated_at"`
}


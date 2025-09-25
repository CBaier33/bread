package models

type Category struct {
	ID          int64  `json:"id"`
	GroupID     *int64 `json:"group_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	ExpenseType bool   `json:"expense_type"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

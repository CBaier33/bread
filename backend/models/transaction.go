package models

type Transaction struct {
	ID          int64  `json:"id"`
	Description string `json:"description"`
	ProjectID   int64  `json:"project_id"`
	CategoryID  *int64 `json:"category_id"`
	Date        string `json:"date"`
	Amount      int64  `json:"amount"`
	ExpenseType bool   `json:"expense_type"`
	Notes       string `json:"notes"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

package models

type Category struct {
    ID          int64   `json:"id"`
		BudgetID    int64   `json:"budget_id"`
    GroupID     *int64  `json:"group_id"`
	  Name 				string 	`json:"name"`
    Description string  `json:"description"`
	  IsDeposit   bool    `json:"is_deposit"`
		Expected    int64   `json:"expected"`
		Actual      int64   `json:"actual"`
		CreatedAt   string  `json:"created_at"`
		UpdatedAt   string  `json:"updated_at"`
}

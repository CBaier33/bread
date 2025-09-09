package models

type Category struct {
    ID          int64   `json:"id"`
		BudgetID    int64   `json:"budget_id"`
    GroupID     *int64  `json:"group_id"`
    Description string  `json:"description"`
	  Name 				string 	`json:"name"`
	  IsDeposit   bool    `json:"is_deposit"`
		Expected    int64   `json:"amt_expected"`
		Actual      int64   `json:"amt_actual"`
		CreatedAt   string  `json:"created_at"`
		UpdatedAt   string  `json:"updated_at"`
}

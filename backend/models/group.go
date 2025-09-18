package models

type Group struct {
    ID          int64   `json:"id"`
    BudgetID    int64   `json:"budget_id"`
	  Name 				string 	`json:"name"`
    Description string  `json:"description"`
		CreatedAt   string  `json:"created_at"`
		UpdatedAt   string  `json:"updated_at"`
}

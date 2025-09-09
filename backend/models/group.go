package models

type Group struct {
    ID          int64   `json:"id"`
    BudgetID    int64   `json:"budget_id"`
    Description string  `json:"description"`
	  Name 				string 	`json:"name"`
		CreatedAt   string  `json:"created_at"`
		UpdatedAt   string  `json:"updated_at"`
}

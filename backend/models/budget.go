package models

type Budget struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	PeriodStart string    `json:"period_start"`
	PeriodEnd   string 		`json:"period_end"`
	CreatedAt   string    `json:"created_at"`
	UpdatedAt   string    `json:"updated_at"`
}

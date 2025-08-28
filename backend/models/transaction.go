package models

type Transaction struct {
    ID          int64   `json:"id"`
    Description string  `json:"description"`
    Amount      float64 `json:"amount"`
		Direction		string  `json:"direction"`
    Category    string  `json:"category"`
    CreatedAt   string  `json:"created_at"` // ISO string for Wails
}


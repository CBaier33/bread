package models

type Tag struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type TransactionTag struct {
	TransactionID int64  `json:"transaction_id"`
	TagID         int64  `json:"tag_id"`
	CreatedAt     string `json:"created_at"`
}

package models

type Group struct {
    ID          int64   `json:"id"`
    ProjectID   int64   `json:"project_id"`
	  Name 				string 	`json:"name"`
    Description string  `json:"description"`
		CreatedAt   string  `json:"created_at"`
		UpdatedAt   string  `json:"updated_at"`
}

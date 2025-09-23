package persistence

import (
	"bread/backend/models"
)

// InsertProject inserts a new project and returns its ID
func InsertProject(b models.Project) (int64, error) {
	res, err := DB.Exec(`
        INSERT INTO projects(name, description, currency, created_at, updated_at)
        VALUES (?, ?, ?, ?, ?)`,
		b.Name,
		b.Description,
		b.Currency,
		b.CreatedAt,
		b.UpdatedAt,
	)
	if err != nil {
		return 0, err
	}

	return res.LastInsertId()
}

// GetProject retrieves a project by ID
func GetProject(id int64) (models.Project, error) {
	row := DB.QueryRow(`
        SELECT id, name, description, currency, created_at, updated_at
        FROM projects
        WHERE id = ?`, id)

	var b models.Project
	if err := row.Scan(
		&b.ID,
		&b.Name,
		&b.Description,
		&b.Currency,
		&b.CreatedAt,
		&b.UpdatedAt,
	); err != nil {
		return b, err
	}
	return b, nil
}

// ListProjects retrieves all projects
func ListProjects() ([]models.Project, error) {
	rows, err := DB.Query(`
        SELECT id, name, description, currency, created_at, updated_at
        FROM projects
        ORDER BY description DESC
    `)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var projects []models.Project
	for rows.Next() {
		var b models.Project
		if err := rows.Scan(
			&b.ID,
			&b.Name,
			&b.Description,
			&b.Currency,
			&b.CreatedAt,
			&b.UpdatedAt,
		); err != nil {
			return nil, err
		}
		projects = append(projects, b)
	}

	return projects, nil
}

// UpdateProject updates a project
func UpdateProject(b models.Project) error {
	_, err := DB.Exec(`
        UPDATE projects
        SET name = ?, description = ?, currency = ?, updated_at = ?
        WHERE id = ?`,
		b.Name,
		b.Description,
		b.Currency,
		b.UpdatedAt,
		b.ID,
	)
	return err
}

// DeleteProject deletes a project by ID
func DeleteProject(id int64) error {
	_, err := DB.Exec(`DELETE FROM projects WHERE id = ?`, id)
	return err
}


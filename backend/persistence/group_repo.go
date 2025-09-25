package persistence

import (
	"bread/backend/models"
)

// InsertGroup inserts a new group into the database and returns its ID.
func InsertGroup(g models.Group) (int64, error) {
	res, err := DB.Exec(`
		INSERT INTO groups(project_id, name, description)
		VALUES (?, ?, ?)`,
		g.ProjectID,
		g.Name,
		g.Description,
	)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

// GetGroup retrieves a group by ID.
func GetGroup(id int64) (models.Group, error) {
	row := DB.QueryRow(`
		SELECT id, project_id, name, description, created_at, updated_at
		FROM groups
		WHERE id = ?`, id)

	var g models.Group
	if err := row.Scan(
		&g.ID,
		&g.ProjectID,
		&g.Name,
		&g.Description,
		&g.CreatedAt,
		&g.UpdatedAt,
	); err != nil {
		return g, err
	}
	return g, nil
}

// ListGroups lists all groups.
func ListGroups() ([]models.Group, error) {
	rows, err := DB.Query(`
		SELECT id, project_id, name, description, created_at, updated_at
		FROM groups
		ORDER BY id`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var groups []models.Group
	for rows.Next() {
		var g models.Group
		if err := rows.Scan(
			&g.ID,
			&g.ProjectID,
			&g.Name,
			&g.Description,
			&g.CreatedAt,
			&g.UpdatedAt,
		); err != nil {
			return nil, err
		}
		groups = append(groups, g)
	}
	return groups, nil
}

// UpdateGroup updates a group.
func UpdateGroup(g models.Group) error {
	_, err := DB.Exec(`
		UPDATE groups
		SET name = ?, description = ?, updated_at = (datetime('now'))
		WHERE id = ?`,
		g.Name,
		g.Description,
		g.ID,
	)
	return err
}

// DeleteGroup deletes a group.
func DeleteGroup(id int64) error {
	_, err := DB.Exec(`DELETE FROM groups WHERE id = ?`, id)
	return err
}


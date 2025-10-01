package persistence

import (
	"bread/backend/models"
)

// InsertGroup inserts a new group into the database and returns its ID.
func InsertGroup(g models.Group, db runner) (int64, error) {


	if db == nil {
		db = DB
	}

	res, err := db.Exec(`
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
func GetGroup(id int64, db runner) (models.Group, error) {


	if db == nil {
		db = DB
	}

	row := db.QueryRow(`
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
func ListGroups(projectID int64, db runner) ([]models.Group, error) {


	if db == nil {
		db = DB
	}

	rows, err := db.Query(`
		SELECT id, project_id, name, description, created_at, updated_at
		FROM groups
		WHERE project_id = ?
		ORDER BY id`, projectID)
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
func UpdateGroup(g models.Group, db runner) error {


	if db == nil {
		db = DB
	}

	_, err := db.Exec(`
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
func DeleteGroup(id int64, db runner) error {


	if db == nil {
		db = DB
	}

	_, err := db.Exec(`DELETE FROM groups WHERE id = ?`, id)
	return err
}


package persistence

import (
	"bread/backend/models"
)

// InsertTag inserts a new tag into the database and returns its ID.
func InsertTag(g models.Tag) (int64, error) {
	res, err := DB.Exec(`
		INSERT INTO tags(project_id, name)
		VALUES (?, ?)`,
		g.ProjectID,
		g.Name,
	)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

// GetTag retrieves a tag by ID.
func GetTag(id int64) (*models.Tag, error) {
	row := DB.QueryRow(`
		SELECT id, project_id, name 
		FROM tags
		WHERE id = ?`, id)

	var g models.Tag
	if err := row.Scan(
		&g.ID,
		&g.ProjectID,
		&g.Name,
	); err != nil {
		return nil, err
	}
	return &g, nil
}

// ListTags lists all tags.
func ListTags() ([]models.Tag, error) {
	rows, err := DB.Query(`
		SELECT id, project_id, name, created_at, updated_at
		FROM tags
		ORDER BY id`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tags []models.Tag
	for rows.Next() {
		var g models.Tag
		if err := rows.Scan(
			&g.ID,
			&g.ProjectID,
			&g.Name,
			&g.CreatedAt,
			&g.UpdatedAt,
		); err != nil {
			return nil, err
		}
		tags = append(tags, g)
	}
	return tags, nil
}

// UpdateTag updates a tag.
func UpdateTag(g models.Tag) error {
	_, err := DB.Exec(`
		UPDATE tags
		SET project_id = ?, name = ?, updated_at = (datetime('now'))
		WHERE id = ?`,
		g.ProjectID,
		g.Name,
		g.ID,
	)
	return err
}

// DeleteTag deletes a tag.
func DeleteTag(id int64) error {
	_, err := DB.Exec(`DELETE FROM tags WHERE id = ?`, id)
	return err
}

// InsertTransactionTag inserts a new tag into the database and returns its ID.
func InsertTransactionTag(g models.TransactionTag) (error) {
	_, err := DB.Exec(`
		INSERT INTO transaction_tags(transaction_id, tag_id)
		VALUES (?, ?)`,
		g.TransactionID,
		g.TagID,
	)
	return err
}

// GetTransactionTag retrieves a tag by ID.
func GetTransactionTag(transaction_id int64, tag_id int64) (*models.TransactionTag, error) {
	row := DB.QueryRow(`
		SELECT created_at
		FROM transaction_tags
		WHERE transaction_id = ? and tag_id = ?`, transaction_id, tag_id)

	var g models.TransactionTag
	if err := row.Scan(
		&g.CreatedAt,
	); err != nil {
		return nil, err
	}
	return &g, nil
}

// DeleteTransactionTag deletes a tag.
func DeleteTransactionTag(transaction_id int64, tag_id int64) error {
	_, err := DB.Exec(`DELETE FROM transaction_tags WHERE transaction_id = ? and tag_id = ?`, transaction_id, tag_id)
	return err
}

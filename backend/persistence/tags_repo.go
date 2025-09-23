package persistence

import (
	"bread/backend/models"
)

// InsertTag inserts a new group into the database and returns its ID.
func InsertTag(g models.Tag) (int64, error) {
	res, err := DB.Exec(`
		INSERT INTO tags(name, created_at, updated_at)
		VALUES (?, ?, ?)`,
		g.Name,
		g.CreatedAt,
		g.UpdatedAt,
	)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

// GetTag retrieves a group by ID.
func GetTag(id int64) (*models.Tag, error) {
	row := DB.QueryRow(`
		SELECT id, name, created_at, updated_at
		FROM tags
		WHERE id = ?`, id)

	var g models.Tag
	if err := row.Scan(
		&g.ID,
		&g.Name,
		&g.CreatedAt,
		&g.UpdatedAt,
	); err != nil {
		return nil, err
	}
	return &g, nil
}

// ListTags lists all tags.
func ListTags() ([]models.Tag, error) {
	rows, err := DB.Query(`
		SELECT id, name, created_at, updated_at
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

// UpdateTag updates a group.
func UpdateTag(g models.Tag) error {
	_, err := DB.Exec(`
		UPDATE tags
		SET name = ?, updated_at = ?
		WHERE id = ?`,
		g.Name,
		g.UpdatedAt,
		g.ID,
	)
	return err
}

// DeleteTag deletes a group.
func DeleteTag(id int64) error {
	_, err := DB.Exec(`DELETE FROM tags WHERE id = ?`, id)
	return err
}

// InsertTransactionTag inserts a new group into the database and returns its ID.
func InsertTransactionTag(g models.TransactionTag) (int64, error) {
	res, err := DB.Exec(`
		INSERT INTO transaction_tag(transaction_id, tag_id, created_at)
		VALUES (?, ?, ?)`,
		g.TransactionID,
		g.TagID,
		g.CreatedAt,
	)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

// GetTransactionTag retrieves a group by ID.
func GetTransactionTag(transaction_id int64, tag_id int64) (*models.TransactionTag, error) {
	row := DB.QueryRow(`
		SELECT created_at
		FROM transaction_tag
		WHERE transaction_id = ? and tag_id = ?`, transaction_id, tag_id)

	var g models.TransactionTag
	if err := row.Scan(
		&g.CreatedAt,
	); err != nil {
		return nil, err
	}
	return &g, nil
}

// DeleteTransactionTag deletes a group.
func DeleteTransactionTag(transaction_id int64, tag_id int64) error {
	_, err := DB.Exec(`DELETE FROM transaction_tag WHERE transaction_id = ? and tag_id = ?`, transaction_id, tag_id)
	return err
}

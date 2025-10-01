package persistence

import (
	"bread/backend/models"
)

// InsertTag inserts a new tag into the database and returns its ID.
func InsertTag(g models.Tag, db runner) (int64, error) {


	if db == nil {
		db = DB
	}

	res, err := db.Exec(`
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
func GetTag(id int64, db runner) (*models.Tag, error) {


	if db == nil {
		db = DB
	}

	row := db.QueryRow(`
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
func ListTags(projectID int64, db runner) ([]models.Tag, error) {


	if db == nil {
		db = DB
	}

	rows, err := db.Query(`
		SELECT id, project_id, name, created_at, updated_at
		FROM tags
		WHERE project_id = ?
		ORDER BY id`, projectID)
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
func UpdateTag(g models.Tag, db runner) error {


	if db == nil {
		db = DB
	}

	_, err := db.Exec(`
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
func DeleteTag(id int64, db runner) error {


	if db == nil {
		db = DB
	}

	_, err := db.Exec(`DELETE FROM tags WHERE id = ?`, id)
	return err
}

// InsertTransactionTag inserts a new tag into the database and returns its ID.
func InsertTransactionTag(g models.TransactionTag, db runner) (error) {


	if db == nil {
		db = DB
	}

	_, err := db.Exec(`
		INSERT INTO transaction_tags(transaction_id, tag_id)
		VALUES (?, ?)`,
		g.TransactionID,
		g.TagID,
	)
	return err
}



// GetTransactionTag retrieves a tag by ID.
func GetTransactionTag(transaction_id int64, tag_id int64, db runner) (*models.TransactionTag, error) {


	if db == nil {
		db = DB
	}

	row := db.QueryRow(`
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
func DeleteTransactionTag(transaction_id int64, tag_id int64, db runner) error {


	if db == nil {
		db = DB
	}

	_, err := db.Exec(`DELETE FROM transaction_tags WHERE transaction_id = ? and tag_id = ?`, transaction_id, tag_id)
	return err
}

// Return a slice of all tags given to a certain transaction
func GetTags(transactionID int64, db runner) ([]models.Tag, error) {

	if db == nil {
		db = DB
	}

	rows, err := db.Query(`
		SELECT t.id, t.name, t.created_at, t.updated_at
		FROM tags t
		RIGHT JOIN transaction_tags tt ON t.id = tt.tag_id AND tt.transaction_id = ?
		ORDER BY id`, transactionID)
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


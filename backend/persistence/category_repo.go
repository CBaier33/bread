package persistence

import (
	"bread/backend/models"
)

// InsertCategory inserts a new category into the database and returns its ID.
func InsertCategory(c models.Category) (int64, error) {
	res, err := DB.Exec(`
		INSERT INTO categories(group_id, name, description, expense_type)
		VALUES (?, ?, ?, ?)`,
		c.GroupID,
		c.Name,
		c.Description,
		c.ExpenseType,
	)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

// GetCategory retrieves a category by ID.
func GetCategory(id int64) (models.Category, error) {
	row := DB.QueryRow(`
		SELECT id, group_id, name, description, expense_type, created_at, updated_at
		FROM categories
		WHERE id = ?`, id)

	var c models.Category
	if err := row.Scan(
		&c.ID,
		&c.GroupID,
		&c.Name,
		&c.Description,
		&c.ExpenseType,
		&c.CreatedAt,
		&c.UpdatedAt,
	); err != nil {
		return c, err
	}
	return c, nil
}

// ListCategories lists all categories.
func ListCategories() ([]models.Category, error) {
	rows, err := DB.Query(`
		SELECT id, group_id, name, description, expense_type, created_at, updated_at
		FROM categories
		ORDER BY id`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []models.Category
	for rows.Next() {
		var c models.Category
		if err := rows.Scan(
			&c.ID,
			&c.GroupID,
			&c.Name,
			&c.Description,
			&c.ExpenseType,
			&c.CreatedAt,
			&c.UpdatedAt,
		); err != nil {
			return nil, err
		}
		categories = append(categories, c)
	}
	return categories, nil
}

// UpdateCategory updates a category.
func UpdateCategory(c models.Category) error {
	_, err := DB.Exec(`
		UPDATE categories
		SET group_id = ?, name = ?, description = ?, expense_type = ?, updated_at = (datetime('now'))
		WHERE id = ?`,
		c.GroupID,
		c.Name,
		c.Description,
		c.ExpenseType,
		c.ID,
	)
	return err
}

// DeleteCategory deletes a category.
func DeleteCategory(id int64) error {
	_, err := DB.Exec(`DELETE FROM categories WHERE id = ?`, id)
	return err
}

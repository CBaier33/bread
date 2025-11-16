package persistence

import (
	"bread/backend/models"
)

// InsertCategory inserts a new category into the database and returns its ID.
func InsertCategory(c models.Category, db runner) (int64, error) {


	if db == nil {
		db = DB
	}

	res, err := db.Exec(`
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
func GetCategory(id int64, db runner) (models.Category, error) {


	if db == nil {
		db = DB
	}

	row := db.QueryRow(`
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

func ListProjectCategories(projectID int64, db runner) ([]models.Category, error) {

	if db == nil {
		db = DB
	}
	rows, err := db.Query(`
	 SELECT c.id, c.group_id, c.name, c.description, c.expense_type, c.created_at, c.updated_at
   FROM categories c
   JOIN groups g on c.group_id = g.id
   WHERE g.project_id = ?`, projectID)

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

// ListCategories lists all categories.
func ListCategories(groupID int64, db runner) ([]models.Category, error) {

	if db == nil {
		db = DB
	}

	rows, err := db.Query(`
		SELECT id, group_id, name, description, expense_type, created_at, updated_at
		FROM categories
		WHERE group_id = ?
		ORDER BY id`, groupID)
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
func UpdateCategory(c models.Category, db runner) error {


	if db == nil {
		db = DB
	}

	_, err := db.Exec(`
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
func DeleteCategory(id int64, db runner) error {


	if db == nil {
		db = DB
	}

	_, err := db.Exec(`DELETE FROM categories WHERE id = ?`, id)
	return err
}

package controllers

import (
	"bread/backend/models"
	"bread/backend/services"
)

// CategoryController exposes category operations to the frontend via Wails
type CategoryController struct{}

// NewCategoryController creates a new controller instance
func NewCategoryController() *CategoryController {
	return &CategoryController{}
}

// AddCategory creates a new category
func (c *CategoryController) AddCategory(budgetID int64, groupID int64, name, description string, expected int64, actual int64) (models.Category, error) {
	return services.CreateCategory(budgetID, groupID, name, description, expected, actual)
}

// GetCategory retrieves a category by ID
func (c *CategoryController) GetCategory(id int64) (models.Category, error) {
	return services.GetCategoryByID(id)
}

// ListCategories returns all categories
func (c *CategoryController) ListCategories() ([]models.Category, error) {
	return services.ListCategories()
}

// UpdateCategory updates an existing category
func (c *CategoryController) UpdateCategory(category models.Category) error {
	return services.UpdateCategory(category)
}

// DeleteCategory deletes a category by ID
func (c *CategoryController) DeleteCategory(id int64) error {
	return services.DeleteCategory(id)
}


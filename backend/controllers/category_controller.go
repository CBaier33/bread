package controllers

import (
	"bread/backend/models"
	"bread/backend/services"
)

type CategoryController struct {
	service *services.CategoryService
}

func NewCategoryController() *CategoryController {
	return &CategoryController{
		service: &services.CategoryService{},
	}
}

// CreateCategory inserts a new category and returns it
func (c *CategoryController) CreateCategory(groupID int64, name, description string, expenseType bool) (models.Category, error) {
	return c.service.CreateCategory(groupID, name, description, expenseType)
}

// GetCategoryByID retrieves a category by ID
func (c *CategoryController) GetCategoryByID(id int64) (models.Category, error) {
	return c.service.GetCategoryByID(id)
}

// ListCategories returns all categories for a group
func (c *CategoryController) ListCategories(groupID int64) ([]models.Category, error) {
	return c.service.ListCategories(groupID)
}

// UpdateCategory updates a category's data
func (c *CategoryController) UpdateCategory(cat models.Category) error {
	return c.service.UpdateCategory(cat)
}

// DeleteCategory removes a category by ID
func (c *CategoryController) DeleteCategory(id int64) error {
	return c.service.DeleteCategory(id)
}


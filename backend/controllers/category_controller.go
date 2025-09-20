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

// CreateCategory creates a new category
func (c *CategoryController) CreateCategory(budgetID int64, groupID int64, name string, description string, expected int64, actual int64) (models.Category, error) {
	return c.service.CreateCategory(budgetID, groupID, name, description, expected, actual)
}

// GetCategory retrieves a category by ID
func (c *CategoryController) GetCategory(id int64) (models.Category, error) {
	return c.service.GetCategoryByID(id)
}

// ListCategories returns all categories
func (c *CategoryController) ListCategories() ([]models.Category, error) {
	return c.service.ListCategories()
}

// UpdateCategory updates an existing category
func (c *CategoryController) UpdateCategory(category models.Category) error {
	return c.service.UpdateCategory(category)
}

// DeleteCategory deletes a category by ID
func (c *CategoryController) DeleteCategory(id int64) error {
	return c.service.DeleteCategory(id)
}


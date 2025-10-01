package services

import (
	"bread/backend/models"
	"bread/backend/persistence"
	"fmt"
)

type CategoryService struct{}

func (s *CategoryService) CreateCategory(groupID int64, name, description string, expenseType bool) (models.Category, error) {
	c := models.Category{
		GroupID:     &groupID,
		Name:        name,
		Description: description,
		ExpenseType: expenseType,
	}

	id, err := persistence.InsertCategory(c, nil)
	if err != nil {
		return c, fmt.Errorf("CreateCategory: %w", err)
	}
	c.ID = id
	return c, nil
}

func (s *CategoryService) GetCategoryByID(id int64) (models.Category, error) {
	c, err := persistence.GetCategory(id, nil)
	if err != nil {
		return models.Category{}, fmt.Errorf("GetCategoryByID: %w", err)
	}
	return c, nil
}

func (s *CategoryService) ListCategories(groupID int64) ([]models.Category, error) {
	return persistence.ListCategories(groupID, nil)
}

func (s *CategoryService) UpdateCategory(c models.Category) error {
	return persistence.UpdateCategory(c, nil)
}

func (s *CategoryService) DeleteCategory(id int64) error {
	return persistence.DeleteCategory(id, nil)
}


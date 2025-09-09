package services

import (
	"bread/backend/models"
	"bread/backend/persistence"
	"fmt"
	"time"
)

func CreateCategory(budgetID int64, groupID int64, name, description string) (models.Category, error) {
	now := time.Now().Format("2006-01-02 15:04:05")
	c := models.Category{
		BudgetID: budgetID,
		GroupID:     &groupID,
		Name:        name,
		Description: description,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	id, err := persistence.InsertCategory(c)
	if err != nil {
		return c, fmt.Errorf("CreateCategory: %w", err)
	}
	c.ID = id
	return c, nil
}

func GetCategoryByID(id int64) (models.Category, error) {
	c, err := persistence.GetCategory(id)
	if err != nil {
		return models.Category{}, fmt.Errorf("GetCategoryByID: %w", err)
	}
	return *c, nil
}

func ListCategories() ([]models.Category, error) {
	return persistence.ListCategories()
}

func UpdateCategory(c models.Category) error {
	if c.UpdatedAt == "" {
		c.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
	}
	return persistence.UpdateCategory(c)
}

func DeleteCategory(id int64) error {
	return persistence.DeleteCategory(id)
}


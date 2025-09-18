package services

import (
	"bread/backend/models"
	"bread/backend/persistence"
	"fmt"
	"time"
)

type GroupService struct{}

func (s *GroupService) CreateGroup(budgetID int64, name string, description string) (models.Group, error) {
	now := time.Now().Format("2006-01-02 15:04:05")
	g := models.Group{
		BudgetID: budgetID,
		Name:        name,
		Description: description,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	id, err := persistence.InsertGroup(g)
	if err != nil {
		return g, fmt.Errorf("CreateGroup: %w", err)
	}
	g.ID = id
	return g, nil
}

func (s *GroupService) GetGroup(id int64) (models.Group, error) {
	g, err := persistence.GetGroup(id)
	if err != nil {
		return models.Group{}, fmt.Errorf("GetGroupByID: %w", err)
	}
	return *g, nil
}

func (s *GroupService) ListGroups() ([]models.Group, error) {
	return persistence.ListGroups()
}

func (s *GroupService) UpdateGroup(g models.Group) error {
	if g.UpdatedAt == "" {
		g.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
	}
	return persistence.UpdateGroup(g)
}

func (s *GroupService) DeleteGroup(id int64) error {
	return persistence.DeleteGroup(id)
}


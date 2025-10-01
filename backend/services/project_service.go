package services

import (
	"bread/backend/models"
	"bread/backend/persistence"
	"fmt"
)

type ProjectService struct{}

func (s *ProjectService) CreateProject(name, description, currency string) (models.Project, error) {
	b := models.Project{
		Name:        name,
		Description: description,
		Currency:    currency,
	}

	id, err := persistence.InsertProject(b, nil)
	if err != nil {
		return b, fmt.Errorf("CreateProject: %w", err)
	}
	b.ID = id
	return b, nil
}

func (s *ProjectService) GetProject(id int64) (models.Project, error) {
	b, err := persistence.GetProject(id, nil)
	if err != nil {
		return b, fmt.Errorf("GetProject: %w", err)
	}
	return b, nil
}

func (s *ProjectService) ListProjects() ([]models.Project, error) {
	budgets, err := persistence.ListProjects(nil)
	if err != nil {
		return nil, fmt.Errorf("ListProjects: %w", err)
	}
	return budgets, nil
}

// UpdateProject updates a budget's name or period
func (s *ProjectService) UpdateProject(b models.Project) error {
	if err := persistence.UpdateProject(b, nil); err != nil {
		return fmt.Errorf("UpdateProject: %w", err)
	}
	return nil
}

func (s *ProjectService) DeleteProject(id int64) error {
	if err := persistence.DeleteProject(id, nil); err != nil {
		return fmt.Errorf("DeleteProject: %w", err)
	}
	return nil
}


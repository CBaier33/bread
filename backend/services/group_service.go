package services

import (
	"bread/backend/models"
	"bread/backend/persistence"
	"fmt"
)

type GroupService struct{}

// CreateGroup inserts a new group and returns the full group with ID populated
func (s *GroupService) CreateGroup(periodID int64, name, description string) (models.Group, error) {

	b := models.Group{
		ProjectID:   periodID,
		Name:        name,
		Description: description,
	}

	id, err := persistence.InsertGroup(b, nil)
	if err != nil {
		return b, fmt.Errorf("CreateGroup: %w", err)
	}
	b.ID = id
	return b, nil
}

// GetGroup retrieves a group by ID
func (s *GroupService) GetGroup(id int64) (models.Group, error) {
	b, err := persistence.GetGroup(id, nil)
	if err != nil {
		return b, fmt.Errorf("GetGroup: %w", err)
	}
	return b, nil
}

// ListGroups returns all groups
func (s *GroupService) ListGroups(projectID int64) ([]models.Group, error) {
	groups, err := persistence.ListGroups(projectID, nil)
	if err != nil {
		return nil, fmt.Errorf("ListGroups: %w", err)
	}
	return groups, nil
}

// UpdateGroup updates a group's name or period
func (s *GroupService) UpdateGroup(b models.Group) error {
	if err := persistence.UpdateGroup(b, nil); err != nil {
		return fmt.Errorf("UpdateGroup: %w", err)
	}
	return nil
}

// DeleteGroup removes a group by ID
func (s *GroupService) DeleteGroup(id int64) error {
	if err := persistence.DeleteGroup(id, nil); err != nil {
		return fmt.Errorf("DeleteGroup: %w", err)
	}
	return nil
}


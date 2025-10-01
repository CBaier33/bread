package controllers

import (
	"bread/backend/models"
	"bread/backend/services"
)

type GroupController struct {
	service *services.GroupService
}

func NewGroupController() *GroupController {
	return &GroupController{
		service: &services.GroupService{},
	}
}

// CreateGroup inserts a new group and returns it
func (c *GroupController) CreateGroup(periodID int64, name, description string) (models.Group, error) {
	return c.service.CreateGroup(periodID, name, description)
}

// GetGroup retrieves a group by ID
func (c *GroupController) GetGroup(id int64) (models.Group, error) {
	return c.service.GetGroup(id)
}

// ListGroups returns all groups for a project
func (c *GroupController) ListGroups(projectID int64) ([]models.Group, error) {
	return c.service.ListGroups(projectID)
}

// UpdateGroup updates a group's data
func (c *GroupController) UpdateGroup(b models.Group) error {
	return c.service.UpdateGroup(b)
}

// DeleteGroup removes a group by ID
func (c *GroupController) DeleteGroup(id int64) error {
	return c.service.DeleteGroup(id)
}


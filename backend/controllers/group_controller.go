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

func (c *GroupController) CreateGroup(budgetID int64, name, description string) (models.Group, error) {
	return c.service.CreateGroup(budgetID, name, description)
}

// GetGroup retrieves a group by ID
func (c *GroupController) GetGroup(id int64) (models.Group, error) {
	return c.service.GetGroup(id)
}

// ListGroups returns all categories
func (c *GroupController) ListGroups() ([]models.Group, error) {
	return c.service.ListGroups()
}

// UpdateGroup updates an existing group
func (c *GroupController) UpdateGroup(group models.Group) error {
	return c.service.UpdateGroup(group)
}

// DeleteGroup deletes a group by ID
func (c *GroupController) DeleteGroup(id int64) error {
	return c.service.DeleteGroup(id)
}


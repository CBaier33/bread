package controllers

import (
	"bread/backend/models"
	"bread/backend/services"
)

type ProjectController struct {
	service *services.ProjectService
}

func NewProjectController() *ProjectController {
	return &ProjectController{
		service: &services.ProjectService{},
	}
}

// CreateProject inserts a new project and returns it
func (c *ProjectController) CreateProject(name, description, currency string) (models.Project, error) {
	return c.service.CreateProject(name, description, currency)
}

// GetProject retrieves a project by ID
func (c *ProjectController) GetProject(id int64) (models.Project, error) {
	return c.service.GetProject(id)
}

// ListProjects returns all projects
func (c *ProjectController) ListProjects() ([]models.Project, error) {
	return c.service.ListProjects()
}

// UpdateProject updates a project's data
func (c *ProjectController) UpdateProject(p models.Project) error {
	return c.service.UpdateProject(p)
}

// DeleteProject removes a project by ID
func (c *ProjectController) DeleteProject(id int64) error {
	return c.service.DeleteProject(id)
}


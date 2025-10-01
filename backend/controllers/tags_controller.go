package controllers

import (
	"bread/backend/models"
	"bread/backend/services"
)

type TagController struct {
	service *services.TagService
}

func NewTagController() *TagController {
	return &TagController{
		service: &services.TagService{},
	}
}

// --- Tag Management ---

// CreateTag inserts a new tag and returns it
func (c *TagController) CreateTag(projectID int64, name string) (models.Tag, error) {
	return c.service.CreateTag(projectID, name)
}

// GetTag retrieves a tag by ID
func (c *TagController) GetTag(id int64) (*models.Tag, error) {
	return c.service.GetTag(id)
}

// ListTags returns all tags for a project
func (c *TagController) ListTags(projectID int64) ([]models.Tag, error) {
	return c.service.ListTags(projectID)
}

// UpdateTag updates a tag's data
func (c *TagController) UpdateTag(tag models.Tag) error {
	return c.service.UpdateTag(tag)
}

// DeleteTag removes a tag by ID
func (c *TagController) DeleteTag(id int64) error {
	return c.service.DeleteTag(id)
}

// --- Transaction ↔ Tag Linking ---

// CreateTransactionTag links a tag to a transaction
func (c *TagController) CreateTransactionTag(transactionID, tagID int64) (models.TransactionTag, error) {
	return c.service.CreateTransactionTag(transactionID, tagID)
}

// GetTransactionTag retrieves a specific transaction-tag link
func (c *TagController) GetTransactionTag(transactionID, tagID int64) (*models.TransactionTag, error) {
	return c.service.GetTransactionTag(transactionID, tagID)
}

// GetTags retrieves all tags attached to a transaction
func (c *TagController) GetTags(transactionID int64) ([]models.Tag, error) {
	return c.service.GetTags(transactionID)
}

// DeleteTransactionTag removes a tag from a transaction
func (c *TagController) DeleteTransactionTag(transactionID, tagID int64) error {
	return c.service.DeleteTransactionTag(transactionID, tagID)
}


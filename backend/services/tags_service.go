package services

import (
	"bread/backend/models"
	"bread/backend/persistence"
	"fmt"
)

type TagService struct{}

// CreateTag inserts a new tag and returns the full tag with ID populated
func (s *TagService) CreateTag(projectID int64, name string) (models.Tag, error) {

	b := models.Tag{
		ProjectID:   projectID,
		Name:        name,
	}

	id, err := persistence.InsertTag(b, nil)
	if err != nil {
		return b, fmt.Errorf("CreateTag: %w", err)
	}
	b.ID = id
	return b, nil
}

// GetTag retrieves a tag by ID
func (s *TagService) GetTag(id int64) (*models.Tag, error) {
	b, err := persistence.GetTag(id, nil)
	if err != nil {
		return b, fmt.Errorf("GetTag: %w", err)
	}
	return b, nil
}

// ListTags returns all tags
func (s *TagService) ListTags(projectID int64) ([]models.Tag, error) {
	tags, err := persistence.ListTags(projectID, nil)
	if err != nil {
		return nil, fmt.Errorf("ListTags: %w", err)
	}
	return tags, nil
}

// UpdateTag updates a tag's name or period
func (s *TagService) UpdateTag(b models.Tag) error {
	if err := persistence.UpdateTag(b, nil); err != nil {
		return fmt.Errorf("UpdateTag: %w", err)
	}
	return nil
}

// DeleteTag removes a tag by ID
func (s *TagService) DeleteTag(id int64) error {
	if err := persistence.DeleteTag(id, nil); err != nil {
		return fmt.Errorf("DeleteTag: %w", err)
	}
	return nil
}

func (s *TagService) CreateTransactionTag(transactionID, tagID int64) (models.TransactionTag, error) {

	tt := models.TransactionTag{
		TransactionID:  transactionID,
		TagID:          tagID,
	}

	err := persistence.InsertTransactionTag(tt, nil)
	if err != nil {
		return tt, fmt.Errorf("CreateTransactionTag: %w", err)
	}
	return tt, nil
}

func (s *TagService) GetTransactionTag(transactionID, tagID int64) (*models.TransactionTag, error) {
	tt, err := persistence.GetTransactionTag(transactionID, tagID, nil)
	if err != nil {
		return tt, fmt.Errorf("GetTransactionTag: %w", err)
	}
	return tt, nil
}

func (s *TagService) GetTags(transactionID int64) ([]models.Tag, error) {
	tags, err := persistence.GetTags(transactionID, nil)
	if err != nil {
		return tags, fmt.Errorf("GetTransactionTag: %w", err)
	}
	return tags, nil

}

func (s *TagService) DeleteTransactionTag(transactionID, tagID int64) error {
	if err := persistence.DeleteTransactionTag(transactionID, tagID, nil); err != nil {
		return fmt.Errorf("DeleteTransactionTag: %w", err)
	}
	return nil
}

package services

import (
	"bread/backend/models"
	"bread/backend/persistence"
	"fmt"
	"time"
)

type BudgetService struct{}

// CreateBudget inserts a new budget and returns the full budget with ID populated
func (s *BudgetService) CreateBudget(name string, periodStart, periodEnd string) (models.Budget, error) {
	now := time.Now().Format("2006-01-02 15:04:05")
	b := models.Budget{
		Name:        name,
		PeriodStart: periodStart,
		PeriodEnd:   periodEnd,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	id, err := persistence.InsertBudget(b)
	if err != nil {
		return b, fmt.Errorf("CreateBudget: %w", err)
	}
	b.ID = id
	return b, nil
}

// GetBudget retrieves a budget by ID
func (s *BudgetService) GetBudget(id int64) (models.Budget, error) {
	b, err := persistence.GetBudget(id)
	if err != nil {
		return b, fmt.Errorf("GetBudget: %w", err)
	}
	return b, nil
}

// ListBudgets returns all budgets
func (s *BudgetService) ListBudgets() ([]models.Budget, error) {
	budgets, err := persistence.ListBudgets()
	if err != nil {
		return nil, fmt.Errorf("ListBudgets: %w", err)
	}
	return budgets, nil
}

// UpdateBudget updates a budget's name or period
func (s *BudgetService) UpdateBudget(b models.Budget) error {
	b.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
	if err := persistence.UpdateBudget(b); err != nil {
		return fmt.Errorf("UpdateBudget: %w", err)
	}
	return nil
}

// DeleteBudget removes a budget by ID
func (s *BudgetService) DeleteBudget(id int64) error {
	if err := persistence.DeleteBudget(id); err != nil {
		return fmt.Errorf("DeleteBudget: %w", err)
	}
	return nil
}


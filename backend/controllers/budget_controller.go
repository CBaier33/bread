package controllers

import (
	"bread/backend/models"
	"bread/backend/services"
)

type BudgetController struct {
	service *services.BudgetService
}

func NewBudgetController() *BudgetController {
	return &BudgetController{
		service: &services.BudgetService{},
	}
}

// CreateBudget creates a new budget
func (c *BudgetController) CreateBudget(periodID int64, name, periodStart, periodEnd string, expectedIncome, startingBalance int64) (models.Budget, error) {
	return c.service.CreateBudget(periodID, name, periodStart, periodEnd, expectedIncome, startingBalance)
}

// GetBudget retrieves a budget by ID
func (c *BudgetController) GetBudget(id int64) (models.Budget, error) {
	return c.service.GetBudget(id)
}

// ListBudgets returns all budgets for a project
func (c *BudgetController) ListBudgets(projectID int64) ([]models.Budget, error) {
	return c.service.ListBudgets(projectID)
}

// UpdateBudget updates an existing budget
func (c *BudgetController) UpdateBudget(b models.Budget) error {
	return c.service.UpdateBudget(b)
}

// DeleteBudget deletes a budget by ID
func (c *BudgetController) DeleteBudget(id int64) error {
	return c.service.DeleteBudget(id)
}

// AddAllocation creates a new allocation for a budget
func (c *BudgetController) AddAllocation(budgetID, categoryID, expectedCost int64) error {
	return c.service.AddAllocation(budgetID, categoryID, expectedCost)
}

// UpdateAllocationCost updates the expected cost for a budget allocation
func (c *BudgetController) UpdateAllocationCost(budgetID, categoryID, newCost int64) error {
	return c.service.UpdateAllocationCost(budgetID, categoryID, newCost)
}

func (c *BudgetController) DeleteAllocation(id int64) error {
	return c.service.DeleteAllocation(id)
}

func (c *BudgetController) ListAllocations(budgetID int64) ([]models.BudgetAllocation, error) {
	return c.service.ListAllocations(budgetID)
}

// DuplicateBudget duplicates an existing budget and its allocations
func (c *BudgetController) DuplicateBudget(periodID, oldBudgetID int64, name, periodStart, periodEnd string) (int64, error) {
	return c.service.DuplicateBudget(periodID, oldBudgetID, name, periodStart, periodEnd)
}


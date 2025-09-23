package controllers

import (
	"fmt"
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

// AddBudget creates a new budget
func (c *BudgetController) CreateBudget(name string, period_start string, period_end string) (models.Budget, error) {
	fmt.Println("Budgets called")
	fmt.Println(name, period_start, period_end)
	return c.service.CreateBudget(name, period_start, period_end)
}

// GetBudget retrieves a budget by ID
func (c *BudgetController) GetBudget(id int64) (models.Budget, error) {
	return c.service.GetBudget(id)
}

// ListBudgets returns all categories
func (c *BudgetController) ListBudgets() ([]models.Budget, error) {
	return c.service.ListBudgets()
}

// UpdateBudget updates an existing budget
func (c *BudgetController) UpdateBudget(budget models.Budget) error {
	return c.service.UpdateBudget(budget)
}

// DeleteBudget deletes a budget by ID
func (c *BudgetController) DeleteBudget(id int64) error {
	return c.service.DeleteBudget(id)
}


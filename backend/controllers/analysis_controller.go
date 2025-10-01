package controllers

import (
	"bread/backend/services"
)

type AnalysisController struct {
	service *services.AnalysisService
}

func NewAnalysisController() *AnalysisController {
	return &AnalysisController{
		service: &services.AnalysisService{},
	}
}

// BudgetProjectedCost returns the projected cost for a budget
func (c *AnalysisController) BudgetProjectedCost(budgetID int64) (int64, error) {
	return c.service.BudgetProjectedCost(budgetID)
}

// BudgetTotalCost returns the actual cost for a budget
func (c *AnalysisController) BudgetTotalCost(budgetID int64) (int64, error) {
	return c.service.BudgetTotalCost(budgetID)
}

// AllocationCost returns the actual cost for a specific allocation
func (c *AnalysisController) AllocationCost(budgetID, categoryID int64) (int64, error) {
	return c.service.AllocationCost(budgetID, categoryID)
}


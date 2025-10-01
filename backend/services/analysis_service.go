package services

import (
	"bread/backend/persistence" 
)

type AnalysisService struct{}

func (s *AnalysisService) BudgetProjectedCost(budgetID int64) (int64, error) {
	cost, err := persistence.BudgetTotalProjectedCost(budgetID, nil)
	if err != nil {
		return 0, err
	}
	return cost, err
}

func (s *AnalysisService) BudgetTotalCost(budgetID int64) (int64, error) {
	cost, err := persistence.BudgetActualCost(budgetID, nil)
	if err != nil {
		return 0, err
	}
	return cost, err
}

func (s *AnalysisService) AllocationCost(budgetID int64, categoryID int64) (int64, error) {
	cost, err := persistence.AllocationCost(budgetID, &categoryID, nil)
	if err != nil {
		return 0, err
	}

	return cost, err
}

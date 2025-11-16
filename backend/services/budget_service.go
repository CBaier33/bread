package services

import (
	"bread/backend/models"
	"bread/backend/persistence" 
	"fmt"
	"time"
)

type BudgetService struct{}

func (s *BudgetService) CreateBudget(periodID int64, name, periodStart, periodEnd string, expectedIncome, startingBalance int64) (models.Budget, error) {
	b := models.Budget{
		ProjectID:   periodID,
		Name:        name,
		PeriodStart: periodStart,
		PeriodEnd:   periodEnd,
		ExpectedIncome: expectedIncome,
		StartingBalance: startingBalance,
	}

	id, err := persistence.InsertBudget(b, nil)
	if err != nil {
		return b, fmt.Errorf("CreateBudget: %w", err)
	}
	b.ID = id
	return b, nil
}

// GetBudget retrieves a budget by ID
func (s *BudgetService) GetBudget(id int64) (models.Budget, error) {
	b, err := persistence.GetBudget(id, nil)
	if err != nil {
		return b, fmt.Errorf("GetBudget: %w", err)
	}
	return b, nil
}

// ListBudgets returns all budgets
func (s *BudgetService) ListBudgets(projectID int64) ([]models.Budget, error) {
	budgets, err := persistence.ListBudgets(projectID, nil)
	if err != nil {
		return nil, fmt.Errorf("ListBudgets: %w", err)
	}
	return budgets, nil
}

func (s *BudgetService) UpdateBudget(b models.Budget) error {
	b.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
	if err := persistence.UpdateBudget(b, nil); err != nil {
		return fmt.Errorf("UpdateBudget: %w", err)
	}
	return nil
}

func (s *BudgetService) DeleteBudget(id int64) error {
	if err := persistence.DeleteBudget(id, nil); err != nil {
		return fmt.Errorf("DeleteBudget: %w", err)
	}
	return nil
}


func (s *BudgetService) AddAllocation(budgetID, categoryID, expectedCost int64) error {

	ba := models.BudgetAllocation {
		BudgetID: budgetID,
		CategoryID: categoryID,
		ExpectedCost: expectedCost,
	}

	_, err := persistence.InsertAllocation(ba, nil)

	return err
}

func (s *BudgetService) UpdateAllocationCost(budgetID, categoryID, newCost int64) error {
	ba, err := persistence.GetAllocation(budgetID, categoryID, nil)
	if err != nil {
		return err
	}

	ba.ExpectedCost = newCost

	err = persistence.UpdateAllocation(ba, nil)

	if err != nil {
		return err
	}

	return nil
}

func (s *BudgetService) DeleteAllocation(id int64) error {
	if err := persistence.DeleteAllocation(id, nil); err != nil {
		return fmt.Errorf("DeleteBudget: %w", err)
	}
	return nil
}

func (s *BudgetService) ListAllocations(budgetID int64) ([]models.BudgetAllocation, error) {
	if allocs, err := persistence.ListAllocations(budgetID, nil); err != nil {
		return allocs, fmt.Errorf("ListAllocations: %w", err)
	} else{
		return allocs, nil
	}
}

// Business Logic

func (s *BudgetService) DuplicateBudget(periodID, oldBudgetID int64, name, periodStart, periodEnd string) (newID int64, err error) {
    tx, err := persistence.DB.Begin()

    if err != nil {
        return 0, fmt.Errorf("DuplicateBudget -> tx.Begin: %w", err)
    }

    defer func() {
        if err != nil {
            _ = tx.Rollback()
        }
    }()

    // Create new budget
    newBudget := models.Budget{
        ProjectID:   periodID,
        Name:        name,
        PeriodStart: periodStart,
        PeriodEnd:   periodEnd,
    }

    newID, err = persistence.InsertBudget(newBudget, tx)
    if err != nil {
        return 0, fmt.Errorf("DuplicateBudget -> InsertBudget: %w", err)
    }

    var oldAllocs []models.BudgetAllocation
    oldAllocs, err = persistence.ListAllocations(oldBudgetID, tx)
    if err != nil {
        return 0, fmt.Errorf("DuplicateBudget -> ListAllocations: %w", err)
    }

    for _, oldAllocation := range oldAllocs {

        ba := models.BudgetAllocation{
            BudgetID:    newID,
            CategoryID:  oldAllocation.CategoryID,
            ExpectedCost: oldAllocation.ExpectedCost,
        }

        _, err = persistence.InsertAllocation(ba, tx)

        if err != nil {
            return 0, fmt.Errorf("DuplicateBudget -> InsertAllocation: %w", err)
        }
    }

    if err = tx.Commit(); err != nil {
        return 0, fmt.Errorf("DuplicateBudget -> tx.Commit: %w", err)
    }

    return newID, nil
}


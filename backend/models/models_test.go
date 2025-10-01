package models

import (
	"testing"
)

func TestProjectModel(t *testing.T) {
	p := Project{
		ID:          1,
		Name:        "Test Project",
		Description: "Testing project persistence",
		Currency:    "USD",
		CreatedAt:   "2025-09-23T00:00:00Z",
		UpdatedAt:   "2025-09-23T00:00:00Z",
	}

	if p.Name != "Test Project" {
		t.Errorf("expected Project name 'Test Project', got %s", p.Name)
	}
}

func TestGroupModel(t *testing.T) {
	g := Group{
		ID:          1,
		ProjectID:   1,
		Name:        "Ops",
		Description: "Operations group",
		CreatedAt:   "2025-09-23T00:00:00Z",
		UpdatedAt:   "2025-09-23T00:00:00Z",
	}

	if g.ProjectID != 1 {
		t.Errorf("expected Group.ProjectID = 1, got %d", g.ProjectID)
	}
}

func TestCategoryModel(t *testing.T) {
	c := Category{
		ID:          1,
		GroupID:     nil,
		Name:        "Utilities",
		Description: "Monthly bills",
		ExpenseType: true, // withdrawal
		Expected:    5000,
		Actual:      4500,
		CreatedAt:   "2025-09-23T00:00:00Z",
		UpdatedAt:   "2025-09-23T00:00:00Z",
	}

	if !c.ExpenseType {
		t.Errorf("expected ExpenseType=true (withdrawal), got false")
	}
}

func TestBudgetModel(t *testing.T) {
	b := Budget{
		ID:              1,
		ProjectID:       1,
		Name:            "Q4 Budget",
		PeriodStart:     "2025-10-01",
		PeriodEnd:       "2025-12-31",
		ExpectedIncome:  100000,
		StartingBalance: 20000,
		CreatedAt:       "2025-09-23T00:00:00Z",
		UpdatedAt:       "2025-09-23T00:00:00Z",
	}

	if b.ExpectedIncome != 100000 {
		t.Errorf("expected ExpectedIncome=100000, got %d", b.ExpectedIncome)
	}
}

func TestBudgetAllocationModel(t *testing.T) {
	ba := BudgetAllocation{
		ID:           1,
		BudgetID:     1,
		CategoryID:   1,
		ExpectedCost: 5000,
		CreatedAt:    "2025-09-23T00:00:00Z",
		UpdatedAt:    "2025-09-23T00:00:00Z",
	}

	if ba.ExpectedCost != 5000 {
		t.Errorf("expected ExpectedCost=5000, got %d", ba.ExpectedCost)
	}
}

func TestTransactionModelWithdraw(t *testing.T) {
	tx := Transaction{
		ID:          1,
		Description: "Grocery shopping",
		ProjectID:   1,
		CategoryID:  nil,
		Date:        "2025-09-23",
		Amount:      1500,
		ExpenseType: true, // withdrawal
		Notes:       "Whole Foods",
		CreatedAt:   "2025-09-23T00:00:00Z",
		UpdatedAt:   "2025-09-23T00:00:00Z",
	}

	if !tx.ExpenseType {
		t.Errorf("expected ExpenseType=true (withdrawal), got false")
	}
}

func TestTransactionModelDeposit(t *testing.T) {
	tx := Transaction{
		ID:          2,
		Description: "Paycheck",
		ProjectID:   1,
		CategoryID:  nil,
		Date:        "2025-09-23",
		Amount:      5000,
		ExpenseType: false, // deposit
		Notes:       "Monthly salary",
		CreatedAt:   "2025-09-23T00:00:00Z",
		UpdatedAt:   "2025-09-23T00:00:00Z",
	}

	if tx.ExpenseType {
		t.Errorf("expected ExpenseType=false (deposit), got true")
	}
}

func TestTagModel(t *testing.T) {
	tag := Tag{
		ID:        1,
		Name:      "groceries",
		CreatedAt: "2025-09-23T00:00:00Z",
		UpdatedAt: "2025-09-23T00:00:00Z",
	}

	if tag.Name != "groceries" {
		t.Errorf("expected Tag.Name='groceries', got %s", tag.Name)
	}
}

func TestTransactionTagModel(t *testing.T) {
	tt := TransactionTag{
		TransactionID: 1,
		TagID:         1,
		CreatedAt:     "2025-09-23T00:00:00Z",
	}

	if tt.TransactionID != 1 {
		t.Errorf("expected TransactionID=1, got %d", tt.TransactionID)
	}
}


package persistence

import (
	"bread/backend/models"
	"testing"
	"time"
)

func TestBudgetPersistence(t *testing.T) {
	db := SetupTestDB(t)
	DB = db

	now := time.Now().UTC().Format(time.RFC3339)

	// --- Budget CRUD ---
	p, err := InsertProject(models.Project{
		Name: "Project 1",
	})
	if err != nil {
		t.Fatalf("InsertProject failed: %v", err)
	}

	g, err := InsertGroup(models.Group{
		ProjectID: p,
		Name: "Test Group",
	})
	if err != nil {
		t.Fatalf("InsertGroup failed: %v", err)
	}

	c, err := InsertCategory(models.Category{
		GroupID: &g,
		Name: "Test Category",
		ExpenseType: true,

	})
	if err != nil {
		t.Fatalf("InsertCategory failed: %v", err)
	}

	b := models.Budget{
		ProjectID:       p,
		Name:            "September Budget",
		PeriodStart:     "2025-09-01",
		PeriodEnd:       "2025-09-30",
		ExpectedIncome:  5000,
		StartingBalance: 1000,
		CreatedAt:       now,
		UpdatedAt:       now,
	}

	// Insert
	id, err := InsertBudget(b)
	if err != nil {
		t.Fatalf("InsertBudget failed: %v", err)
	}
	if id == 0 {
		t.Fatal("expected non-zero ID after insert")
	}
	b.ID = id

	// Get
	got, err := GetBudget(id)
	if err != nil {
		t.Fatalf("GetBudget failed: %v", err)
	}
	if got.Name != b.Name {
		t.Errorf("expected name=%s, got=%s", b.Name, got.Name)
	}

	// List
	list, err := ListBudgets()
	if err != nil {
		t.Fatalf("ListBudgets failed: %v", err)
	}
	if len(list) != 1 {
		t.Errorf("expected 1 budget, got %d", len(list))
	}

	// Update
	b.Name = "Updated Budget"
	b.ExpectedIncome = 6000
	b.UpdatedAt = time.Now().UTC().Format(time.RFC3339)
	if err := UpdateBudget(b); err != nil {
		t.Fatalf("UpdateBudget failed: %v", err)
	}

	got2, _ := GetBudget(id)
	if got2.ExpectedIncome != b.ExpectedIncome {
		t.Errorf("expected update income=%d, got=%d", b.ExpectedIncome, got2.ExpectedIncome)
	}
	if got2.Name != b.Name {
		t.Errorf("expected updated name=%s, got=%s", b.Name, got2.Name)
	}

	// --- Budget Allocation CRUD ---

	a := models.BudgetAllocation{
		BudgetID:     got2.ID,
		CategoryID:   c, // assume a category ID exists in migrations
		ExpectedCost: 200,
		CreatedAt:    now,
		UpdatedAt:    now,
	}

	// Insert
	allocID, err := InsertAllocation(a)
	if err != nil {
		t.Fatalf("InsertAllocation failed: %v", err)
	}
	a.ID = allocID

	// Get
	allocGot, err := GetAllocation(allocID)
	if err != nil {
		t.Fatalf("GetAllocation failed: %v", err)
	}
	if allocGot.ExpectedCost != a.ExpectedCost {
		t.Errorf("expected ExpectedCost=%d, got=%d", a.ExpectedCost, allocGot.ExpectedCost)
	}

	// List
	allocList, err := ListAllocations()
	if err != nil {
		t.Fatalf("ListAllocations failed: %v", err)
	}
	if len(allocList) != 1 {
		t.Errorf("expected 1 allocation, got %d", len(allocList))
	}

	// Update
	a.ExpectedCost = 250
	a.UpdatedAt = time.Now().UTC().Format(time.RFC3339)
	if err := UpdateAllocation(a); err != nil {
		t.Fatalf("UpdateAllocation failed: %v", err)
	}
	allocGot2, _ := GetAllocation(allocID)
	if allocGot2.ExpectedCost != a.ExpectedCost {
		t.Errorf("expected updated ExpectedCost=%d, got=%d", a.ExpectedCost, allocGot2.ExpectedCost)
	}

	// Delete Allocation
	if err := DeleteAllocation(allocID); err != nil {
		t.Fatalf("DeleteAllocation failed: %v", err)
	}
	if _, err := GetAllocation(allocID); err == nil {
		t.Errorf("expected error after deleting allocation, got nil")
	}

	// Delete Budget
	if err := DeleteBudget(id); err != nil {
		t.Fatalf("DeleteBudget failed: %v", err)
	}
	if _, err := GetBudget(id); err == nil {
		t.Errorf("expected error after deleting budget, got nil")
	}
}


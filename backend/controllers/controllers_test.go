package controllers_test

import (
	"bread/backend/controllers"
	"bread/backend/persistence"
	"testing"
)

func TestControllers(t *testing.T) {
	db := persistence.SetupTestDB(t)
	defer db.Close()
	persistence.DB = db

	// Instantiate controllers using their constructors
	bc := controllers.NewBudgetController()
	gc := controllers.NewGroupController()
	cc := controllers.NewCategoryController()
	pc := controllers.NewProjectController()
	tc := controllers.NewTransactionController()
	tgc := controllers.NewTagController()
	ac := controllers.NewAnalysisController()

	// --- Projects ---
	project1, err := pc.CreateProject("Personal", "for budgeting income", "USD")
	if err != nil {
		t.Fatal(err)
	}

	// --- Budgets ---
	budget1, err := bc.CreateBudget(project1.ID, "Test", "2025-09-01", "2025-09-30")
	if err != nil {
		t.Fatal(err)
	}

	budget2, err := bc.CreateBudget(project1.ID, "test2", "2025-01-01", "2025-06-30")
	if err != nil {
		t.Fatal(err)
	}

	budget2.Name = "bread"
	if err := bc.UpdateBudget(budget2); err != nil {
		t.Fatal(err)
	}

	if err := bc.DeleteBudget(budget2.ID); err != nil {
		t.Fatal(err)
	}

	budgets, err := bc.ListBudgets(project1.ID)
	if err != nil {
		t.Fatal(err)
	}
	if len(budgets) != 1 {
		t.Errorf("expected budgets length 1, got %d", len(budgets))
	}

	// --- Groups & Categories ---
	group1, err := gc.CreateGroup(project1.ID, "Test", "test group")
	if err != nil {
		t.Fatal(err)
	}

	category1, err := cc.CreateCategory(group1.ID, "TestCategory", "Desc", true)
	if err != nil {
		t.Fatal(err)
	}

	// --- Budget Allocations ---
	if err := bc.AddAllocation(budget1.ID, category1.ID, 200); err != nil {
		t.Fatal(err)
	}

	budget3ID, err := bc.DuplicateBudget(project1.ID, budget1.ID, "Dup Budget", "2025-10-01", "2025-10-31")
	if err != nil {
		t.Fatal(err)
	}

	budget3Allocs, err := bc.ListAllocations(budget3ID)
	budget1Allocs, err := bc.ListAllocations(budget1.ID)
	if len(budget1Allocs) != len(budget3Allocs) || len(budget1Allocs) == 0 {
		t.Errorf("DuplicateBudget failed: Budget1 %d vs Budget3 %d", len(budget1Allocs), len(budget3Allocs))
	}

	// --- Transactions ---
	_, err = tc.CreateTransaction(project1.ID, nil, "burger", 100, "2025-10-13", true, "")
	if err != nil {
		t.Fatal(err)
	}

	tx2, err := tc.CreateTransaction(project1.ID, &category1.ID, "grocery", 100, "2025-10-20", true, "")
	if err != nil {
		t.Fatal(err)
	}

	// --- Tags ---
	tag1, err := tgc.CreateTag(project1.ID, "publix")
	if err != nil {
		t.Fatal(err)
	}

	if _, err := tgc.CreateTransactionTag(tx2.ID, tag1.ID); err != nil {
		t.Fatal(err)
	}

	tags, err := tgc.GetTags(tx2.ID)
	if err != nil {
		t.Fatal(err)
	}
	if len(tags) == 0 {
		t.Errorf("expected 1 tag for transaction, got 0")
	}

	// --- Analysis ---
	totalCost, err := ac.BudgetTotalCost(budget3ID)
	if err != nil {
		t.Fatal(err)
	}
	if totalCost != 200 {
		t.Errorf("expected total cost 200, got %d", totalCost)
	}

	totalProjectedCost, err := ac.BudgetProjectedCost(budget1.ID)
	if err != nil {
		t.Fatal(err)
	}
	if totalProjectedCost != 200 {
		t.Errorf("expected projected cost 200, got %d", totalProjectedCost)
	}

	allocCost, err := ac.AllocationCost(budget3ID, category1.ID)
	if err != nil {
		t.Fatal(err)
	}
	if allocCost != 100 {
		t.Errorf("expected allocation cost 100, got %d", allocCost)
	}
}


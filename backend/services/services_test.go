package services

import (
	"bread/backend/persistence"
	"testing"
)

func TestServices(t *testing.T) {
	db := persistence.SetupTestDB(t)
	defer db.Close()
	persistence.DB = db 

	ps := &ProjectService{}
	bs := &BudgetService{}
	gs := &GroupService{}
	cs := &CategoryService{}
	ts := &TransactionService{}
	tgs := &TagService{}
	as := &AnalysisService{}

	project1, err := ps.CreateProject("Personal", "for budgeting income", "USD" )

	budget1, err := bs.CreateBudget(project1.ID, "Test", "2025-09-01", "2025-09-31")

	if err != nil {
		t.Fatal(err)
	}

	group1, err := gs.CreateGroup(project1.ID, "Test", "test group")
	category1, err := cs.CreateCategory(group1.ID, "Test", "2025-09-01", true)

	if err != nil {
		t.Fatal(err)
	}

	budget2, err := bs.CreateBudget(project1.ID,"test2", "2025-01-01", "2025-06-31")

	if err != nil {
		t.Fatal(err)
	}

	budget2.Name = "bread"

	err = bs.UpdateBudget(budget2)
	if err != nil {
		t.Fatal(err)
	}

	err = bs.DeleteBudget(budget2.ID)

	budgets, err := bs.ListBudgets(project1.ID)

	if len(budgets) != 1 {
    t.Errorf("expected budges length of 1, got %d", len(budgets))
	}

	err = bs.AddAllocation(budget1.ID, category1.ID, 200)

	if err != nil {
		t.Fatal(err)
	}

	budget3_id, err := bs.DuplicateBudget(project1.ID, budget1.ID, "Dup Budget", "2025-10-01", "2025-10-31")

	if err != nil {
		t.Fatal(err)
	}

	budgets, err = bs.ListBudgets(project1.ID)

	if len(budgets) != 2 {
    t.Errorf("expected budges length of 2, got %d", len(budgets))
	}

	budget3_allocs, err := bs.ListAllocations(budget3_id)

	budget1_allocs, err := bs.ListAllocations(budget1.ID)

	if len(budget1_allocs) != len(budget3_allocs)  || len(budget1_allocs) == 0{
		t.Errorf("DuplicateBudget: Duplication failed, Budget1 : %d | Budget2: %d", len(budget1_allocs), len(budget3_allocs))
	}

	_, err = ts.CreateTransaction(project1.ID, nil, "burgor", 100, "2025-10-13", true, "")

	if err != nil {
		t.Errorf("CreateTransaction: %s", err)
	}

	tag1, err := tgs.CreateTag(project1.ID, "publix")
	transaction1, err := ts.CreateTransaction(project1.ID, &category1.ID, "grocery", 100, "2025-10-20", true, "")
	_, err = tgs.CreateTransactionTag(transaction1.ID, tag1.ID)

	t1Tags, err := tgs.GetTags(transaction1.ID)
	if err != nil {
		t.Errorf("GetTags: %s", err)
	}
	if len(t1Tags) == 0 {
		t.Errorf("GetTags length expected 1 got 0.")
	}

	_, err = ts.CreateTransaction(project1.ID, &category1.ID, "grocery", 100, "2025-11-01", true, "") // not in budget3



	if err != nil {
		t.Errorf("CreateTransaction: %s", err)
	}

	totalCost, err := as.BudgetTotalCost(budget3_id)

	if err != nil {
		t.Errorf("BudgetTotalCost: %s", err)
	}

	if totalCost != 200 {
		t.Errorf("BudgetTotalCost: Incorrect value")
	}

	totalProjectedCost, err := as.BudgetProjectedCost(budget1.ID)
	if err != nil {
		t.Errorf("BudgetProjectedCost: %s", err)
	}

	if totalProjectedCost != 200 {
		t.Errorf("BudgetProjectedCost: %s", err)
	}

	allocCost, err := as.AllocationCost(budget3_id, category1.ID)
	if err != nil {
		t.Errorf("AllocationCost: %s", err)
	}

	if allocCost != 100 {
		t.Errorf("Unexpected Allocation Cost amount. Expected 100, got %d", allocCost)
	}

}

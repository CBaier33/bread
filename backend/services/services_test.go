package services

import (
	"bread/backend/models"
	"bread/backend/persistence"
	"testing"
)

func TestBudget(t *testing.T) {
	db := persistence.SetupTestDB(t)
	defer db.Close()
	persistence.DB = db 

	bs := &BudgetService{}

	// Optional: create a budget/group/category if needed
	budget1, err := bs.CreateBudget("Test", "2025-09-01", "2025-09-31")

	if err != nil {
		t.Fatal(err)
	}

	budget2, err := bs.CreateBudget("test2", "2025-01-01", "2025-06-31")

	if err != nil {
		t.Fatal(err)
	}

	budget2.Name = "bread"

	err = bs.UpdateBudget(budget2)
	if err != nil {
		t.Fatal(err)
	}

	err = bs.DeleteBudget(budget1.ID)

	budgets, err := bs.ListBudgets()

	if len(budgets) != 1 {
    t.Errorf("expected budges length of 1, got %d", len(budgets))
	}

	nextBudget, _ := bs.GetBudget(budget2.ID)

	if err != nil {
		t.Fatal(err)
	}

	if nextBudget.Name != "bread" {
		t.Errorf("expected budget name to be 'bread', got " + nextBudget.Name)
	}
	

}

func TestGroup(t *testing.T) {
	db := persistence.SetupTestDB(t)
	defer db.Close()
	persistence.DB = db 

	bs := &BudgetService{}

	// Optional: create a budget/group/category if needed
	budget1, err := bs.CreateBudget("Test", "2025-09-01", "2025-09-31")

	gs := &GroupService{}

	var budgetid int64 = budget1.ID

	group1, err := gs.CreateGroup(budgetid, "Test Group", "description")

	if err != nil {
		t.Fatal(err)
	}

	group2, err := gs.CreateGroup(budgetid, "Test Group 2", "description")

	if err != nil {
		t.Fatal(err)
	}

	group2.Name = "bread"

	err = gs.UpdateGroup(group2)
	if err != nil {
		t.Fatal(err)
	}

	err = gs.DeleteGroup(group1.ID)

	budgets, err := gs.ListGroups()

	if len(budgets) != 1 {
    t.Errorf("expected budges length of 1, got %d", len(budgets))
	}

	nextGroup, _ := gs.GetGroup(group2.ID)

	if err != nil {
		t.Fatal(err)
	}

	if nextGroup.Name != "bread" {
		t.Errorf("expected budget name to be 'bread', got " + nextGroup.Name)
	}

}


func TestCreateTransaction(t *testing.T) {

	// Create a fresh test DB
	db := persistence.SetupTestDB(t)
	defer db.Close()
	persistence.DB = db // point repo layer to this test DB

	// Optional: create a budget/group/category if needed
	budgetID, _ := persistence.InsertBudget(models.Budget{Name: "Test Budget"})
	groupID, _ := persistence.InsertGroup(models.Group{BudgetID: budgetID, Name: "Test Group"})
	categoryID, _ := persistence.InsertCategory(models.Category{BudgetID: budgetID, GroupID: &groupID, Name: "Test Category"})

	ts := &TransactionService{}

	t.Run("Transaction with category", func(t *testing.T) {
		tx, err := ts.CreateTransaction("Lunch", 12345, "2025-09-09", "Sushi", "", budgetID, nil, &categoryID)
		if err != nil {
			t.Fatal(err)
		}
		if tx.ID == 0 {
			t.Fatal("expected transaction ID > 0")
		}
	})

	t.Run("Transaction without category", func(t *testing.T) {
		tx, err := ts.CreateTransaction("Cash deposit", 5000, "2025-09-09", "ATM", "money:", budgetID, nil, nil)
		if err != nil {
			t.Fatal(err)
		}
		if tx.ID == 0 {
			t.Fatal("expected transaction ID > 0")
		}
	})
}


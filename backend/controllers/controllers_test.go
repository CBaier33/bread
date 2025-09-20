package controllers

import (
	"bread/backend/persistence"
	"testing"
	"strconv"
)

func TestBudget(t *testing.T) {
	db := persistence.SetupTestDB(t)
	defer db.Close()
	persistence.DB = db 


	bc := NewBudgetController()

	// Optional: create a budget/group/category if needed
	budget1, err := bc.CreateBudget("Test", "2025-09-01", "2025-09-31")

	if err != nil {
		t.Fatal(err)
	}

	budget2, err := bc.CreateBudget("test2", "2025-01-01", "2025-06-31")

	if err != nil {
		t.Fatal(err)
	}

	budget2.Name = "bread"

	err = bc.UpdateBudget(budget2)
	if err != nil {
		t.Fatal(err)
	}

	err = bc.DeleteBudget(budget1.ID)

	budgets, err := bc.ListBudgets()

	if len(budgets) != 1 {
    t.Errorf("expected budges length of 1, got %d", len(budgets))
	}

	nextBudget, _ := bc.GetBudget(budget2.ID)

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

	bc := NewBudgetController()

	// Optional: create a budget/group/category if needed
	budget1, err := bc.CreateBudget("Test", "2025-09-01", "2025-09-31")
	budgetid := budget1.ID

	gc := NewGroupController()

	group1, err := gc.CreateGroup(budgetid, "Test Group", "description")

	if err != nil {
		t.Fatal(err)
	}

	group2, err := gc.CreateGroup(budgetid, "Test Group 2", "description")

	if err != nil {
		t.Fatal(err)
	}

	group2.Name = "bread"

	err = gc.UpdateGroup(group2)
	if err != nil {
		t.Fatal(err)
	}

	err = gc.DeleteGroup(group1.ID)

	groups, err := gc.ListGroups()

	if len(groups) != 1 {
    t.Errorf("expected budges length of 1, got %d", len(groups))
	}

	nextGroup, _ := gc.GetGroup(group2.ID)

	if err != nil {
		t.Fatal(err)
	}

	if nextGroup.Name != "bread" {
		t.Errorf("expected budget name to be 'bread', got " + nextGroup.Name)
	}

}

func TestCategory(t *testing.T) {
	db := persistence.SetupTestDB(t)
	defer db.Close()
	persistence.DB = db 

	bc := NewBudgetController()

	// Optional: create a budget/group/category if needed
	budget1, err := bc.CreateBudget("Test", "2025-09-01", "2025-09-31")
	budgetid := budget1.ID

	gc := NewGroupController()
	group1, err := gc.CreateGroup(budgetid, "Test Group", "description")
	groupid := group1.ID

	cg := NewCategoryController()

	category1, err := cg.CreateCategory(budgetid, groupid, "Test Category", "description", 1000, 0)

	if err != nil {
		t.Fatal(err)
	}

	category2, err := cg.CreateCategory(budgetid, groupid, "Test Category", "description", 200, 0)

	if err != nil {
		t.Fatal(err)
	}

	category2.Expected += 1

	err = cg.UpdateCategory(category2)
	if err != nil {
		t.Fatal(err)
	}

	err = cg.DeleteCategory(category1.ID)

	categories, err := cg.ListCategories()

	if len(categories) != 1 {
    t.Errorf("expected budges length of 1, got %d", len(categories))
	}

	nextCategory, _ := cg.GetCategory(category2.ID)

	if err != nil {
		t.Fatal(err)
	}

	if nextCategory.Expected != 201 {
		t.Errorf("expected category amount to be 201, got " + strconv.FormatInt(nextCategory.Expected, 10))
	}

}

func TestTransaction(t *testing.T) {
	db := persistence.SetupTestDB(t)
	defer db.Close()
	persistence.DB = db 

	bc := NewBudgetController()

	// Optional: create a budget/group/category if needed
	budget1, err := bc.CreateBudget("Test", "2025-09-01", "2025-09-31")
	budgetid := budget1.ID

	gc := NewGroupController()
	group1, err := gc.CreateGroup(budgetid, "Test Group", "description")
	groupid := group1.ID

	cg := NewCategoryController()

	category1, err := cg.CreateCategory(budgetid, groupid, "Test Category Food", "description", 1000, 0)
	category1Id := category1.ID

	tc := NewTransactionController()

	transaction1, err := tc.CreateTransaction("butter-1", 100, "2025-09-22", "", "", budgetid, &groupid, &category1Id)

	if err != nil {
		t.Fatal(err)
	}

	// UnGrouped category
	transaction2, err := tc.CreateTransaction("butter", 200, "2025-09-21", "", "", budgetid, nil, nil)

	if err != nil {
		t.Fatal(err)
	}

	transaction3, err := tc.CreateTransaction("test", 200, "2025-09-23", "", "", budgetid, &groupid, &category1Id)

	if err != nil {
		t.Fatal(err)
	}

	transaction2.Description = "bread"

	err = tc.UpdateTransaction(transaction2)

	if err != nil {
		t.Fatal(err)
	}

	transaction1.Amount = 20

	err = tc.UpdateTransaction(transaction1)

	if err != nil {
		t.Fatal(err)
	}

	otherTransaction, err := tc.GetTransaction(transaction1.ID)

	if err != nil {
		t.Fatal(err)
	}

	if transaction1.Amount != 20 {
    t.Errorf("expected transaction amount of 20, got %d", otherTransaction.Amount)
	}

	err = tc.DeleteTransaction(transaction3.ID)

	transactions, err := tc.ListTransactions(&category1Id)

	if len(transactions) != 1 {
    t.Errorf("expected budges length of 1, got %d", len(transactions))
	}

	nextTransaction, _ := tc.GetTransaction(transaction2.ID)

	if err != nil {
		t.Fatal(err)
	}

	if nextTransaction.Description != "bread" {
		t.Errorf("expected budget name to be 'bread', got " + nextTransaction.Description)
	}

}

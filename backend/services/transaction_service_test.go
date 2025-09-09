package services

import (
	"bread/backend/models"
	"bread/backend/persistence"
	"testing"
)

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
		tx, err := ts.CreateTransaction("Lunch", 12345, "2025-09-09", "Sushi", budgetID, nil, &categoryID)
		if err != nil {
			t.Fatal(err)
		}
		if tx.ID == 0 {
			t.Fatal("expected transaction ID > 0")
		}
	})

	t.Run("Transaction without category", func(t *testing.T) {
		tx, err := ts.CreateTransaction("Cash deposit", 5000, "2025-09-09", "ATM", budgetID, nil, nil)
		if err != nil {
			t.Fatal(err)
		}
		if tx.ID == 0 {
			t.Fatal("expected transaction ID > 0")
		}
	})
}


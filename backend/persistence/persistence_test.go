package persistence

import (
	"bread/backend/models"
	"testing"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func TestPersistence(t *testing.T) {
	db := SetupTestDB(t)
	defer db.Close()
	DB = db

	currentTS := time.Now().Format("2006-01-02 15:04:05")

	var budgetID, groupID, categoryID int64

	t.Run("Budget Test", func(t *testing.T) {
		b := models.Budget{
			Name:        "September Budget",
			PeriodStart: "2025-09-01",
			PeriodEnd:   "2025-09-30",
			CreatedAt:   currentTS,
			UpdatedAt:   currentTS,
		}
		var err error
		budgetID, err = InsertBudget(b)
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("Created budget ID=%d", budgetID)

		got, err := GetBudget(budgetID)
		if err != nil {
			t.Fatal(err)
		}
		if got.Name != b.Name {
			t.Fatalf("expected %s, got %s", b.Name, got.Name)
		}
		t.Logf("Retrieved budget successfully")

		b.Name = "Updated Budget"
		b.ID = budgetID
		if err := UpdateBudget(b); err != nil {
			t.Fatal(err)
		}
		t.Logf("Updated budget successfully")
	})

	t.Run("Group Test", func(t *testing.T) {
		g := models.Group{
			BudgetID:    budgetID,
			Name:        "Test Group",
			Description: "Group description",
			CreatedAt:   currentTS,
			UpdatedAt:   currentTS,
		}
		var err error
		groupID, err = InsertGroup(g)
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("Created group ID=%d", groupID)

		got, err := GetGroup(groupID)
		if err != nil {
			t.Fatal(err)
		}
		if got.Name != g.Name {
			t.Fatalf("expected %s, got %s", g.Name, got.Name)
		}
		t.Logf("Retrieved group successfully")
	})

	t.Run("Category Test", func(t *testing.T) {
		c := models.Category{
			BudgetID:    budgetID,
			GroupID:     &groupID,
			Name:        "Food",
			Description: "Food expenses",
			CreatedAt:   currentTS,
			UpdatedAt:   currentTS,
		}
		var err error
		categoryID, err = InsertCategory(c)
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("Created category ID=%d", categoryID)

		got, err := GetCategory(categoryID)
		if err != nil {
			t.Fatal(err)
		}
		if got.Name != c.Name {
			t.Fatalf("expected %s, got %s", c.Name, got.Name)
		}
		t.Logf("Retrieved category successfully")
	})

	t.Run("Transaction Test", func(t *testing.T) {
		// Transaction WITH category
		txWithCat := models.Transaction{
			Description: "Lunch",
			BudgetID:    budgetID,
			GroupID:     &groupID,
			CategoryID:  &categoryID,
			Date:        currentTS[:10],
			Amount:      12345,
			Notes:       "Sushi",
			CreatedAt:   currentTS,
			UpdatedAt:   currentTS,
		}
		id1, err := InsertTransaction(txWithCat)
		if err != nil {
			t.Fatalf("InsertTransaction with category failed: %v", err)
		}
		t.Logf("Inserted transaction with category ID=%d", id1)

		// Transaction WITHOUT category OR group
		txNoCat := models.Transaction{
			Description: "Cash deposit",
			BudgetID:    budgetID,
			GroupID:     nil,
			CategoryID:  nil,
			Date:        currentTS[:10],
			Amount:      5000,
			Notes:       "ATM deposit",
			CreatedAt:   currentTS,
			UpdatedAt:   currentTS,
		}
		id2, err := InsertTransaction(txNoCat)
		if err != nil {
			t.Fatalf("InsertTransaction without category failed: %v", err)
		}
		t.Logf("Inserted transaction without category ID=%d", id2)

		// Transaction WITHOUT category OR group
		txNoGRP := models.Transaction{
			Description: "Bitcoin deposit",
			BudgetID:    budgetID,
			GroupID:     &groupID,
			CategoryID:  nil,
			Date:        currentTS[:10],
			Amount:      5012,
			Notes:       "BTC deposit",
			CreatedAt:   currentTS,
			UpdatedAt:   currentTS,
		}

		id3, err := InsertTransaction(txNoGRP)
		if err != nil {
			t.Fatalf("InsertTransaction without category failed: %v", err)
		}
		t.Logf("Inserted transaction without category ID=%d", id3)
		// Retrieve transactions
		txns, err := GetTransactions()
		if err != nil {
			t.Fatalf("GetTransactions failed: %v", err)
		}
		if len(txns) != 3 {
			t.Fatalf("Expected 2 transactions, got %d", len(txns))
		}
		t.Logf("Retrieved %d transactions", len(txns))

		// Delete transactions
		_, err = DB.Exec("DELETE FROM transactions")
		if err != nil {
			t.Fatalf("Failed to delete transactions: %v", err)
		}
		txns, _ = GetTransactions()
		if len(txns) != 0 {
			t.Fatalf("Expected 0 transactions after delete, got %d", len(txns))
		}
		t.Logf("Transactions deleted successfully: %d & %d", id1, id2)
	})
}


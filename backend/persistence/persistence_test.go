package persistence

import (
	"bread/backend/models"
	"testing"
)

func TestPersistenceBasic(t *testing.T) {
	db := SetupTestDB(t)
	DB = db

	// --- Project ---
	projectModel := models.Project{
		Name:        "Test Project",
		Description: "CRUD test",
		Currency:    "USD",
	}

	project_id, err := InsertProject(projectModel)

	if err != nil {
		t.Fatalf("InsertProject failed: %v", err)
	}
	if project_id == 0 {
		t.Fatalf("expected non-zero ID after insert")
	}

	project1, err := GetProject(project_id)
	if err != nil {
		t.Fatalf("InsertProject failed: %v", err)
	}

	project1Desc := project1.Description

	project1.Description = "New Description"

	err = UpdateProject(project1)

	if err != nil {
		t.Fatalf("UpdateProject failed: %v", err)

	}

	project1, err = GetProject(project1.ID)
	if err != nil {
		t.Fatalf("GetProject failed: %v", err)
	}

	if project1.Description == project1Desc {
		t.Errorf("Description remains unchanged.")
	}

	
	project_id2, err := InsertProject(projectModel)

	project2, err := GetProject(project_id2)
	if err != nil {
		t.Fatalf("GetProject failed: %v", err)

	}

	projects, err := ListProjects()

	if err != nil {
		t.Fatalf("ListProjects failed: %v", err)
	}
	if len(projects) != 2 {
		t.Errorf("expected 2 projects, got %d", len(projects))
	}

	err = DeleteProject(project2.ID)
	if err != nil {
		t.Fatalf("DeleteProject failed: %v", err)
	}

	projects, err = ListProjects();

	if err != nil {
		t.Fatalf("ListProjects failed: %v", err)
	}
	if len(projects) != 1 {
		t.Errorf("expected 1 projects, got %d", len(projects))
	}

	// --- Budget ---

	budgetModel := models.Budget{
		ProjectID:   project1.ID,
		Name:        "Test Budget",
		PeriodStart: "2025-09-01",
		PeriodEnd:   "2025-09-31",
		ExpectedIncome: 5200,
		StartingBalance: 100,
	}

	budget_id, err := InsertBudget(budgetModel)

	if err != nil {
		t.Fatalf("InsertBudget failed: %v", err)
	}
	if budget_id == 0 {
		t.Fatalf("expected non-zero ID after insert")
	}

	budget1, err := GetBudget(budget_id)
	if err != nil {
		t.Fatalf("InsertBudget failed: %v", err)
	}

	budget1Name := budget1.Name

	budget1.Name = "New Name"

	err = UpdateBudget(budget1)

	if err != nil {
		t.Fatalf("UpdateBudget failed: %v", err)

	}

	budget1, err = GetBudget(budget1.ID)
	if err != nil {
		t.Fatalf("GetBudget failed: %v", err)
	}

	if budget1.Name == budget1Name {
		t.Errorf("Description remains unchanged.")
	}

	
	budget_id2, err := InsertBudget(budgetModel)

	budget2, err := GetBudget(budget_id2)
	if err != nil {
		t.Fatalf("GetBudget failed: %v", err)

	}

	budgets, err := ListBudgets()

	if err != nil {
		t.Fatalf("ListBudgets failed: %v", err)
	}
	if len(budgets) != 2 {
		t.Errorf("expected 2 budgets, got %d", len(budgets))
	}

	err = DeleteBudget(budget2.ID)
	if err != nil {
		t.Fatalf("DeleteBudget failed: %v", err)
	}

	budgets, err = ListBudgets();

	if err != nil {
		t.Fatalf("ListBudgets failed: %v", err)
	}
	if len(budgets) != 1 {
		t.Errorf("expected 1 budgets, got %d", len(budgets))
	}

	// --- Groups ---

	groupModel := models.Group{
		ProjectID:   project1.ID,
		Name:        "Test Group",
	}

	group_id, err := InsertGroup(groupModel)

	if err != nil {
		t.Fatalf("InsertGroup failed: %v", err)
	}
	if group_id == 0 {
		t.Fatalf("expected non-zero ID after insert")
	}

	group1, err := GetGroup(group_id)
	if err != nil {
		t.Fatalf("InsertGroup failed: %v", err)
	}

	group1Name := group1.Name

	group1.Name = "New Name"

	err = UpdateGroup(group1)

	if err != nil {
		t.Fatalf("UpdateGroup failed: %v", err)

	}

	group1, err = GetGroup(group1.ID)
	if err != nil {
		t.Fatalf("GetGroup failed: %v", err)
	}

	if group1.Name == group1Name {
		t.Errorf("Description remains unchanged.")
	}

	
	group_id2, err := InsertGroup(groupModel)

	group2, err := GetGroup(group_id2)
	if err != nil {
		t.Fatalf("GetGroup failed: %v", err)

	}

	groups, err := ListGroups()

	if err != nil {
		t.Fatalf("ListGroups failed: %v", err)
	}
	if len(groups) != 2 {
		t.Errorf("expected 2 groups, got %d", len(groups))
	}

	err = DeleteGroup(group2.ID)
	if err != nil {
		t.Fatalf("DeleteGroup failed: %v", err)
	}

	groups, err = ListGroups();

	if err != nil {
		t.Fatalf("ListGroups failed: %v", err)
	}
	if len(groups) != 1 {
		t.Errorf("expected 1 groups, got %d", len(groups))
	}

	// --- Category ---

	categoryModel := models.Category{
		GroupID:         &group1.ID,
		Name:            "Test Category",
		ExpenseType:     true,
	}

	category_id, err := InsertCategory(categoryModel)

	if err != nil {
		t.Fatalf("InsertCategory failed: %v", err)
	}
	if category_id == 0 {
		t.Fatalf("expected non-zero ID after insert")
	}

	category1, err := GetCategory(category_id)
	if err != nil {
		t.Fatalf("GetCategory failed: %v", err)
	}

	category1Name := category1.Name

	category1.Name = "New Name"

	err = UpdateCategory(category1)

	if err != nil {
		t.Fatalf("UpdateCategory failed: %v", err)

	}

	category1, err = GetCategory(category1.ID)
	if err != nil {
		t.Fatalf("GetCategory failed: %v", err)
	}

	if category1.Name == category1Name {
		t.Errorf("Description remains unchanged.")
	}

	
	category_id2, err := InsertCategory(categoryModel)

	category2, err := GetCategory(category_id2)
	if err != nil {
		t.Fatalf("GetCategory failed: %v", err)

	}

	categories, err := ListCategories()

	if err != nil {
		t.Fatalf("ListCategories failed: %v", err)
	}
	if len(categories) != 2 {
		t.Errorf("expected 2 categories, got %d", len(categories))
	}

	err = DeleteCategory(category2.ID)
	if err != nil {
		t.Fatalf("DeleteCategory failed: %v", err)
	}

	categories, err = ListCategories();

	if err != nil {
		t.Fatalf("ListCategories failed: %v", err)
	}
	if len(categories) != 1 {
		t.Errorf("expected 1 categories, got %d", len(categories))
	}

	// --- BudgetAllocations

	allocationModel := models.BudgetAllocation{
		BudgetID: budget_id,
		CategoryID: category_id,
		ExpectedCost: 777,
	}

	allocation_id, err := InsertAllocation(allocationModel)

	if err != nil {
		t.Fatalf("InsertAllocation failed: %v", err)
	}
	if allocation_id == 0 {
		t.Fatalf("expected non-zero ID after insert")
	}

	allocation1, err := GetAllocation(allocation_id)
	if err != nil {
		t.Fatalf("InsertAllocation failed: %v", err)
	}

	allocation1Cost := allocation1.ExpectedCost

	allocation1.ExpectedCost = 200

	err = UpdateAllocation(allocation1)

	if err != nil {
		t.Fatalf("UpdateAllocation failed: %v", err)

	}

	allocation1, err = GetAllocation(allocation1.ID)
	if err != nil {
		t.Fatalf("GetAllocation failed: %v", err)
	}

	if allocation1.ExpectedCost == allocation1Cost {
		t.Errorf("Description remains unchanged.")
	}

	
	allocation_id2, err := InsertAllocation(allocationModel)
	if err != nil {
		t.Fatalf("InsertAllocation failed: %v", err)
	}

	allocation2, err := GetAllocation(allocation_id2)
	if err != nil {
		t.Fatalf("GetAllocation failed: %v", err)
	}

	allocations, err := ListAllocations()

	if err != nil {
		t.Fatalf("ListAllocations failed: %v", err)
	}
	if len(allocations) != 2 {
		t.Errorf("expected 2 allocations, got %d", len(allocations))
	}

	err = DeleteAllocation(allocation2.ID)
	if err != nil {
		t.Fatalf("DeleteAllocation failed: %v", err)
	}

	allocations, err = ListAllocations();

	if err != nil {
		t.Fatalf("ListAllocations failed: %v", err)
	}
	if len(allocations) != 1 {
		t.Errorf("expected 1 allocations, got %d", len(allocations))
	}

	// --- Transactions ---

	transactionModel := models.Transaction{
		Description: "CRUD test",
		ProjectID: project_id,
		CategoryID: &category_id,
		Date: "2025-09-02",
		Amount: 10,
		ExpenseType: true,
	}

	transaction_id, err := InsertTransaction(transactionModel)

	if err != nil {
		t.Fatalf("InsertTransaction failed: %v", err)
	}
	if transaction_id == 0 {
		t.Fatalf("expected non-zero ID after insert")
	}

	transaction1, err := GetTransaction(transaction_id)
	if err != nil {
		t.Fatalf("InsertTransaction failed: %v", err)
	}

	transaction1Desc := transaction1.Description

	transaction1.Description = "New Description"

	err = UpdateTransaction(transaction1)

	if err != nil {
		t.Fatalf("UpdateTransaction failed: %v", err)

	}

	transaction1, err = GetTransaction(transaction1.ID)
	if err != nil {
		t.Fatalf("GetTransaction failed: %v", err)
	}

	if transaction1.Description == transaction1Desc {
		t.Errorf("Description remains unchanged.")
	}

	
	transaction_id2, err := InsertTransaction(transactionModel)

	transaction2, err := GetTransaction(transaction_id2)
	if err != nil {
		t.Fatalf("GetTransaction failed: %v", err)

	}

	transactions, err := ListTransactions()

	if err != nil {
		t.Fatalf("ListTransactions failed: %v", err)
	}
	if len(transactions) != 2 {
		t.Errorf("expected 2 transactions, got %d", len(transactions))
	}

	err = DeleteTransaction(transaction2.ID)
	if err != nil {
		t.Fatalf("DeleteTransaction failed: %v", err)
	}

	transactions, err = ListTransactions();

	if err != nil {
		t.Fatalf("ListTransactions failed: %v", err)
	}
	if len(transactions) != 1 {
		t.Errorf("expected 1 transactions, got %d", len(transactions))
	}

	// -- Tags --

	tagModel := models.Tag{
		ProjectID: project_id,
		Name: "beezdurger",
	}

	tag_id, err := InsertTag(tagModel)

	if err != nil {
		t.Fatalf("InsertTag failed: %v", err)
	}
	if tag_id == 0 {
		t.Fatalf("expected non-zero ID after insert")
	}
	
	tag_id2, err := InsertTag(tagModel)

	tag2, err := GetTag(tag_id2)
	if err != nil {
		t.Fatalf("GetTag failed: %v", err)

	}

	tags, err := ListTags()

	if err != nil {
		t.Fatalf("ListTags failed: %v", err)
	}
	if len(tags) != 2 {
		t.Errorf("expected 2 tags, got %d", len(tags))
	}

	err = DeleteTag(tag2.ID)
	if err != nil {
		t.Fatalf("DeleteTag failed: %v", err)
	}

	tags, err = ListTags();

	if err != nil {
		t.Fatalf("ListTags failed: %v", err)
	}
	if len(tags) != 1 {
		t.Errorf("expected 1 tags, got %d", len(tags))
	}

	// TransactionTags

	transaction_tagModel := models.TransactionTag{
		TransactionID: transaction_id,
		TagID: tag_id,
	}

	err = InsertTransactionTag(transaction_tagModel)

	if err != nil {
		t.Fatalf("InsertTransactionTag failed: %v", err)
	}
	
	transaction_tagModel.TransactionID = transaction_id2
	err = InsertTransactionTag(transaction_tagModel)

	transaction_tag2, err := GetTransactionTag(transaction_id, tag_id)
	if err != nil {
		t.Fatalf("GetTransactionTag failed: %v", err)

	}

	if transaction_tag2.CreatedAt == "" {
		t.Fatalf("GetTransactionTag return empty")
	}

}

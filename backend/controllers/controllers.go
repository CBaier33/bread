package controllers

// All returns a slice of all controllers
func All() []interface{} {
	return []interface{}{
		NewBudgetController(),
		NewGroupController(),
		NewCategoryController(),
		NewProjectController(),
		NewTransactionController(),
		NewTagController(),
		NewAnalysisController(),
	}
}


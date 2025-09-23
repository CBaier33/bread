-- 004_create_budgets.sql
-- Creates the budgets and allocations tables

CREATE TABLE IF NOT EXISTS budgets (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    project_id INTEGER NOT NULL,
    name TEXT NOT NULL,
    period_start TEXT NOT NULL,  -- YYYY-MM-DD
    period_end TEXT NOT NULL,    -- YYYY-MM-DD
    expected_income INTEGER NOT NULL,
    starting_balance INTEGER NOT NULL,
    created_at TEXT NOT NULL DEFAULT (datetime('now')),
    updated_at TEXT NOT NULL DEFAULT (datetime('now')),
    FOREIGN KEY (project_id) REFERENCES projects(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS budget_allocations (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    budget_id INTEGER NOT NULL,
    category_id INTEGER NOT NULL,
    expected_cost INTEGER NOT NULL, 
    created_at TEXT NOT NULL DEFAULT (datetime('now')),
    updated_at TEXT NOT NULL DEFAULT (datetime('now')),
    FOREIGN KEY (budget_id) REFERENCES budgets(id) ON DELETE CASCADE,
    FOREIGN KEY (category_id) REFERENCES categories(id) ON DELETE CASCADE,
    UNIQUE (budget_id, category_id)        -- prevent duplicate allocations
);

CREATE INDEX IF NOT EXISTS idx_budget_allocations_budget_category
ON budget_allocations(budget_id, category_id);

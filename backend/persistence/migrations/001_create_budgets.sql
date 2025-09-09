-- 001_create_budgets.sql
-- Creates the budgets table

CREATE TABLE IF NOT EXISTS budgets (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    period_start TEXT NOT NULL,  -- YYYY-MM-DD
    period_end TEXT NOT NULL,    -- YYYY-MM-DD
    created_at TEXT NOT NULL DEFAULT (datetime('now')),
    updated_at TEXT NOT NULL DEFAULT (datetime('now'))
);


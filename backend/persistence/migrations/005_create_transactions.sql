-- 004_create_transactions.sql
-- Creates the transactions table

CREATE TABLE IF NOT EXISTS transactions (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    project_id INTEGER NOT NULL,
    category_id INTEGER,
    description TEXT NOT NULL CHECK(description <> ''), 
    date TEXT NOT NULL CHECK(date <> ''),                       -- YYYY-MM-DD
    amount INTEGER NOT NULL,                  -- in cents
    expense_type BOOLEAN NOT NULL, -- True -> Withdrawl | False -> Deposit
    notes TEXT,
    created_at TEXT NOT NULL DEFAULT (datetime('now')),
    updated_at TEXT NOT NULL DEFAULT (datetime('now')),
    FOREIGN KEY (project_id) REFERENCES projects(id) ON DELETE CASCADE,
    FOREIGN KEY (category_id) REFERENCES categories(id) ON DELETE SET NULL
);


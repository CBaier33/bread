-- 004_create_transactions.sql
-- Creates the transactions table

CREATE TABLE IF NOT EXISTS transactions (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    budget_id INTEGER NOT NULL,
    group_id INTEGER,
    category_id INTEGER,                      -- normalized FK
    description TEXT NOT NULL,
    date TEXT NOT NULL,                       -- YYYY-MM-DD
    amount INTEGER NOT NULL,                  -- in cents
    tags TEXT,
    notes TEXT,
    created_at TEXT NOT NULL DEFAULT (datetime('now')),
    updated_at TEXT NOT NULL DEFAULT (datetime('now')),
    FOREIGN KEY (budget_id) REFERENCES budgets(id) ON DELETE CASCADE
    FOREIGN KEY (group_id) REFERENCES groups(id) ON DELETE SET NULL
    FOREIGN KEY (category_id) REFERENCES categories(id) ON DELETE SET NULL
);


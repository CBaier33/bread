-- 001_create_projects.sql
-- Creates the projects table

CREATE TABLE IF NOT EXISTS projects (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL CHECK(name <> ''),
    description TEXT,
    currency TEXT NOT NULL CHECK(currency <> ''),
    created_at TEXT NOT NULL DEFAULT (datetime('now')),
    updated_at TEXT NOT NULL DEFAULT (datetime('now'))
);

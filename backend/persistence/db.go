package persistence

import (
    "database/sql"
    _ "github.com/mattn/go-sqlite3"
    "log"
)

var DB *sql.DB

func InitDB() {
    var err error
    DB, err = sql.Open("sqlite3", "transactions.db?_foreign_keys=on")
    if err != nil {
        log.Fatal(err)
    }

    createTable := `
    CREATE TABLE IF NOT EXISTS transactions (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        description TEXT NOT NULL,
        amount REAL NOT NULL,
				direction TEXT NOT NULL,
        category TEXT,
        created_at TEXT NOT NULL
    );`
    _, err = DB.Exec(createTable)
    if err != nil {
        log.Fatal(err)
    }
}


package persistence

import (
    "database/sql"
    _ "github.com/mattn/go-sqlite3"
    "log"
)

var DB *sql.DB

func InitDB() {
    var err error
    DB, err = sql.Open("sqlite3", "bread.db?_foreign_keys=on")
    if err != nil {
        log.Fatal(err)
    }

		if err := RunMigrations(DB); err != nil {
			log.Fatal(err)
		}

}


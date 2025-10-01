package persistence

import (
    "database/sql"
    _ "github.com/mattn/go-sqlite3"
    "log"
)

type runner interface {
    Exec(query string, args ...any) (sql.Result, error)
    Query(query string, args ...any) (*sql.Rows, error)
		QueryRow(query string, args...any) *sql.Row
}

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


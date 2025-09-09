package persistence

import (
	"database/sql"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func SetupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("sqlite3", ":memory:?_foreign_keys=on")
	if err != nil {
		t.Fatal(err)
	}
	if err := RunMigrations(db); err != nil {
		t.Fatal(err)
	}
	return db
}


package persistence

import (
	"database/sql"
	"embed"
	"fmt"
	"log"
	"sort"
)

// Embed all SQL files in migrations/ folder
//go:embed migrations/*.sql
var migrationFiles embed.FS

func RunMigrations(db *sql.DB) error {
	// Ensure schema_migrations exists
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS schema_migrations (version TEXT PRIMARY KEY)`)
	if err != nil {
		return fmt.Errorf("create schema_migrations: %w", err)
	}

	// Get already applied migrations
	applied := make(map[string]bool)
	rows, err := db.Query(`SELECT version FROM schema_migrations`)
	if err != nil {
		return fmt.Errorf("read applied migrations: %w", err)
	}
	defer rows.Close()
	for rows.Next() {
		var v string
		if err := rows.Scan(&v); err != nil {
			return fmt.Errorf("scan applied migration: %w", err)
		}
		applied[v] = true
	}

	// Read migration filenames from embedded FS
	entries, err := migrationFiles.ReadDir("migrations")
	if err != nil {
		return fmt.Errorf("read embedded migrations: %w", err)
	}

	names := []string{}
	for _, e := range entries {
		if !e.IsDir() && len(e.Name()) > 4 && e.Name()[len(e.Name())-4:] == ".sql" {
			names = append(names, e.Name())
		}
	}
	sort.Strings(names) // run in order

	// Apply migrations
	for _, name := range names {
		if applied[name] {
			continue
		}

		content, err := migrationFiles.ReadFile("migrations/" + name)
		if err != nil {
			return fmt.Errorf("read migration %s: %w", name, err)
		}

		if _, err := db.Exec(string(content)); err != nil {
			return fmt.Errorf("exec migration %s: %w", name, err)
		}

		if _, err := db.Exec(`INSERT INTO schema_migrations(version) VALUES(?)`, name); err != nil {
			return fmt.Errorf("record migration %s: %w", name, err)
		}

		log.Printf("Applied migration: %s", name)
	}

	return nil
}


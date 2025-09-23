package persistence

import (
	"bread/backend/models"
	"testing"
	"time"
)

func TestProjectPersistence(t *testing.T) {
	// Setup a fresh in-memory DB
	db := SetupTestDB(t)
	DB = db // make sure the package-level DB is set

	now := time.Now().UTC().Format(time.RFC3339)

	// --- Create ---
	p := models.Project{
		Name:        "Test Project",
		Description: "CRUD test",
		Currency:    "USD",
		CreatedAt:   now,
		UpdatedAt:   now,
	}
	id, err := InsertProject(p)
	if err != nil {
		t.Fatalf("InsertProject failed: %v", err)
	}
	if id == 0 {
		t.Fatalf("expected non-zero ID after insert")
	}

	// --- Read ---
	got, err := GetProject(id)
	if err != nil {
		t.Fatalf("GetProject failed: %v", err)
	}
	if got.Name != p.Name {
		t.Errorf("expected name=%s, got=%s", p.Name, got.Name)
	}

	// --- List ---
	list, err := ListProjects()
	if err != nil {
		t.Fatalf("ListProjects failed: %v", err)
	}
	if len(list) != 1 {
		t.Errorf("expected 1 project, got %d", len(list))
	}

	// --- Update ---
	updated := got
	updated.Name = "Updated Project"
	updated.Description = "Updated description"
	updated.Currency = "EUR"
	updated.UpdatedAt = time.Now().UTC().Format(time.RFC3339)
	if err := UpdateProject(updated); err != nil {
		t.Fatalf("UpdateProject failed: %v", err)
	}

	got2, err := GetProject(id)
	if err != nil {
		t.Fatalf("GetProject after update failed: %v", err)
	}
	if got2.Name != updated.Name {
		t.Errorf("expected updated name=%s, got=%s", updated.Name, got2.Name)
	}

	// --- Delete ---
	if err := DeleteProject(id); err != nil {
		t.Fatalf("DeleteProject failed: %v", err)
	}
	_, err = GetProject(id)
	if err == nil {
		t.Errorf("expected error after deleting project, got nil")
	}
}


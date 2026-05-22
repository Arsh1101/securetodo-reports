package repositories

import (
	"context"
	"database/sql"
	"testing"

	_ "modernc.org/sqlite"
)

func setupTestDB(t *testing.T) *sql.DB {
	t.Helper()

	db, err := sql.Open("sqlite", ":memory:")
	if err != nil {
		t.Fatalf("open sqlite test db: %v", err)
	}

	if err := InitTodoSchema(context.Background(), db); err != nil {
		t.Fatalf("init schema: %v", err)
	}

	return db
}

func TestSQLiteTodoRepositoryCreateAndList(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	repo := NewSQLiteTodoRepository(db)

	created, err := repo.Create(context.Background(), "Write Terraform demo")
	if err != nil {
		t.Fatalf("create todo: %v", err)
	}

	if created.ID == 0 {
		t.Fatal("expected created todo to have ID")
	}

	todos, err := repo.List(context.Background())
	if err != nil {
		t.Fatalf("list todos: %v", err)
	}

	if len(todos) != 1 {
		t.Fatalf("expected 1 todo, got %d", len(todos))
	}

	if todos[0].Title != "Write Terraform demo" {
		t.Fatalf("unexpected title: %s", todos[0].Title)
	}
}

func TestSQLiteTodoRepositoryComplete(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	repo := NewSQLiteTodoRepository(db)

	created, err := repo.Create(context.Background(), "Complete me")
	if err != nil {
		t.Fatalf("create todo: %v", err)
	}

	completed, err := repo.Complete(context.Background(), created.ID)
	if err != nil {
		t.Fatalf("complete todo: %v", err)
	}

	if !completed.IsCompleted {
		t.Fatal("expected todo to be completed")
	}
}

package services

import (
	"context"
	"testing"

	"securetodo-reports/app/internal/testutil"
)

func TestTodoServiceCreateTodoSuccessfully(t *testing.T) {
	repo := testutil.NewFakeTodoRepository()
	service := NewTodoService(repo)

	todo, err := service.CreateTodo(context.Background(), "Learn Terraform")
	if err != nil {
		t.Fatalf("create todo: %v", err)
	}

	if todo.ID == 0 {
		t.Fatal("expected todo ID to be set")
	}

	if todo.Title != "Learn Terraform" {
		t.Fatalf("unexpected title: %s", todo.Title)
	}
}

func TestTodoServiceRejectsInvalidTodoTitle(t *testing.T) {
	repo := testutil.NewFakeTodoRepository()
	service := NewTodoService(repo)

	_, err := service.CreateTodo(context.Background(), "   ")
	if err == nil {
		t.Fatal("expected error for invalid title")
	}
}

func TestTodoServiceCompleteTodo(t *testing.T) {
	repo := testutil.NewFakeTodoRepository()
	service := NewTodoService(repo)

	todo, err := service.CreateTodo(context.Background(), "Finish app")
	if err != nil {
		t.Fatalf("create todo: %v", err)
	}

	completed, err := service.CompleteTodo(context.Background(), todo.ID)
	if err != nil {
		t.Fatalf("complete todo: %v", err)
	}

	if !completed.IsCompleted {
		t.Fatal("expected todo to be completed")
	}
}

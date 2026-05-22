package handlers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"securetodo-reports/app/internal/services"
	"securetodo-reports/app/internal/testutil"
)

func TestTodoHandlerListTodosReturnsOK(t *testing.T) {
	repo := testutil.NewFakeTodoRepository()
	service := services.NewTodoService(repo)
	handler := NewTodoHandler(service)

	req := httptest.NewRequest(http.MethodGet, "/todos", nil)
	rec := httptest.NewRecorder()

	handler.HandleTodos(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rec.Code)
	}

	if !strings.Contains(rec.Body.String(), "SecureTodo Reports") {
		t.Fatal("expected page title in response")
	}
}

func TestTodoHandlerCreateTodoRedirects(t *testing.T) {
	repo := testutil.NewFakeTodoRepository()
	service := services.NewTodoService(repo)
	handler := NewTodoHandler(service)

	body := strings.NewReader("title=Learn+Terraform")
	req := httptest.NewRequest(http.MethodPost, "/todos", body)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	rec := httptest.NewRecorder()

	handler.HandleTodos(rec, req)

	if rec.Code != http.StatusSeeOther {
		t.Fatalf("expected redirect, got %d", rec.Code)
	}

	if len(repo.Todos) != 1 {
		t.Fatalf("expected 1 todo, got %d", len(repo.Todos))
	}
}

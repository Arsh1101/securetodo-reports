package handlers

import (
	"context"
	"fmt"
	"html"
	"net/http"
	"strconv"

	"securetodo-reports/app/internal/models"
)

type TodoService interface {
	CreateTodo(ctx context.Context, title string) (models.Todo, error)
	ListTodos(ctx context.Context) ([]models.Todo, error)
	CompleteTodo(ctx context.Context, id int64) (models.Todo, error)
	DeleteTodo(ctx context.Context, id int64) error
}

type TodoHandler struct {
	service TodoService
}

func NewTodoHandler(service TodoService) *TodoHandler {
	return &TodoHandler{service: service}
}

func (h *TodoHandler) HandleTodos(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.listTodos(w, r)
	case http.MethodPost:
		h.createTodo(w, r)
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *TodoHandler) CompleteTodo(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	id, err := strconv.ParseInt(r.FormValue("id"), 10, 64)
	if err != nil {
		http.Error(w, "invalid todo id", http.StatusBadRequest)
		return
	}

	if _, err := h.service.CompleteTodo(r.Context(), id); err != nil {
		http.Error(w, "failed to complete todo", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/todos", http.StatusSeeOther)
}

func (h *TodoHandler) createTodo(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "invalid form", http.StatusBadRequest)
		return
	}

	title := r.FormValue("title")

	if _, err := h.service.CreateTodo(r.Context(), title); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	http.Redirect(w, r, "/todos", http.StatusSeeOther)
}

func (h *TodoHandler) listTodos(w http.ResponseWriter, r *http.Request) {
	todos, err := h.service.ListTodos(r.Context())
	if err != nil {
		http.Error(w, "failed to list todos", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	fmt.Fprintln(w, "<!doctype html><html><head><title>SecureTodo Reports</title><link rel='stylesheet' href='/static/css/styles.css'></head><body>")
	fmt.Fprintln(w, "<main>")
	fmt.Fprintln(w, "<h1>SecureTodo Reports</h1>")
	fmt.Fprintln(w, "<form method='POST' action='/todos'>")
	fmt.Fprintln(w, "<input name='title' placeholder='New todo' required>")
	fmt.Fprintln(w, "<button type='submit'>Add Todo</button>")
	fmt.Fprintln(w, "</form>")

	fmt.Fprintln(w, "<form method='POST' action='/reports/generate'>")
	fmt.Fprintln(w, "<button type='submit'>Generate JSON Report</button>")
	fmt.Fprintln(w, "</form>")

	fmt.Fprintln(w, "<p><a href='/reports'>View reports</a></p>")
	fmt.Fprintln(w, "<ul>")

	for _, todo := range todos {
		status := "pending"
		if todo.IsCompleted {
			status = "completed"
		}

		fmt.Fprintf(w, "<li><strong>%s</strong> - %s", html.EscapeString(todo.Title), status)

		if !todo.IsCompleted {
			fmt.Fprintf(w, `<form method="POST" action="/todos/complete" style="display:inline;"><input type="hidden" name="id" value="%d"><button type="submit">Complete</button></form>`, todo.ID)
		}

		fmt.Fprintln(w, "</li>")
	}

	fmt.Fprintln(w, "</ul>")
	fmt.Fprintln(w, "</main></body></html>")
}

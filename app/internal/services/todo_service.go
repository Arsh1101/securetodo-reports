package services

import (
	"context"
	"strings"

	"securetodo-reports/app/internal/models"
	"securetodo-reports/app/internal/repositories"
	"securetodo-reports/app/internal/validation"
)

type TodoService struct {
	repo repositories.TodoRepository
}

func NewTodoService(repo repositories.TodoRepository) *TodoService {
	return &TodoService{repo: repo}
}

func (s *TodoService) CreateTodo(ctx context.Context, title string) (models.Todo, error) {
	if err := validation.ValidateTodoTitle(title); err != nil {
		return models.Todo{}, err
	}

	return s.repo.Create(ctx, strings.TrimSpace(title))
}

func (s *TodoService) ListTodos(ctx context.Context) ([]models.Todo, error) {
	return s.repo.List(ctx)
}

func (s *TodoService) CompleteTodo(ctx context.Context, id int64) (models.Todo, error) {
	return s.repo.Complete(ctx, id)
}

func (s *TodoService) DeleteTodo(ctx context.Context, id int64) error {
	return s.repo.Delete(ctx, id)
}

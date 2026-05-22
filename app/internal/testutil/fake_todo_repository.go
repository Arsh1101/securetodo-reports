package testutil

import (
	"context"
	"errors"
	"time"

	"securetodo-reports/app/internal/models"
)

type FakeTodoRepository struct {
	Todos  []models.Todo
	NextID int64
}

func NewFakeTodoRepository() *FakeTodoRepository {
	return &FakeTodoRepository{
		Todos:  []models.Todo{},
		NextID: 1,
	}
}

func (r *FakeTodoRepository) Create(ctx context.Context, title string) (models.Todo, error) {
	now := time.Now().UTC()

	todo := models.Todo{
		ID:          r.NextID,
		Title:       title,
		IsCompleted: false,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	r.NextID++
	r.Todos = append(r.Todos, todo)

	return todo, nil
}

func (r *FakeTodoRepository) List(ctx context.Context) ([]models.Todo, error) {
	return r.Todos, nil
}

func (r *FakeTodoRepository) Complete(ctx context.Context, id int64) (models.Todo, error) {
	for i, todo := range r.Todos {
		if todo.ID == id {
			r.Todos[i].IsCompleted = true
			r.Todos[i].UpdatedAt = time.Now().UTC()
			return r.Todos[i], nil
		}
	}

	return models.Todo{}, errors.New("todo not found")
}

func (r *FakeTodoRepository) Delete(ctx context.Context, id int64) error {
	for i, todo := range r.Todos {
		if todo.ID == id {
			r.Todos = append(r.Todos[:i], r.Todos[i+1:]...)
			return nil
		}
	}

	return errors.New("todo not found")
}

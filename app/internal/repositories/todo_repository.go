package repositories

import (
	"context"

	"securetodo-reports/app/internal/models"
)

type TodoRepository interface {
	Create(ctx context.Context, title string) (models.Todo, error)
	List(ctx context.Context) ([]models.Todo, error)
	Complete(ctx context.Context, id int64) (models.Todo, error)
	Delete(ctx context.Context, id int64) error
}

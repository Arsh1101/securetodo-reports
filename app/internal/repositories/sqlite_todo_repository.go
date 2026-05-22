package repositories

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"securetodo-reports/app/internal/models"

	_ "modernc.org/sqlite"
)

var ErrTodoNotFound = errors.New("todo not found")

type SQLiteTodoRepository struct {
	db *sql.DB
}

func NewSQLiteTodoRepository(db *sql.DB) *SQLiteTodoRepository {
	return &SQLiteTodoRepository{db: db}
}

func InitTodoSchema(ctx context.Context, db *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS todos (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		is_completed INTEGER NOT NULL DEFAULT 0,
		created_at TEXT NOT NULL,
		updated_at TEXT NOT NULL
	);`

	_, err := db.ExecContext(ctx, query)
	return err
}

func (r *SQLiteTodoRepository) Create(ctx context.Context, title string) (models.Todo, error) {
	now := time.Now().UTC()

	result, err := r.db.ExecContext(
		ctx,
		`INSERT INTO todos (title, is_completed, created_at, updated_at) VALUES (?, 0, ?, ?)`,
		title,
		now.Format(time.RFC3339),
		now.Format(time.RFC3339),
	)
	if err != nil {
		return models.Todo{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return models.Todo{}, err
	}

	return models.Todo{
		ID:          id,
		Title:       title,
		IsCompleted: false,
		CreatedAt:   now,
		UpdatedAt:   now,
	}, nil
}

func (r *SQLiteTodoRepository) List(ctx context.Context) ([]models.Todo, error) {
	rows, err := r.db.QueryContext(
		ctx,
		`SELECT id, title, is_completed, created_at, updated_at FROM todos ORDER BY id DESC`,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todos []models.Todo

	for rows.Next() {
		var todo models.Todo
		var completed int
		var createdAt string
		var updatedAt string

		if err := rows.Scan(&todo.ID, &todo.Title, &completed, &createdAt, &updatedAt); err != nil {
			return nil, err
		}

		todo.IsCompleted = completed == 1
		todo.CreatedAt, _ = time.Parse(time.RFC3339, createdAt)
		todo.UpdatedAt, _ = time.Parse(time.RFC3339, updatedAt)

		todos = append(todos, todo)
	}

	return todos, rows.Err()
}

func (r *SQLiteTodoRepository) Complete(ctx context.Context, id int64) (models.Todo, error) {
	now := time.Now().UTC()

	result, err := r.db.ExecContext(
		ctx,
		`UPDATE todos SET is_completed = 1, updated_at = ? WHERE id = ?`,
		now.Format(time.RFC3339),
		id,
	)
	if err != nil {
		return models.Todo{}, err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return models.Todo{}, err
	}

	if affected == 0 {
		return models.Todo{}, ErrTodoNotFound
	}

	row := r.db.QueryRowContext(
		ctx,
		`SELECT id, title, is_completed, created_at, updated_at FROM todos WHERE id = ?`,
		id,
	)

	var todo models.Todo
	var completed int
	var createdAt string
	var updatedAt string

	if err := row.Scan(&todo.ID, &todo.Title, &completed, &createdAt, &updatedAt); err != nil {
		return models.Todo{}, err
	}

	todo.IsCompleted = completed == 1
	todo.CreatedAt, _ = time.Parse(time.RFC3339, createdAt)
	todo.UpdatedAt, _ = time.Parse(time.RFC3339, updatedAt)

	return todo, nil
}

func (r *SQLiteTodoRepository) Delete(ctx context.Context, id int64) error {
	result, err := r.db.ExecContext(ctx, `DELETE FROM todos WHERE id = ?`, id)
	if err != nil {
		return err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if affected == 0 {
		return ErrTodoNotFound
	}

	return nil
}

package validation

import (
	"errors"
	"strings"
)

var (
	ErrEmptyTodoTitle = errors.New("todo title cannot be empty")
	ErrTodoTitleTooLong = errors.New("todo title cannot be longer than 120 characters")
)

func ValidateTodoTitle(title string) error {
	trimmed := strings.TrimSpace(title)

	if trimmed == "" {
		return ErrEmptyTodoTitle
	}

	if len(trimmed) > 120 {
		return ErrTodoTitleTooLong
	}

	return nil
}

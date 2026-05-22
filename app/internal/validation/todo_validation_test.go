package validation

import (
	"errors"
	"strings"
	"testing"
)

func TestValidateTodoTitleRejectsEmptyTitle(t *testing.T) {
	err := ValidateTodoTitle("")

	if !errors.Is(err, ErrEmptyTodoTitle) {
		t.Fatalf("expected ErrEmptyTodoTitle, got %v", err)
	}
}

func TestValidateTodoTitleRejectsWhitespaceOnlyTitle(t *testing.T) {
	err := ValidateTodoTitle("     ")

	if !errors.Is(err, ErrEmptyTodoTitle) {
		t.Fatalf("expected ErrEmptyTodoTitle, got %v", err)
	}
}

func TestValidateTodoTitleRejectsVeryLongTitle(t *testing.T) {
	longTitle := strings.Repeat("a", 121)

	err := ValidateTodoTitle(longTitle)

	if !errors.Is(err, ErrTodoTitleTooLong) {
		t.Fatalf("expected ErrTodoTitleTooLong, got %v", err)
	}
}

func TestValidateTodoTitleAcceptsValidTitle(t *testing.T) {
	err := ValidateTodoTitle("Learn Terraform modules")

	if err != nil {
		t.Fatalf("expected valid title, got %v", err)
	}
}

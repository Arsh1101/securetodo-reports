package models

import "testing"

func TestTodoDefaultCompletionState(t *testing.T) {
	todo := Todo{Title: "Learn Terraform"}

	if todo.IsCompleted {
		t.Fatal("expected new todo to be incomplete by default")
	}
}

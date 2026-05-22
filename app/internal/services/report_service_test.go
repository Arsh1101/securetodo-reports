package services

import (
	"context"
	"testing"

	"securetodo-reports/app/internal/testutil"
)

func TestReportServiceGenerateAndSaveReport(t *testing.T) {
	todoRepo := testutil.NewFakeTodoRepository()
	reportRepo := testutil.NewFakeReportRepository()

	todoService := NewTodoService(todoRepo)
	reportService := NewReportService(todoRepo, reportRepo)

	first, err := todoService.CreateTodo(context.Background(), "Learn Go TDD")
	if err != nil {
		t.Fatalf("create first todo: %v", err)
	}

	_, err = todoService.CreateTodo(context.Background(), "Build Terraform demo")
	if err != nil {
		t.Fatalf("create second todo: %v", err)
	}

	_, err = todoService.CompleteTodo(context.Background(), first.ID)
	if err != nil {
		t.Fatalf("complete todo: %v", err)
	}

	report, err := reportService.GenerateAndSaveReport(context.Background())
	if err != nil {
		t.Fatalf("generate report: %v", err)
	}

	if report.Summary.Total != 2 {
		t.Fatalf("expected total 2, got %d", report.Summary.Total)
	}

	if report.Summary.Completed != 1 {
		t.Fatalf("expected completed 1, got %d", report.Summary.Completed)
	}

	if report.Summary.Pending != 1 {
		t.Fatalf("expected pending 1, got %d", report.Summary.Pending)
	}

	if len(reportRepo.SavedFiles) != 1 {
		t.Fatalf("expected 1 saved report, got %d", len(reportRepo.SavedFiles))
	}
}

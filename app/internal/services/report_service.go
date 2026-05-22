package services

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"securetodo-reports/app/internal/models"
	"securetodo-reports/app/internal/repositories"
)

type ReportService struct {
	todoRepo   repositories.TodoRepository
	reportRepo repositories.ReportRepository
}

func NewReportService(todoRepo repositories.TodoRepository, reportRepo repositories.ReportRepository) *ReportService {
	return &ReportService{
		todoRepo:   todoRepo,
		reportRepo: reportRepo,
	}
}

func (s *ReportService) GenerateAndSaveReport(ctx context.Context) (models.TodoReport, error) {
	todos, err := s.todoRepo.List(ctx)
	if err != nil {
		return models.TodoReport{}, err
	}

	report := buildTodoReport(todos)

	data, err := json.MarshalIndent(report, "", "  ")
	if err != nil {
		return models.TodoReport{}, err
	}

	if err := s.reportRepo.Save(ctx, report.FileName, data); err != nil {
		return models.TodoReport{}, err
	}

	return report, nil
}

func (s *ReportService) ListReports(ctx context.Context) ([]string, error) {
	return s.reportRepo.List(ctx)
}

func buildTodoReport(todos []models.Todo) models.TodoReport {
	completed := 0

	for _, todo := range todos {
		if todo.IsCompleted {
			completed++
		}
	}

	now := time.Now().UTC()

	return models.TodoReport{
		FileName:    fmt.Sprintf("todo-report-%s.json", now.Format("20060102-150405")),
		GeneratedAt: now,
		Summary: models.ReportSummary{
			Total:     len(todos),
			Completed: completed,
			Pending:   len(todos) - completed,
		},
		Items: todos,
	}
}

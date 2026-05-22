package handlers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"securetodo-reports/app/internal/services"
	"securetodo-reports/app/internal/testutil"
)

func TestReportHandlerGenerateReportRedirects(t *testing.T) {
	todoRepo := testutil.NewFakeTodoRepository()
	reportRepo := testutil.NewFakeReportRepository()

	reportService := services.NewReportService(todoRepo, reportRepo)
	handler := NewReportHandler(reportService)

	req := httptest.NewRequest(http.MethodPost, "/reports/generate", nil)
	rec := httptest.NewRecorder()

	handler.GenerateReport(rec, req)

	if rec.Code != http.StatusSeeOther {
		t.Fatalf("expected redirect, got %d", rec.Code)
	}
}

func TestReportHandlerListReportsReturnsOK(t *testing.T) {
	todoRepo := testutil.NewFakeTodoRepository()
	reportRepo := testutil.NewFakeReportRepository()
	reportRepo.SavedFiles["todo-report-test.json"] = []byte(`{}`)

	reportService := services.NewReportService(todoRepo, reportRepo)
	handler := NewReportHandler(reportService)

	req := httptest.NewRequest(http.MethodGet, "/reports", nil)
	rec := httptest.NewRecorder()

	handler.ListReports(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rec.Code)
	}

	if !strings.Contains(rec.Body.String(), "todo-report-test.json") {
		t.Fatal("expected report file name in response")
	}
}

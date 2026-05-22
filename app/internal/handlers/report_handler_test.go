package handlers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"securetodo-reports/app/internal/services"
	"securetodo-reports/app/internal/testutil"
)

func TestReportHandlerDownloadReportReturnsJSON(t *testing.T) {
	todoRepo := testutil.NewFakeTodoRepository()
	reportRepo := testutil.NewFakeReportRepository()
	reportRepo.SavedFiles["todo-report-test.json"] = []byte(`{"ok":true}`)

	reportService := services.NewReportService(todoRepo, reportRepo)
	handler := NewReportHandler(reportService)

	req := httptest.NewRequest(http.MethodGet, "/reports/download?file=todo-report-test.json", nil)
	rec := httptest.NewRecorder()

	handler.DownloadReport(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rec.Code)
	}

	if rec.Header().Get("Content-Type") != "application/json" {
		t.Fatalf("expected application/json, got %s", rec.Header().Get("Content-Type"))
	}

	if !strings.Contains(rec.Header().Get("Content-Disposition"), "todo-report-test.json") {
		t.Fatalf("expected content disposition to include file name")
	}

	if rec.Body.String() != `{"ok":true}` {
		t.Fatalf("unexpected body: %s", rec.Body.String())
	}
}

func TestReportHandlerDownloadReportWithoutFileReturnsBadRequest(t *testing.T) {
	todoRepo := testutil.NewFakeTodoRepository()
	reportRepo := testutil.NewFakeReportRepository()

	reportService := services.NewReportService(todoRepo, reportRepo)
	handler := NewReportHandler(reportService)

	req := httptest.NewRequest(http.MethodGet, "/reports/download", nil)
	rec := httptest.NewRecorder()

	handler.DownloadReport(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("expected status 400, got %d", rec.Code)
	}
}

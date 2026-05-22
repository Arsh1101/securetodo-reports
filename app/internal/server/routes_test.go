package server

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"securetodo-reports/app/internal/handlers"
	"securetodo-reports/app/internal/services"
	"securetodo-reports/app/internal/testutil"
)

func TestRouterHealthRoute(t *testing.T) {
	todoRepo := testutil.NewFakeTodoRepository()
	reportRepo := testutil.NewFakeReportRepository()

	todoService := services.NewTodoService(todoRepo)
	reportService := services.NewReportService(todoRepo, reportRepo)

	router := NewRouter(
		handlers.NewTodoHandler(todoService),
		handlers.NewReportHandler(reportService),
	)

	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rec.Code)
	}
}

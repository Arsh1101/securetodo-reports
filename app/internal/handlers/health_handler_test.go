package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthHandlerReturnsOK(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	rec := httptest.NewRecorder()

	HealthHandler(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rec.Code)
	}

	if rec.Body.String() != "ok" {
		t.Fatalf("expected body ok, got %q", rec.Body.String())
	}
}

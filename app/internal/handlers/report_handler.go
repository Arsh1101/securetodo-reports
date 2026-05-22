package handlers

import (
	"context"
	"fmt"
	"html"
	"net/http"

	"securetodo-reports/app/internal/models"
)

type ReportService interface {
	GenerateAndSaveReport(ctx context.Context) (models.TodoReport, error)
	ListReports(ctx context.Context) ([]string, error)
}

type ReportHandler struct {
	service ReportService
}

func NewReportHandler(service ReportService) *ReportHandler {
	return &ReportHandler{service: service}
}

func (h *ReportHandler) GenerateReport(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if _, err := h.service.GenerateAndSaveReport(r.Context()); err != nil {
		http.Error(w, "failed to generate report", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/reports", http.StatusSeeOther)
}

func (h *ReportHandler) ListReports(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	reports, err := h.service.ListReports(r.Context())
	if err != nil {
		http.Error(w, "failed to list reports", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	fmt.Fprintln(w, "<!doctype html><html><head><title>Reports</title><link rel='stylesheet' href='/static/css/styles.css'></head><body>")
	fmt.Fprintln(w, "<main>")
	fmt.Fprintln(w, "<h1>Generated Reports</h1>")
	fmt.Fprintln(w, "<p><a href='/todos'>Back to todos</a></p>")
	fmt.Fprintln(w, "<ul>")

	for _, report := range reports {
		fmt.Fprintf(w, "<li>%s</li>", html.EscapeString(report))
	}

	fmt.Fprintln(w, "</ul>")
	fmt.Fprintln(w, "</main></body></html>")
}

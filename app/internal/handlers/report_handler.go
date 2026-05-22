package handlers

import (
	"context"
	"fmt"
	"html"
	"net/http"
	"net/url"

	"securetodo-reports/app/internal/models"
)

type ReportService interface {
	GenerateAndSaveReport(ctx context.Context) (models.TodoReport, error)
	ListReports(ctx context.Context) ([]string, error)
	ReadReport(ctx context.Context, fileName string) ([]byte, error)
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
		escapedReport := html.EscapeString(report)
		downloadURL := "/reports/download?file=" + url.QueryEscape(report)

		fmt.Fprintf(
			w,
			`<li>%s - <a href="%s">Download</a></li>`,
			escapedReport,
			downloadURL,
		)
	}

	fmt.Fprintln(w, "</ul>")
	fmt.Fprintln(w, "</main></body></html>")
}

func (h *ReportHandler) DownloadReport(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	fileName := r.URL.Query().Get("file")
	if fileName == "" {
		http.Error(w, "missing report file name", http.StatusBadRequest)
		return
	}

	data, err := h.service.ReadReport(r.Context(), fileName)
	if err != nil {
		http.Error(w, "report not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Disposition", fmt.Sprintf(`attachment; filename="%s"`, fileName))
	w.WriteHeader(http.StatusOK)

	_, _ = w.Write(data)
}

package server

import (
	"net/http"

	"securetodo-reports/app/internal/handlers"
)

func NewRouter(todoHandler *handlers.TodoHandler, reportHandler *handlers.ReportHandler) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/todos", http.StatusSeeOther)
	})

	mux.HandleFunc("/health", handlers.HealthHandler)
	mux.HandleFunc("/todos", todoHandler.HandleTodos)
	mux.HandleFunc("/todos/complete", todoHandler.CompleteTodo)
	mux.HandleFunc("/reports", reportHandler.ListReports)
	mux.HandleFunc("/reports/generate", reportHandler.GenerateReport)
	mux.HandleFunc("/reports/download", reportHandler.DownloadReport)

	fileServer := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fileServer))

	return mux
}

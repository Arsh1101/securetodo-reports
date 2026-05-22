package main

import (
	"context"
	"database/sql"
	"log"
	"os"
	"path/filepath"

	"securetodo-reports/app/internal/config"
	"securetodo-reports/app/internal/handlers"
	"securetodo-reports/app/internal/repositories"
	"securetodo-reports/app/internal/server"
	"securetodo-reports/app/internal/services"

	_ "modernc.org/sqlite"
)

func main() {
	cfg := config.Load()

	if err := os.MkdirAll(filepath.Dir(cfg.DBPath), 0755); err != nil {
		log.Fatalf("create data directory: %v", err)
	}

	db, err := sql.Open("sqlite", cfg.DBPath)
	if err != nil {
		log.Fatalf("open sqlite database: %v", err)
	}
	defer db.Close()

	if err := repositories.InitTodoSchema(context.Background(), db); err != nil {
		log.Fatalf("init todo schema: %v", err)
	}

	todoRepo := repositories.NewSQLiteTodoRepository(db)
	reportRepo := repositories.NewLocalReportRepository(cfg.ReportsDir)

	todoService := services.NewTodoService(todoRepo)
	reportService := services.NewReportService(todoRepo, reportRepo)

	todoHandler := handlers.NewTodoHandler(todoService)
	reportHandler := handlers.NewReportHandler(reportService)

	router := server.NewRouter(todoHandler, reportHandler)

	appServer := server.New(cfg.Addr, router)

	log.Printf("starting server on %s", cfg.Addr)

	if err := appServer.ListenAndServe(); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}

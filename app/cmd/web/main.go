package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"

	appconfig "securetodo-reports/app/internal/config"
	"securetodo-reports/app/internal/handlers"
	"securetodo-reports/app/internal/repositories"
	"securetodo-reports/app/internal/server"
	"securetodo-reports/app/internal/services"

	_ "modernc.org/sqlite"
)

func main() {
	cfg := appconfig.Load()

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

	reportRepo, err := buildReportRepository(context.Background(), cfg)
	if err != nil {
		log.Fatalf("create report repository: %v", err)
	}

	todoService := services.NewTodoService(todoRepo)
	reportService := services.NewReportService(todoRepo, reportRepo)

	todoHandler := handlers.NewTodoHandler(todoService)
	reportHandler := handlers.NewReportHandler(reportService)

	router := server.NewRouter(todoHandler, reportHandler)

	appServer := server.New(cfg.Addr, router)

	log.Printf("starting server on %s", cfg.Addr)
	log.Printf("report storage mode: %s", cfg.ReportStorage)

	if err := appServer.ListenAndServe(); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}

func buildReportRepository(ctx context.Context, cfg appconfig.Config) (repositories.ReportRepository, error) {
	switch strings.ToLower(cfg.ReportStorage) {
	case "local":
		return repositories.NewLocalReportRepository(cfg.ReportsDir), nil

	case "s3":
		if cfg.S3BucketName == "" {
			return nil, fmt.Errorf("S3_BUCKET_NAME is required when REPORT_STORAGE=s3")
		}

		awsCfg, err := config.LoadDefaultConfig(ctx, config.WithRegion(cfg.AWSRegion))
		if err != nil {
			return nil, fmt.Errorf("load AWS config: %w", err)
		}

		s3Client := s3.NewFromConfig(awsCfg)

		return repositories.NewS3ReportRepository(
			s3Client,
			cfg.S3BucketName,
			cfg.S3ReportsPrefix,
		), nil

	default:
		return nil, fmt.Errorf("unsupported REPORT_STORAGE value: %s", cfg.ReportStorage)
	}
}

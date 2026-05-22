package config

import "testing"

func TestLoadUsesDefaults(t *testing.T) {
	t.Setenv("APP_ADDR", "")
	t.Setenv("DB_PATH", "")
	t.Setenv("REPORTS_DIR", "")
	t.Setenv("REPORT_STORAGE", "")
	t.Setenv("AWS_REGION", "")
	t.Setenv("S3_BUCKET_NAME", "")
	t.Setenv("S3_REPORTS_PREFIX", "")

	cfg := Load()

	if cfg.Addr != ":8080" {
		t.Fatalf("expected default addr, got %s", cfg.Addr)
	}

	if cfg.DBPath == "" {
		t.Fatal("expected default db path")
	}

	if cfg.ReportsDir == "" {
		t.Fatal("expected default reports dir")
	}

	if cfg.ReportStorage != "local" {
		t.Fatalf("expected default report storage local, got %s", cfg.ReportStorage)
	}

	if cfg.AWSRegion != "ca-central-1" {
		t.Fatalf("expected default AWS region ca-central-1, got %s", cfg.AWSRegion)
	}

	if cfg.S3ReportsPrefix != "reports/" {
		t.Fatalf("expected default S3 reports prefix reports/, got %s", cfg.S3ReportsPrefix)
	}
}

func TestLoadUsesEnvironmentVariables(t *testing.T) {
	t.Setenv("APP_ADDR", ":9090")
	t.Setenv("DB_PATH", "./test.db")
	t.Setenv("REPORTS_DIR", "./test-reports")
	t.Setenv("REPORT_STORAGE", "s3")
	t.Setenv("AWS_REGION", "us-west-2")
	t.Setenv("S3_BUCKET_NAME", "my-test-bucket")
	t.Setenv("S3_REPORTS_PREFIX", "todo-reports/")

	cfg := Load()

	if cfg.Addr != ":9090" {
		t.Fatalf("expected env addr, got %s", cfg.Addr)
	}

	if cfg.DBPath != "./test.db" {
		t.Fatalf("expected env db path, got %s", cfg.DBPath)
	}

	if cfg.ReportsDir != "./test-reports" {
		t.Fatalf("expected env reports dir, got %s", cfg.ReportsDir)
	}

	if cfg.ReportStorage != "s3" {
		t.Fatalf("expected env report storage s3, got %s", cfg.ReportStorage)
	}

	if cfg.AWSRegion != "us-west-2" {
		t.Fatalf("expected env AWS region, got %s", cfg.AWSRegion)
	}

	if cfg.S3BucketName != "my-test-bucket" {
		t.Fatalf("expected env bucket name, got %s", cfg.S3BucketName)
	}

	if cfg.S3ReportsPrefix != "todo-reports/" {
		t.Fatalf("expected env prefix, got %s", cfg.S3ReportsPrefix)
	}
}

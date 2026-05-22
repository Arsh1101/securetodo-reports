package config

import "testing"

func TestLoadUsesDefaults(t *testing.T) {
	t.Setenv("APP_ADDR", "")
	t.Setenv("DB_PATH", "")
	t.Setenv("REPORTS_DIR", "")

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
}

func TestLoadUsesEnvironmentVariables(t *testing.T) {
	t.Setenv("APP_ADDR", ":9090")
	t.Setenv("DB_PATH", "./test.db")
	t.Setenv("REPORTS_DIR", "./test-reports")

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
}

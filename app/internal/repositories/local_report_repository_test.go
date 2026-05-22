package repositories

import (
	"context"
	"testing"
)

func TestLocalReportRepositorySaveAndList(t *testing.T) {
	dir := t.TempDir()

	repo := NewLocalReportRepository(dir)

	err := repo.Save(context.Background(), "report.json", []byte(`{"ok":true}`))
	if err != nil {
		t.Fatalf("save report: %v", err)
	}

	reports, err := repo.List(context.Background())
	if err != nil {
		t.Fatalf("list reports: %v", err)
	}

	if len(reports) != 1 {
		t.Fatalf("expected 1 report, got %d", len(reports))
	}

	if reports[0] != "report.json" {
		t.Fatalf("unexpected report name: %s", reports[0])
	}
}

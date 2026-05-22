package repositories

import (
	"context"
	"errors"
	"testing"
)

func TestLocalReportRepositorySaveListAndRead(t *testing.T) {
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

	data, err := repo.Read(context.Background(), "report.json")
	if err != nil {
		t.Fatalf("read report: %v", err)
	}

	if string(data) != `{"ok":true}` {
		t.Fatalf("unexpected report content: %s", string(data))
	}
}

func TestValidateReportFileNameRejectsPathTraversal(t *testing.T) {
	err := ValidateReportFileName("../secret.json")

	if !errors.Is(err, ErrInvalidReportFileName) {
		t.Fatalf("expected ErrInvalidReportFileName, got %v", err)
	}
}

func TestValidateReportFileNameRejectsNonJSONFile(t *testing.T) {
	err := ValidateReportFileName("report.txt")

	if !errors.Is(err, ErrInvalidReportFileName) {
		t.Fatalf("expected ErrInvalidReportFileName, got %v", err)
	}
}

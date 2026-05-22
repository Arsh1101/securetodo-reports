package models

import "testing"

func TestReportSummaryPendingCount(t *testing.T) {
	summary := ReportSummary{
		Total:     3,
		Completed: 1,
		Pending:   2,
	}

	if summary.Pending != summary.Total-summary.Completed {
		t.Fatal("pending count should match total minus completed")
	}
}

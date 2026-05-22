package models

import "time"

type TodoReport struct {
	FileName    string        `json:"file_name"`
	GeneratedAt time.Time     `json:"generated_at"`
	Summary     ReportSummary `json:"summary"`
	Items       []Todo        `json:"items"`
}

type ReportSummary struct {
	Total     int `json:"total"`
	Completed int `json:"completed"`
	Pending   int `json:"pending"`
}

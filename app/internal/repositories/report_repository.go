package repositories

import "context"

type ReportRepository interface {
	Save(ctx context.Context, fileName string, data []byte) error
	List(ctx context.Context) ([]string, error)
	Read(ctx context.Context, fileName string) ([]byte, error)
}

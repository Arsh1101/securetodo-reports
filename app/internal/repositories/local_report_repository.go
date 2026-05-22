package repositories

import (
	"context"
	"os"
	"path/filepath"
	"sort"
)

type LocalReportRepository struct {
	dir string
}

func NewLocalReportRepository(dir string) *LocalReportRepository {
	return &LocalReportRepository{dir: dir}
}

func (r *LocalReportRepository) Save(ctx context.Context, fileName string, data []byte) error {
	if err := os.MkdirAll(r.dir, 0755); err != nil {
		return err
	}

	path := filepath.Join(r.dir, fileName)

	return os.WriteFile(path, data, 0644)
}

func (r *LocalReportRepository) List(ctx context.Context) ([]string, error) {
	entries, err := os.ReadDir(r.dir)
	if err != nil {
		if os.IsNotExist(err) {
			return []string{}, nil
		}
		return nil, err
	}

	var reports []string

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		reports = append(reports, entry.Name())
	}

	sort.Sort(sort.Reverse(sort.StringSlice(reports)))

	return reports, nil
}

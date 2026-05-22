package repositories

import (
	"context"
	"errors"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

var ErrInvalidReportFileName = errors.New("invalid report file name")

type LocalReportRepository struct {
	dir string
}

func NewLocalReportRepository(dir string) *LocalReportRepository {
	return &LocalReportRepository{dir: dir}
}

func (r *LocalReportRepository) Save(ctx context.Context, fileName string, data []byte) error {
	if err := ValidateReportFileName(fileName); err != nil {
		return err
	}

	if err := os.MkdirAll(r.dir, 0755); err != nil {
		return err
	}

	return os.WriteFile(filepath.Join(r.dir, fileName), data, 0644)
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

		if strings.HasSuffix(entry.Name(), ".json") {
			reports = append(reports, entry.Name())
		}
	}

	sort.Sort(sort.Reverse(sort.StringSlice(reports)))

	return reports, nil
}

func (r *LocalReportRepository) Read(ctx context.Context, fileName string) ([]byte, error) {
	if err := ValidateReportFileName(fileName); err != nil {
		return nil, err
	}

	return os.ReadFile(filepath.Join(r.dir, fileName))
}

func ValidateReportFileName(fileName string) error {
	if fileName == "" {
		return ErrInvalidReportFileName
	}

	if filepath.Base(fileName) != fileName {
		return ErrInvalidReportFileName
	}

	if !strings.HasSuffix(fileName, ".json") {
		return ErrInvalidReportFileName
	}

	return nil
}

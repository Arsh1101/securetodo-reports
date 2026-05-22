package testutil

import (
	"context"
	"errors"
)

type FakeReportRepository struct {
	SavedFiles map[string][]byte
}

func NewFakeReportRepository() *FakeReportRepository {
	return &FakeReportRepository{
		SavedFiles: map[string][]byte{},
	}
}

func (r *FakeReportRepository) Save(ctx context.Context, fileName string, data []byte) error {
	r.SavedFiles[fileName] = data
	return nil
}

func (r *FakeReportRepository) List(ctx context.Context) ([]string, error) {
	files := make([]string, 0, len(r.SavedFiles))

	for fileName := range r.SavedFiles {
		files = append(files, fileName)
	}

	return files, nil
}

func (r *FakeReportRepository) Read(ctx context.Context, fileName string) ([]byte, error) {
	data, ok := r.SavedFiles[fileName]
	if !ok {
		return nil, errors.New("report not found")
	}

	return data, nil
}

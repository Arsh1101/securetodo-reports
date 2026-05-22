package repositories

import (
	"bytes"
	"context"
	"io"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

type fakeS3Client struct {
	objects map[string][]byte
}

func newFakeS3Client() *fakeS3Client {
	return &fakeS3Client{
		objects: map[string][]byte{},
	}
}

func (c *fakeS3Client) PutObject(ctx context.Context, input *s3.PutObjectInput, optFns ...func(*s3.Options)) (*s3.PutObjectOutput, error) {
	data, err := io.ReadAll(input.Body)
	if err != nil {
		return nil, err
	}

	c.objects[*input.Key] = data

	return &s3.PutObjectOutput{}, nil
}

func (c *fakeS3Client) GetObject(ctx context.Context, input *s3.GetObjectInput, optFns ...func(*s3.Options)) (*s3.GetObjectOutput, error) {
	data := c.objects[*input.Key]

	return &s3.GetObjectOutput{
		Body: io.NopCloser(bytes.NewReader(data)),
	}, nil
}

func (c *fakeS3Client) ListObjectsV2(ctx context.Context, input *s3.ListObjectsV2Input, optFns ...func(*s3.Options)) (*s3.ListObjectsV2Output, error) {
	var contents []types.Object

	for key := range c.objects {
		contents = append(contents, types.Object{
			Key: aws.String(key),
		})
	}

	return &s3.ListObjectsV2Output{
		Contents: contents,
	}, nil
}

func TestS3ReportRepositorySaveListAndRead(t *testing.T) {
	client := newFakeS3Client()
	repo := NewS3ReportRepository(client, "test-bucket", "reports/")

	err := repo.Save(context.Background(), "todo-report-test.json", []byte(`{"ok":true}`))
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

	if reports[0] != "todo-report-test.json" {
		t.Fatalf("unexpected report name: %s", reports[0])
	}

	data, err := repo.Read(context.Background(), "todo-report-test.json")
	if err != nil {
		t.Fatalf("read report: %v", err)
	}

	if string(data) != `{"ok":true}` {
		t.Fatalf("unexpected report content: %s", string(data))
	}
}

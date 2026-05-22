package repositories

import (
	"bytes"
	"context"
	"io"
	"path"
	"sort"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type S3Client interface {
	PutObject(ctx context.Context, params *s3.PutObjectInput, optFns ...func(*s3.Options)) (*s3.PutObjectOutput, error)
	GetObject(ctx context.Context, params *s3.GetObjectInput, optFns ...func(*s3.Options)) (*s3.GetObjectOutput, error)
	ListObjectsV2(ctx context.Context, params *s3.ListObjectsV2Input, optFns ...func(*s3.Options)) (*s3.ListObjectsV2Output, error)
}

type S3ReportRepository struct {
	client S3Client
	bucket string
	prefix string
}

func NewS3ReportRepository(client S3Client, bucket string, prefix string) *S3ReportRepository {
	if prefix != "" && !strings.HasSuffix(prefix, "/") {
		prefix += "/"
	}

	return &S3ReportRepository{
		client: client,
		bucket: bucket,
		prefix: prefix,
	}
}

func (r *S3ReportRepository) Save(ctx context.Context, fileName string, data []byte) error {
	if err := ValidateReportFileName(fileName); err != nil {
		return err
	}

	_, err := r.client.PutObject(ctx, &s3.PutObjectInput{
		Bucket:      aws.String(r.bucket),
		Key:         aws.String(r.objectKey(fileName)),
		Body:        bytes.NewReader(data),
		ContentType: aws.String("application/json"),
	})

	return err
}

func (r *S3ReportRepository) List(ctx context.Context) ([]string, error) {
	output, err := r.client.ListObjectsV2(ctx, &s3.ListObjectsV2Input{
		Bucket: aws.String(r.bucket),
		Prefix: aws.String(r.prefix),
	})
	if err != nil {
		return nil, err
	}

	reports := make([]string, 0, len(output.Contents))

	for _, object := range output.Contents {
		if object.Key == nil {
			continue
		}

		fileName := strings.TrimPrefix(*object.Key, r.prefix)
		if fileName == "" {
			continue
		}

		if strings.HasSuffix(fileName, ".json") {
			reports = append(reports, fileName)
		}
	}

	sort.Sort(sort.Reverse(sort.StringSlice(reports)))

	return reports, nil
}

func (r *S3ReportRepository) Read(ctx context.Context, fileName string) ([]byte, error) {
	if err := ValidateReportFileName(fileName); err != nil {
		return nil, err
	}

	output, err := r.client.GetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String(r.bucket),
		Key:    aws.String(r.objectKey(fileName)),
	})
	if err != nil {
		return nil, err
	}
	defer output.Body.Close()

	return io.ReadAll(output.Body)
}

func (r *S3ReportRepository) objectKey(fileName string) string {
	return path.Join(r.prefix, fileName)
}

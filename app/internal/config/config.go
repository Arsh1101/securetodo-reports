package config

import "os"

type Config struct {
	Addr            string
	DBPath          string
	ReportsDir      string
	ReportStorage   string
	AWSRegion       string
	S3BucketName    string
	S3ReportsPrefix string
}

func Load() Config {
	return Config{
		Addr:            getEnv("APP_ADDR", ":8080"),
		DBPath:          getEnv("DB_PATH", "./data/securetodo.db"),
		ReportsDir:      getEnv("REPORTS_DIR", "./reports"),
		ReportStorage:   getEnv("REPORT_STORAGE", "local"),
		AWSRegion:       getEnv("AWS_REGION", "ca-central-1"),
		S3BucketName:    getEnv("S3_BUCKET_NAME", ""),
		S3ReportsPrefix: getEnv("S3_REPORTS_PREFIX", "reports/"),
	}
}

func getEnv(key string, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}

	return value
}

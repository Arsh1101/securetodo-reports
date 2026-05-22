package config

import "os"

type Config struct {
	Addr       string
	DBPath     string
	ReportsDir string
}

func Load() Config {
	return Config{
		Addr:       getEnv("APP_ADDR", ":8080"),
		DBPath:     getEnv("DB_PATH", "./data/securetodo.db"),
		ReportsDir: getEnv("REPORTS_DIR", "./reports"),
	}
}

func getEnv(key string, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}

	return value
}

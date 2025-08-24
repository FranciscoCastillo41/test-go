package config

import (
	"os"
)

// AppConfig holds basic app settings.
type AppConfig struct {
	AppEnvironment string // "development" | "production" | etc.
	HTTPAddress    string // e.g. ":8080"
	CORSOrigins    string // comma-separated
}

// Load reads environment variables and applies sensible defaults.
func Load() AppConfig {
	return AppConfig{
		AppEnvironment: getEnv("APP_ENVIRONMENT", "development"),
		HTTPAddress:    getEnv("HTTP_ADDRESS", ":8080"),
		CORSOrigins:    getEnv("CORS_ORIGINS", "*"),
	}
}

// getEnv returns the value of key or a default if empty.
func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

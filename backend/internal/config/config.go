package config

import (
	"os"
)

// AppConfig holds basic app settings.
type AppConfig struct {
	AppEnvironment    string // "development" | "production" | etc.
	HTTPAddress       string // e.g. ":8080"
	CORSOrigins       string // comma-separated
	DBURL             string // Database connection URL
	SupabaseURL       string // Supabase project URL
	SupabaseJWTSecret string // Supabase JWT secret
}

// Load reads environment variables and applies sensible defaults.
func Load() AppConfig {
	return AppConfig{
		AppEnvironment:    getEnv("APP_ENVIRONMENT", "development"),
		HTTPAddress:       getEnv("HTTP_ADDRESS", ":8080"),
		CORSOrigins:       getEnv("CORS_ORIGINS", "*"),
		DBURL:             getEnv("DB_URL", ""),
		SupabaseURL:       getEnv("SUPABASE_URL", ""),
		SupabaseJWTSecret: getEnv("SUPABASE_JWT_SECRET", ""),
	}
}

// getEnv returns the value of key or a default if empty.
func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

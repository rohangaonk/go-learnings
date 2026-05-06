package config

import (
	"fmt"
	"os"
)

type Config struct {
	DatabaseURL string
	Port        string
	LogLevel    string
}

// Load reads configuration from environment variables.
// Fails fast if required values are missing — crash at startup, not at 3am.
//
// Node.js bridge:
//   process.env.PORT ?? '8080'  ≈  getEnvOrDefault("PORT", "8080")
//   if (!process.env.DATABASE_URL) throw new Error(...)  ≈  return Config{}, fmt.Errorf(...)
func Load() (Config, error) {
	cfg := Config{
		DatabaseURL: os.Getenv("DATABASE_URL"),
		Port:        getEnvOrDefault("PORT", "8080"),
		LogLevel:    getEnvOrDefault("LOG_LEVEL", "info"),
	}

	if cfg.DatabaseURL == "" {
		// Not required for the in-memory version but keeping the pattern
		// so you're used to it when you wire up a real DB.
		cfg.DatabaseURL = "postgres://localhost:5432/go_learning?sslmode=disable"
	}

	if cfg.Port == "" {
		return Config{}, fmt.Errorf("PORT must not be empty")
	}

	return cfg, nil
}

func getEnvOrDefault(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}

package config

import (
	"os"
)

type Config struct {
	Port         string
	DatabasePath string
	JWTSecret    string
}

func Load() (*Config, error) {
	cfg := &Config{
		Port:         getEnv("APP_PORT", "8080"),
		DatabasePath: getEnv("DB_PATH", "data.db"),
		JWTSecret:    getEnv("JWT_SECRET", "supersecretjwtkey"),
	}

	return cfg, nil
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

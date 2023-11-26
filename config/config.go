package config

import (
	"errors"
	"os"

	_ "github.com/joho/godotenv/autoload"
)

// getEnv retrieves an environment variable or returns the default value
func getEnv(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		// Log or handle the absence of the environment variable
		return defaultValue
	}
	return value
}

// validateConfig validates the AppConfig
func validateConfig(cfg AppConfig) error {
	if cfg.DBDriver == "" {
		return errors.New("DB_DRIVER is not set, using the default value")
	}
	return nil
}

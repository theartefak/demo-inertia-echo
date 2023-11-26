package config

import (
	"errors"
	"os"

	_ "github.com/joho/godotenv/autoload"
)

// AppConfig represents the application configuration
type AppConfig struct {
    DatabaseURL string
}

// LoadConfig loads the application configuration from a .env file
func LoadConfig() (*AppConfig, error) {
    var config AppConfig

    // Retrieve configuration values from environment variables
    config.DatabaseURL = GetEnvOrDefault("DATABASE", "sqlite")
    if config.DatabaseURL == "" {
        return nil, errors.New("DATABASE is not set")
    }

    // Add more configuration parameters as needed

    return &config, nil
}

// Get an environment variable or return the default value
func GetEnvOrDefault(key string, defaultValue string) string {
    value, exists := os.LookupEnv(key)
    if !exists {
        // Log or handle the absence of the environment variable
        return defaultValue
    }

    return value
}

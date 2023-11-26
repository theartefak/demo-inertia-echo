package config

// AppConfig represents the application configuration
type AppConfig struct {
	DBDriver   string
	DBHost     string
	DBPort     string
	DBDatabase string
	DBUsername string
	DBPassword string
}

// LoadConfig loads the application configuration from a .env file
func Load() (*AppConfig, error) {
	config := AppConfig{
		DBDriver   : getEnv("DB_DRIVER", "sqlite"),
		DBHost     : getEnv("DB_HOST", "127.0.0.1"),
		DBPort     : getEnv("DB_PORT", "5432"),
		DBDatabase : getEnv("DB_DATABASE", "artefak"),
		DBUsername : getEnv("DB_USERNAME", "root"),
		DBPassword : getEnv("DB_PASSWORD", ""),
	}

	return &config, validateConfig(config)
}

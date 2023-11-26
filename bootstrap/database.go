package bootstrap

import (
	"errors"
	"log"

	"github.com/glebarez/sqlite"
	"github.com/theartefak/artefak/config"
	"github.com/theartefak/artefak/models"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// InitDB initializes the GORM database
func InitDB() (*gorm.DB, error) {
	var db *gorm.DB
	var err error

	// Load configurations
	config, err := config.Load()
	if err != nil {
		log.Fatalf("Error load config: %s", err)
	}

	// Switch database driver based on the configuration
	switch config.DBDriver {
	case "sqlite":
		db, err = gorm.Open(sqlite.Open("models/database.db"), &gorm.Config{})
	case "pgsql":
		dsn := "host=" + config.DBHost +
			" user=" + config.DBUsername +
			" password=" + config.DBPassword +
			" dbname=" + config.DBDatabase +
			" port=" + config.DBPort +
			" sslmode=disable TimeZone=Asia/Makassar"
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	case "mysql":
		dsn := config.DBUsername + ":" + config.DBPassword +
			"@tcp(" + config.DBHost + ":" + config.DBPort + ")/" +
			config.DBDatabase + "?charset=utf8mb4&parseTime=True&loc=Local"
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	default:
		return nil, errors.New("unsupported database driver")
	}

	if err != nil {
		return nil, err
	}

	// Automigrate the models
	if err := db.AutoMigrate(&models.User{}); err != nil {
		return nil, err
	}

	return db, nil
}

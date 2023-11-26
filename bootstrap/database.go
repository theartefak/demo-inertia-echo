package bootstrap

import (
	"github.com/glebarez/sqlite"
	"github.com/theartefak/artefak/models"
	"gorm.io/gorm"
)

// InitDB initializes the GORM database
func InitDB() (*gorm.DB, error) {
    // Connect to SQLite database
    db, err := gorm.Open(sqlite.Open("models/database.db"), &gorm.Config{})
    if err != nil {
        return nil, err
    }

    // Automigrate the models
    if err := db.AutoMigrate(&models.User{}); err != nil {
        return nil, err
    }

    return db, nil
}

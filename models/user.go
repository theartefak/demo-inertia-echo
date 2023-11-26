package models

import "gorm.io/gorm"

// User represents a user model
type User struct {
	gorm.Model
	Username string
	Email    string
}

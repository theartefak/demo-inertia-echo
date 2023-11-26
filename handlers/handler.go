package handlers

import "gorm.io/gorm"

// Handler is a struct that holds instances
// of handler functions or other dependencies.
type Handler struct{
	DB *gorm.DB
}

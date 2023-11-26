package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/theartefak/artefak/handlers"
)

// RegisterRoutes registers all routes for the application
func HttpRoutes(e *echo.Echo, h *handlers.Handler) {
	// Define your routes here using the Echo instance and your handler functions
	e.GET("/", h.Welcome).Name = "welcome"
}

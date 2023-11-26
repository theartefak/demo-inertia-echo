package bootstrap

import (
	"github.com/labstack/echo/v4"
	"github.com/theartefak/artefak/config"
	"github.com/theartefak/artefak/handlers"
	"github.com/theartefak/artefak/routes"
	"github.com/theartefak/inertia-echo"
	"gorm.io/gorm"
)

// Kernel represents the application kernel
type Kernel struct {
	Echo   *echo.Echo
	Config *config.AppConfig
	DB     *gorm.DB
}

// StartKernel initializes a new application kernel
func StartKernel() (*Kernel, error) {
	// Initialize your Echo instance and other components here
	e := echo.New()
	e.Use(inertia.Middleware(e))
	e.Static("/", "./public")

	// Load configurations
	cfg, err := config.Load()
	if err != nil {
		return nil, err
	}

	// Initialize database
	db, err := InitDB()
	if err != nil {
		// Handle the error, e.g., log and exit the application
		return nil, err
	}

	// Register routes
	routes.HttpRoutes(e, &handlers.Handler{DB: db})

	return &Kernel{
		Echo:   e,
		Config: cfg,
		DB:     db,
	}, nil
}

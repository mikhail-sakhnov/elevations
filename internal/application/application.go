package application

import (
	"github.com/gin-gonic/gin"
	"github.com/soider/elevations/internal"
	"github.com/soider/elevations/internal/mapbox"
)

// Application entry point structure
type Application struct {
	gin *gin.Engine
	cfg Config

	elevationService *internal.ElevationService
}

// Build builds and setup new application instance
func Build(cfg Config) *Application {
	app := &Application{
		cfg: cfg,
		gin: gin.New(),
	}

	app.setupDependencies()
	app.setupRoutes()

	return app
}

func (app *Application) setupDependencies() {
	app.elevationService = internal.NewElevationService(
		mapbox.NewClient(app.cfg.MapboxToken),
		mapbox.NewMapboxElevationDecoder(),
	)
}

// Run runs gin main loop
func (app *Application) Run() error {
	return app.gin.Run(app.cfg.ListenOn)
}

package application

import (
	"github.com/gin-gonic/gin"
	"github.com/soider/elevations/internal"
)

func (app *Application) setupRoutes() {
	// Routes
	app.gin.GET("/health", func(c *gin.Context) {
		app.healthHandler(c)
	})

	app.gin.GET("/elevation", func(c *gin.Context) {
		app.elevationHandler(c)
	})
}

func (app *Application) healthHandler(c *gin.Context) {
	c.JSON(200, map[string]string{
		"status": "ok",
	})
}

func (app *Application) elevationHandler(c *gin.Context) {
	location := internal.Location{}
	if err := c.ShouldBind(&location); err != nil {
		c.JSON(400, map[string]string{
			"error": err.Error(), // XXX: in the production ready system we must not expose all internal errors
		})
		return
	}
	data, err := app.elevationService.GetElevation(c, location)
	if err != nil {
		c.JSON(500, map[string]string{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, map[string]interface{}{
		"result": data,
	})
}

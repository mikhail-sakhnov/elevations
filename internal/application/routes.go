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

	app.gin.POST("/route", func(c *gin.Context) {
		app.routeElevation(c)
	})

}

func (app *Application) healthHandler(c *gin.Context) {
	// XXX: in the production application we must do real health check
	c.JSON(200, map[string]string{
		"status": "ok",
	})
}

func (app *Application) routeElevation(c *gin.Context) {
	var route []internal.Location
	if err := c.ShouldBindJSON(&route); err != nil {
		c.JSON(400, map[string]string{
			"error": err.Error(), // XXX: in the production ready system we must not expose all internal errors
		})
		return
	}
	data, err := app.elevationService.GetElevation(c, route)
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

package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/mersanuzun/go-bff-boilerplate/controllers"
)

func SetMonitoringRoutes(e *echo.Echo) {
	r := e.Group("/_monitoring")

	r.GET("/info", controllers.Monitoring)
	r.GET("/health", controllers.Health)
}
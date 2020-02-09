package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/mersanuzun/go-bff-boilerplate/controllers"
)

func SetRoutes(e *echo.Echo)  {
	e.GET("/", controllers.Home)

	SetMonitoringRoutes(e)
	SetTodoRoutes(e)
}
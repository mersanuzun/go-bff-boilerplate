package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/mersanuzun/go-bff-boilerplate/controllers"
)

func SetTodoRoutes(e *echo.Echo) {
	r := e.Group("/todos")

	r.GET("/:id", controllers.GetTodo)
}
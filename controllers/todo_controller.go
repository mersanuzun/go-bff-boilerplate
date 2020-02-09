package controllers

import (
	"github.com/labstack/echo/v4"
	"github.com/mersanuzun/go-bff-boilerplate/models/todo"
	"github.com/mersanuzun/go-bff-boilerplate/services"
	"net/http"
)

// @Summary Get A Todo
// @Description Get A Todo
// @Accept  json
// @Produce  json
// @Param todo path string true "fileName"
// @Success 200 {object} todo.GetTodoResponse
// @Router /todos/{todo} [get]
// @tags Cargo Desi
func GetTodo(c echo.Context) error {
	getTodoRequest := new(todo.GetTodoRequest)

	c.Bind(getTodoRequest)

	return c.JSON(http.StatusOK, services.GetTodo(getTodoRequest))
}
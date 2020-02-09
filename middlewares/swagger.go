package middlewares

import (
	"github.com/labstack/echo/v4"
	"github.com/mersanuzun/go-bff-boilerplate/config"
	"github.com/mersanuzun/go-bff-boilerplate/docs"
	"github.com/swaggo/echo-swagger"
)

func InitSwagger(e *echo.Echo) {
	docs.SwaggerInfo.Title = config.AppName()
	docs.SwaggerInfo.Description = config.AppName()
	docs.SwaggerInfo.Version = config.Version()
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	e.GET("/swagger/*", echoSwagger.WrapHandler)
}

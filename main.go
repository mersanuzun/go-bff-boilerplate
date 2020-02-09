package main

import (
	"github.com/labstack/echo/v4"
	"github.com/mersanuzun/go-bff-boilerplate/config"
	"github.com/mersanuzun/go-bff-boilerplate/lib/logger"
	"github.com/mersanuzun/go-bff-boilerplate/middlewares"
	"github.com/mersanuzun/go-bff-boilerplate/routes"
)


func main() {
	e := echo.New()

	logger.InitZapLogger()

	middlewares.HandleHttpError(e)
	middlewares.InitSwagger(e)

	routes.SetRoutes(e)

	e.Logger.Fatal(e.Start(":" + config.Port()))
}


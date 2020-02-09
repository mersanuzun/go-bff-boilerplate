package controllers

import (
	"github.com/labstack/echo/v4"
	"github.com/mersanuzun/go-bff-boilerplate/config"
	"github.com/mersanuzun/go-bff-boilerplate/models/shared"
	"net/http"
)

// @Summary Get App Info
// @Description Get App Info
// @Accept  json
// @Produce  json
// @Success 200 {object} shared.AppInformation
// @Router /_monitoring/info [get]
// @tags Monitoring
func Monitoring(c echo.Context) error {
	return c.JSON(http.StatusOK, shared.AppInformation{
		Env:     config.Env(),
		AppName: config.AppName(),
		Version: config.Version(),
	})
}

// @Summary Get App Health
// @Description Get App Health
// @Accept  json
// @Produce  json
// @Success 200 {object} shared.AppStatus
// @Router /_monitoring/health [get]
// @tags Monitoring
func Health(c echo.Context) error {
	return c.JSON(http.StatusOK, shared.AppStatus{
		Status: "Ok",
	})
}
package controllers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func Home(c echo.Context) error {
	return c.Redirect(http.StatusPermanentRedirect, "/swagger/index.html")
}
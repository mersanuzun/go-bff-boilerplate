package middlewares

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/mersanuzun/go-bff-boilerplate/models/shared"
	"go.uber.org/zap"
	"net/http"
)

func HandleHttpError(e *echo.Echo) {
	e.Use(middleware.Recover())

	e.HTTPErrorHandler = func(error error, c echo.Context) {
		var errorResponse shared.ErrorResponse

		if he, ok := error.(*echo.HTTPError); ok {
			if he.Message == "missing or malformed jwt" {
				zap.S().Warn(he)

				c.JSON(http.StatusUnauthorized, map[string]string{
					"message": "missing or malformed jwt",
				})
			} else {
				zap.S().Error(he)

				c.JSON(he.Code, he)
			}
		} else if err := json.Unmarshal([]byte(error.Error()), &errorResponse); err == nil {
			zap.S().Error(errorResponse)

			c.JSONBlob(errorResponse.Status, []byte(errorResponse.Body))
		} else {
			zap.S().Error(error)

			c.NoContent(http.StatusInternalServerError)
		}
	}
}
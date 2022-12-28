package webapi

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Authorize(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authToken := c.Request().Header.Get("Authorization")
		if authToken == "" || authToken != "November 10, 2009" {
			return c.JSON(http.StatusUnauthorized, ErrorResponse{Message: "Unauthorized request"})
		}
		return next(c)
	}
}

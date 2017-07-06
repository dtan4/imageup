package handler

import (
	"net/http"

	"github.com/labstack/echo"
)

// Ping represents the handler for "/ping"
func Ping(c echo.Context) error {
	return c.String(http.StatusOK, "pong")
}

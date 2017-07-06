package handler

import (
	"net/http"

	"github.com/labstack/echo"
)

// PingHandler represents the handler for "/ping"
func PingHandler(c echo.Context) error {
	return c.String(http.StatusOK, "pong")
}

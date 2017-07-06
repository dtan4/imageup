package handler

import (
	"net/http"

	"github.com/labstack/echo"
)

// Root represents the handler for "/"
func Root(c echo.Context) error {
	return c.String(http.StatusOK, "ImageUp")
}

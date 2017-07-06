package handler

import (
	"net/http"

	"github.com/labstack/echo"
)

// RootHandler represents the handler for "/"
func RootHandler(c echo.Context) error {
	return c.String(http.StatusOK, "ImageUp")
}

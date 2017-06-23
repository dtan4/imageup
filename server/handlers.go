package server

import (
	"net/http"

	"github.com/labstack/echo"
)

func rootHandler(c echo.Context) error {
	return c.String(http.StatusOK, "ImageUp")
}

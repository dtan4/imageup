package server

import (
	"github.com/labstack/echo"
)

func drawRoutes(e *echo.Echo) {
	e.GET("/", rootHandler)
}

package server

import (
	"github.com/labstack/echo"
)

func drawRoutes(e *echo.Echo) {
	e.GET("/", rootHandler)
	e.GET("/ping", pingHandler)

	e.POST("/webhooks/quay", webhooksQuayHandler)
}

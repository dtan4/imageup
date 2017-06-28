package server

import (
	"github.com/dtan4/imageup/config"
	"github.com/labstack/echo"
	echoMW "github.com/labstack/echo/middleware"
)

func drawRoutes(e *echo.Echo, conf *config.Config) {
	e.GET("/ping", pingHandler)

	if conf.BasicAuthUsername != "" && conf.BasicAuthPassword != "" {
		basicAuth := echoMW.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
			return username == conf.BasicAuthUsername && password == conf.BasicAuthPassword, nil
		})

		e.GET("/", rootHandler, basicAuth)
		e.POST("/webhooks/quay", webhooksQuayHandler, basicAuth)
	} else {
		e.GET("/", rootHandler)
		e.POST("/webhooks/quay", webhooksQuayHandler)
	}
}

package server

import (
	"github.com/dtan4/imageup/config"
	"github.com/dtan4/imageup/server/handler"
	"github.com/labstack/echo"
	echoMW "github.com/labstack/echo/middleware"
)

func drawRoutes(e *echo.Echo, conf *config.Config) {
	e.GET("/ping", handler.PingHandler)

	if conf.BasicAuthUsername != "" && conf.BasicAuthPassword != "" {
		basicAuth := echoMW.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
			return username == conf.BasicAuthUsername && password == conf.BasicAuthPassword, nil
		})

		e.GET("/", handler.RootHandler, basicAuth)
		e.POST("/webhooks/quay", handler.WebhooksQuayHandler, basicAuth)
	} else {
		e.GET("/", handler.RootHandler)
		e.POST("/webhooks/quay", handler.WebhooksQuayHandler)
	}
}

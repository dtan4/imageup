package server

import (
	"github.com/dtan4/imageup/config"
	"github.com/dtan4/imageup/server/handler"
	"github.com/labstack/echo"
	echoMW "github.com/labstack/echo/middleware"
)

func drawRoutes(e *echo.Echo, conf *config.Config) {
	e.GET("/ping", handler.Ping)

	if conf.BasicAuthUsername != "" && conf.BasicAuthPassword != "" {
		basicAuth := echoMW.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
			return username == conf.BasicAuthUsername && password == conf.BasicAuthPassword, nil
		})

		e.GET("/", handler.Root, basicAuth)
		e.POST("/webhooks/quay", handler.WebhooksQuay, basicAuth)
	} else {
		e.GET("/", handler.Root)
		e.POST("/webhooks/quay", handler.WebhooksQuay)
	}
}

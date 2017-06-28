package server

import (
	"fmt"

	"github.com/dtan4/imageup/config"
	"github.com/dtan4/imageup/docker"
	"github.com/dtan4/imageup/server/middleware"
	"github.com/labstack/echo"
	echoMW "github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
)

// Run starts ImageUp server
func Run(conf *config.Config) {
	e := echo.New()

	e.Use(echoMW.Logger())
	e.Use(echoMW.Recover())

	dockerClient, err := docker.NewClient()
	if err != nil {
		e.Logger.Fatal(err)
	}
	e.Use(middleware.SetDockerClient(dockerClient))

	drawRoutes(e, conf)

	addr := fmt.Sprintf(":%d", conf.Port)

	e.Logger.SetLevel(log.ERROR)
	e.Logger.Fatal(e.Start(addr))
}

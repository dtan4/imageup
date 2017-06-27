package server

import (
	"fmt"

	"github.com/dtan4/imageup/docker"
	"github.com/dtan4/imageup/server/middleware"
	"github.com/labstack/echo"
	echoMW "github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
)

// Run starts ImageUp server
func Run(port int) {
	e := echo.New()

	e.Use(echoMW.Logger())
	e.Use(echoMW.Recover())

	dockerClient, err := docker.NewClient()
	if err != nil {
		e.Logger.Fatal(err)
	}
	e.Use(middleware.SetDockerClient(dockerClient))

	drawRoutes(e)

	addr := fmt.Sprintf(":%d", port)

	e.Logger.SetLevel(log.ERROR)
	e.Logger.Fatal(e.Start(addr))
}

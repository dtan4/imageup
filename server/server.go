package server

import (
	"fmt"

	"github.com/dtan4/imageup/docker"
	imageupMW "github.com/dtan4/imageup/server/middleware"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// Run starts ImageUp server
func Run(port int) {
	e := echo.New()

	dockerClient, err := docker.NewClient()
	if err != nil {
		e.Logger.Fatal(err)
	}

	e.Use(imageupMW.SetDockerClient(dockerClient))

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	drawRoutes(e)

	addr := fmt.Sprintf(":%d", port)

	e.Logger.Fatal(e.Start(addr))
}

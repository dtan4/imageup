package server

import (
	"fmt"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// Run starts ImageUp server
func Run(port int) {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	drawRoutes(e)

	addr := fmt.Sprintf(":%d", port)

	e.Logger.Fatal(e.Start(addr))
}

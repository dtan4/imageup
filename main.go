package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

const (
	defaultPort = 8000
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "imageup")
	})

	var port int

	if os.Getenv("PORT") == "" {
		port = defaultPort
	} else {
		p, err := strconv.Atoi(os.Getenv("PORT"))
		if err != nil {
			e.Logger.Fatal(err)
		}

		port = p
	}

	addr := fmt.Sprintf(":%d", port)

	e.Logger.Fatal(e.Start(addr))
}

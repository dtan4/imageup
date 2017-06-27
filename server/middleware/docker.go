package middleware

import (
	"github.com/dtan4/imageup/docker"
	"github.com/labstack/echo"
	"github.com/pkg/errors"
)

const (
	dockerContextKey = "DOCKER"
)

// SetDockerClients sets the given Docker client to request context
func SetDockerClient(client docker.Client) echo.MiddlewareFunc {
	return func(h echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set(dockerContextKey, client)

			return h(c)
		}
	}
}

// GetDockerClient fetches Docker client from request context
func GetDockerClient(c echo.Context) (docker.Client, error) {
	v := c.Get(dockerContextKey)

	if v == nil {
		return nil, errors.New("Docker client is not found in context")
	}

	client, ok := v.(docker.Client)
	if !ok {
		return nil, errors.Errorf("invalid object is stored in context %q", dockerContextKey)
	}

	return client, nil
}

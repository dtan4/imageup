package middleware

import (
	"github.com/labstack/echo"
	"github.com/pkg/errors"
)

const (
	imageWhitelistContextKey = "IMAGE_WHITELIST"
)

// SetImageWhitelist sets the given Docker client to request context
func SetImageWhitelist(imageWhitelist []string) echo.MiddlewareFunc {
	return func(h echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set(imageWhitelistContextKey, imageWhitelist)

			return h(c)
		}
	}
}

// GetImageWhitelist fetches Docker client from request context
func GetImageWhitelist(c echo.Context) ([]string, error) {
	v := c.Get(imageWhitelistContextKey)

	if v == nil {
		return nil, nil
	}

	imageWhitelist, ok := v.([]string)
	if !ok {
		return nil, errors.Errorf("invalid object is stored in context %q", imageWhitelistContextKey)
	}

	return imageWhitelist, nil
}

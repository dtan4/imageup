package config

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
)

const (
	prefix = "imageup"
)

// Config represents configurations of ImageUp
type Config struct {
	BasicAuthPassword string   `envconfig:"basic_auth_password"`
	BasicAuthUsername string   `envconfig:"basic_auth_username"`
	ImageWhiteList    []string `envconfig:"image_whitelist"`
	Port              int      `envconfig:"port" default:"8000"`
}

// Load loads configurations from environment variables
func Load() (*Config, error) {
	var config Config

	if err := envconfig.Process(prefix, &config); err != nil {
		return nil, errors.Wrap(err, "failed to load configurations from env")
	}

	return &config, nil
}

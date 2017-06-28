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
	Port int `envconfig:"port"`
}

// Load loads configurations from environment variables
func Load() (*Config, error) {
	var config Config

	if err := envconfig.Process(prefix, &config); err != nil {
		return nil, errors.Wrap(err, "failed to load configurations from env")
	}

	return &config, nil
}

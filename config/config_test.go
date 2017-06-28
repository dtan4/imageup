package config

import (
	"os"
	"testing"
)

func TestLoad(t *testing.T) {
	testcases := []struct {
		envs      map[string]string
		expectErr bool
	}{
		{
			envs:      map[string]string{},
			expectErr: false,
		},
		{
			envs: map[string]string{
				"IMAGEUP_PORT": "12345",
			},
			expectErr: false,
		},
		{
			envs: map[string]string{
				"IMAGEUP_PORT": "foobar",
			},
			expectErr: true,
		},
	}

	for _, tc := range testcases {
		config, err := Load()
	}
}

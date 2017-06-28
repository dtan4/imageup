package config

import (
	"os"
	"reflect"
	"testing"
)

func TestLoad(t *testing.T) {
	testcases := []struct {
		envs           map[string]string
		expectedConfig *Config
		expectErr      bool
	}{
		{
			envs: map[string]string{},
			expectedConfig: &Config{
				BasicAuthPassword: "",
				BasicAuthUsername: "",
				Port:              8000,
			},
			expectErr: false,
		},
		{
			envs: map[string]string{
				"IMAGEUP_BASIC_AUTH_PASSWORD": "password",
				"IMAGEUP_BASIC_AUTH_USERNAME": "username",
				"IMAGEUP_PORT":                "12345",
			},
			expectedConfig: &Config{
				BasicAuthPassword: "password",
				BasicAuthUsername: "username",
				Port:              12345,
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
		os.Clearenv()

		for k, v := range tc.envs {
			os.Setenv(k, v)
		}

		config, err := Load()
		if tc.expectErr {
			if err == nil {
				t.Errorf("error should be raised")
			}

			continue
		}

		if err != nil {
			t.Errorf("error should not be raised: %s", err)
		}

		if !reflect.DeepEqual(config, tc.expectedConfig) {
			t.Errorf("config does not match, expected: %#v, got: %#v", tc.expectedConfig, config)
		}
	}
}

package config

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	InputStreamPath  string `envconfig:"INPUT_STREAM_PATH"`
	OutputStreamPath string `envconfig:"OUTPUT_STREAM_PATH"`
	ServerOnePort    string `envconfig:"SERVER_ONE_PORT"`
	ServerTwoPort    string `envconfig:"SERVER_TWO_PORT"`
}

// New generate new Config struct from environment variables
func New() (*Config, error) {
	var c Config
	if err := envconfig.Process("", &c); err != nil {
		return nil, err
	}

	return &c, nil
}

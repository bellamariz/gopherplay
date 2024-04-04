package config

import (
	"time"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	InputStreamPath       string        `envconfig:"INPUT_STREAM_PATH"`
	DiscoveryRunFrequency time.Duration `envconfig:"DISCOVERY_RUN_FREQUENCY"`
	LiveSignalName        string        `envconfig:"LIVE_SIGNAL_NAME"`
	LocalHost             string        `envconfig:"LOCAL_HOST"`
	MaxAgePlaylist        time.Duration `envconfig:"MAX_AGE_PLAYLIST"`
	OutputStreamPath      string        `envconfig:"OUTPUT_STREAM_PATH"`
	ReporterPort          string        `envconfig:"REPORTER_PORT"`
	ServerOnePort         string        `envconfig:"SERVER_ONE_PORT"`
	ServerTwoPort         string        `envconfig:"SERVER_TWO_PORT"`
}

// New generate new Config struct from environment variables
func New() (*Config, error) {
	var c Config
	if err := envconfig.Process("", &c); err != nil {
		return nil, err
	}

	return &c, nil
}

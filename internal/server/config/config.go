package config

import (
	"fmt"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Auth   map[string]string `yaml:"auth" env-description:"username password set" env-default:""`
	Proto  string            `yaml:"proto" env-description:"listen proto tcp/udp" env-default:"tcp"`
	Listen string            `yaml:"listen" env-description:"listen address" env-default:"0.0.0.0:1080"`
}

func New(filename string) (*Config, error) {
	var _config = Config{
		Auth: make(map[string]string),
	}
	if err := cleanenv.ReadConfig(filename, &_config); err != nil {
		if !os.IsNotExist(err) {
			return &_config, fmt.Errorf("failed to parse config: %w", err)
		}
	}
	return &_config, nil
}

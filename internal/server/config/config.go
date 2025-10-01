package config

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/ilyakaznacheev/cleanenv"
	"gopkg.in/yaml.v3"
)

var (
	_true = true
)

var _config Config = Config{
	Auth: map[string]string{
		"username": "password",
	},
	Socks: Socks{
		Enable: &_true,
		Proto:  "tcp",
		Listen: "0.0.0.0:1080",
	},
	Https: Https{
		Enable: &_true,
		Listen: "0.0.0.0:8443",
	},
}

type Config struct {
	Auth  map[string]string `yaml:"auth" env-default:""`
	Socks Socks             `yaml:"socks"`
	Https Https             `yaml:"https"`
}

type Socks struct {
	Enable *bool  `yaml:"enable" env-default:"true"`
	Proto  string `yaml:"proto" env-default:"tcp"`
	Listen string `yaml:"listen" env-default:"0.0.0.0:1080"`
}

type Https struct {
	Enable *bool  `yaml:"enable" env-default:"true"`
	Listen string `yaml:"listen" env-default:"0.0.0.0:8443"`
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

func Default(path string) error {
	return save(_config, path)
}

func save(config any, path string) error {
	data, err := yaml.Marshal(&config)
	if err != nil {
		return fmt.Errorf("failed to marshall config: %w", err)
	}
	if err := os.MkdirAll(filepath.Dir(path), 0o777); err != nil {
		return fmt.Errorf("failed to save default config: %w", err)
	}
	if err := writeFile(path, data); err != nil {
		return fmt.Errorf("failed to save default config: %w", err)
	}
	return nil
}

func writeFile(
	path string,
	data []byte,
	perm ...int,
) error {
	if len(perm) == 0 {
		perm = append(perm, 0644)
	}
	if err := os.WriteFile(path, data, fs.FileMode(perm[0])); err != nil {
		return fmt.Errorf("failed to save file: %w", err)
	}
	return nil
}

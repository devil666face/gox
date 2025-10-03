package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"gox/pkg/fs"

	"github.com/ilyakaznacheev/cleanenv"
	"gopkg.in/yaml.v3"
)

var (
	_true = true
)

var config Config = Config{
	Auth: map[string]string{
		"username": "password",
	},
	Socks: Socks{
		Enable: &_true,
		Listen: "0.0.0.0:31080",
	},
	Https: Https{
		Enable: &_true,
		Listen: "0.0.0.0:38443",
	},
}

type Config struct {
	Auth  map[string]string `yaml:"auth" env-default:""`
	Socks Socks             `yaml:"socks"`
	Https Https             `yaml:"https"`
}

type Socks struct {
	Enable *bool  `yaml:"enable" env-default:"true"`
	Listen string `yaml:"listen" env-default:"0.0.0.0:31080"`
}

type Https struct {
	Enable *bool  `yaml:"enable" env-default:"true"`
	Listen string `yaml:"listen" env-default:"0.0.0.0:38443"`
}

func New(filename string, username, password string) (*Config, error) {
	if err := cleanenv.ReadConfig(filename, &config); err != nil {
		log.Println(err, "use default config")
	}
	if username != "" && password != "" {
		config.Auth = map[string]string{
			username: password,
		}
	}
	log.Println("auth", config.Auth)
	return &config, nil
}

func Default(path string) error {
	return save(config, path)
}

func save(config any, path string) error {
	data, err := yaml.Marshal(&config)
	if err != nil {
		return fmt.Errorf("failed to marshall config: %w", err)
	}
	if err := os.MkdirAll(filepath.Dir(path), 0o777); err != nil {
		return fmt.Errorf("failed to save default config: %w", err)
	}
	if err := fs.WriteFile(path, data); err != nil {
		return fmt.Errorf("failed to save default config: %w", err)
	}
	return nil
}

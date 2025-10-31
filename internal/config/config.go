package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type App struct {
	Name     string
	Env      string
	LogLevel string
}

type SMTP struct {
	Host     string
	Port     int
	Username string
	Password string
}

type Config struct {
	AppConfig  App
	SMTPConfig SMTP
}

func Load(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}

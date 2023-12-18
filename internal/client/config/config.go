// Package config отвечает за конфигурацию клиента.
package config

import (
	"encoding/json"
	"os"

	"github.com/caarlos0/env"
)

// Config - структура конфига, содержит в себе настройки приложения.
type Config struct {
	// Address - адрес, на котором запускается клиент.
	Address string `json:"address" env:"ADDRESS"`
	// LogLevel – уровень логгирования.
	LogLevel string `json:"log_level" env:"LOG_LEVEL"`
}

// NewConfig – конструктор Config.
func NewConfig(configPath string) (*Config, error) {
	configFile, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}

	config := &Config{}
	err = json.NewDecoder(configFile).Decode(config)
	if err != nil {
		return nil, err
	}

	err = env.Parse(config)
	return config, err
}

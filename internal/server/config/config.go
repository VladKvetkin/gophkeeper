package config

import (
	"encoding/json"
	"os"

	"github.com/caarlos0/env"
)

// Config – объект конфигурации сервера.
type Config struct {
	// DatabaseDSN – dsn для подключения к БД.
	DatabaseDSN string `json:"database_dsn" env:"DATABASE_DSN"`
	// Address – адрес сервера.
	Address string `json:"address" env:"ADDRESS"`
	// LogLevel – уровень логгирования.
	LogLevel string `json:"log_level" env:"LOG_LEVEL"`
	// SecretKey – ключ шифрования.
	SecretKey string `json:"secret_key"`
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

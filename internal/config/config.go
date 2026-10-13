package config

import (
	"path/filepath"

	"github.com/spf13/viper"
)

type Config struct {
	DatabaseDSN string `mapstructure:"DATABASE_DSN"`
	LogLevel    string `mapstructure:"LOG_LEVEL"`
	AppPort     int    `mapstructure:"APP_PORT"`
}

func NewConfig(path string) (*Config, error) {
	v := viper.New()

	v.SetConfigFile(filepath.Join(path, ".env"))
	v.SetConfigType("env")
	v.AutomaticEnv()

	v.SetDefault("LOG_LEVEL", "info")
	v.SetDefault("APP_PORT", 8080)
	v.SetDefault("DATABASE_DSN", "localhost")

	err := v.ReadInConfig()
	if err != nil {
		return nil, err
	}

	var cfg Config
	err = v.Unmarshal(&cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}

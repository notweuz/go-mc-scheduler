package config

import (
	"os"

	"github.com/goccy/go-yaml"
	"github.com/rs/zerolog/log"
)

type Config struct {
	GlobalConfig CommonConfig `yaml:"go_mc_restart"`
}

type CommonConfig struct {
	Scheduler Scheduler  `yaml:"scheduler"`
	Rcon      RconConfig `yaml:"rcon"`
}

type RconConfig struct {
	Address  string `yaml:"address"`
	Password string `yaml:"password"`
}

type Scheduler struct {
	Cron     string `yaml:"cron"`
	Timezone string `yaml:"timezone"`
}

var Instance *Config

func LoadConfig(path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to read config yml")
	}

	Instance = &Config{}

	log.Info().Msg("Loaded config yml")
	err = yaml.Unmarshal(data, Instance)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to parse config yml")
	}
	log.Info().Msg("Parsed application configuration")

	return nil
}

func GetConfig() *CommonConfig {
	return &Instance.GlobalConfig
}

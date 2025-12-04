package main

import (
	"gorestart-minecraft/internal/config"
	"gorestart-minecraft/internal/rcon"
	"gorestart-minecraft/internal/scheduler"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"
)

var Version = "2025.12.1"

func main() {
	setupLogger()
	log.Info().Str("version", Version).Msg("Starting go-mc-scheduler service")

	err := config.LoadConfig("configs/application.yml")
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to load config")
	}

	rcon.SetupRCONClient(&config.GetConfig().Rcon)

	sched := scheduler.NewScheduler(&config.GetConfig().Scheduler)
	if err := sched.Start(); err != nil {
		log.Fatal().Err(err).Msg("Failed to start scheduler")
	}
	defer sched.Stop()

	select {}
}

func setupLogger() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr}).With().Caller().Logger()
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
}

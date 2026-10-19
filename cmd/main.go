package main

import (
	"context"
	"go-rcon-panel/internal/config"
	"go-rcon-panel/internal/logger"
	"go-rcon-panel/internal/server"
	"os"
	"os/signal"
	"syscall"

	"github.com/rs/zerolog/log"
)

func main() {
	logger.Setup()
	log.Info().Msg("Starting server")

	cfg, err := config.NewConfig(".")
	if err != nil {
		log.Fatal().Err(err).Msg("failed to load config")
	}
	logger.UpdateLogLevel(cfg.LogLevel)

	app, err := server.NewApp(cfg)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to init app")
	}
	app.Start()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	log.Debug().Msg("Waiting for shutdown signal")
	<-ctx.Done()

	app.Shutdown()
}

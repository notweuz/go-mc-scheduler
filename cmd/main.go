package main

import (
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
		log.Panic().Err(err).Msg("Error loading config")
	}
	logger.UpdateLogLevel(cfg.LogLevel)

	app, err := server.NewApp(cfg)
	if err != nil {
		log.Panic().Err(err).Msg("Error loading config")
	}
	app.Start()

	log.Debug().Msg("Initialized graceful shutdown")
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	<-quit
	app.Shutdown()
}

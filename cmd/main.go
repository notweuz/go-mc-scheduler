package main

import (
	"go-mc-scheduler/internal/config"
	"go-mc-scheduler/internal/logger"
	"go-mc-scheduler/internal/server"
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

	app := server.NewApp(cfg)
	app.Start()

	log.Debug().Msg("Initialized graceful shutdown")
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	<-quit
	app.Shutdown()
}
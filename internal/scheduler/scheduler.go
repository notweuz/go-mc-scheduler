package scheduler

import (
	"gorestart-minecraft/internal/config"
	"gorestart-minecraft/internal/rcon"
	"time"

	"github.com/robfig/cron/v3"
	"github.com/rs/zerolog/log"
)

type Scheduler struct {
	cron   *cron.Cron
	config *config.Scheduler
}

func NewScheduler(cfg *config.Scheduler) *Scheduler {
	location, err := time.LoadLocation(cfg.Timezone)
	if err != nil {
		log.Warn().Err(err).Msg("Failed to load timezone, using UTC")
		location = time.UTC
	}

	return &Scheduler{
		cron:   cron.New(cron.WithLocation(location)),
		config: cfg,
	}
}

func (s *Scheduler) Start() error {
	_, err := s.cron.AddFunc(s.config.Cron, s.restartServer)
	if err != nil {
		return err
	}

	s.cron.Start()
	log.Info().Str("cron", s.config.Cron).Msg("Scheduler started")
	return nil
}

func (s *Scheduler) Stop() {
	s.cron.Stop()
	log.Info().Msg("Scheduler stopped")
}

func (s *Scheduler) restartServer() {
	log.Info().Msg("Scheduled task triggered: restarting server")

	log.Info().Msg("Connecting to server using RCON")
	if err := rcon.Client.Connect(); err != nil {
		log.Error().Err(err).Msg("Failed to connect to RCON")
		return
	}

	for _, w := range RestartSequence {
		for _, err := rcon.Client.Execute(w.message); err != nil; {
			log.Error().Err(err).Str("message", w.message).Msg("Failed run restart sequence entry")
		}
		time.Sleep(w.delay)
	}

	log.Info().Msg("Saving server")
	for _, err := rcon.Client.Execute("save-all"); err != nil; {
		log.Error()
	}

	log.Info().Msg("Stopping server")
	for _, err := rcon.Client.Execute("stop"); err != nil; {
		log.Error().Err(err).Msg("Failed to stop server")
	}

	log.Info().Msg("Restarting server finished!")

	log.Info().Msg("Closing rcon connection")
	if err := rcon.Client.Close(); err != nil {
		log.Error().Err(err).Msg("Failed to close rcon connection")
	}
}

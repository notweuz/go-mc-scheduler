package scheduler

import (
	"go-mc-scheduler/internal/config"
	"go-mc-scheduler/internal/rcon"
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
	for _, job := range s.config.Jobs {
		_, err := s.cron.AddFunc(job.Cron, s.runJob(job))
		if err != nil {
			return err
		}
		log.Info().Str("job", job.Name).Str("cron", job.Cron).Msg("Scheduled job added")
	}

	s.cron.Start()
	log.Info().Msg("Scheduler started with all jobs")
	return nil
}

func (s *Scheduler) Stop() {
	s.cron.Stop()
	log.Info().Msg("Scheduler stopped")
}

func (s *Scheduler) runJob(job config.Job) func() {
	return func() {
		log.Info().Str("job", job.Name).Msg("Scheduled job triggered")

		log.Info().Str("job", job.Name).Msg("Connecting to server using RCON")
		conn, err := rcon.Connect()
		if err != nil {
			log.Error().Str("job", job.Name).Err(err).Msg("Failed to connect to RCON")
			return
		}
		defer func() {
			log.Info().Str("job", job.Name).Msg("Scheduled job completed")
			if err := conn.Close(); err != nil {
				log.Error().Err(err).Str("job", job.Name).Msg("Failed to close RCON connection")
			}
		}()

		for _, w := range job.Steps {
			if w.Execute != nil {
				for _, err := conn.Execute(*w.Execute); err != nil; {
					log.Error().Err(err).Msg("Failed to execute command")
				}
			}
			if w.Wait != nil {
				duration, err := time.ParseDuration(*w.Wait)
				if err != nil {
					log.Error().Err(err).Msg("Failed to parse wait duration")
					continue
				}
				log.Info().Str("job", job.Name).Str("duration", duration.String()).Msg("Waiting")
				time.Sleep(duration)
			}
		}
	}
}

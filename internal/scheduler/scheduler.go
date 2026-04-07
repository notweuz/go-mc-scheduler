package scheduler

import (
	"context"
	"go-mc-scheduler/internal/config"
	"go-mc-scheduler/internal/rcon"
	"sync"
	"time"

	"github.com/robfig/cron/v3"
	"github.com/rs/zerolog/log"
)

type Scheduler struct {
	cron   *cron.Cron
	config *config.Scheduler
	wg     sync.WaitGroup
	jobs   []config.Job
	done   chan struct{}
}

func NewScheduler(cfg *config.Scheduler) *Scheduler {
	location, err := time.LoadLocation(cfg.Timezone)
	if err != nil {
		log.Warn().Err(err).Msg("Failed to load timezone, using UTC")
		location = time.UTC
	}

	return &Scheduler{
		cron:   cron.New(cron.WithLocation(location)),
		done:   make(chan struct{}),
		config: cfg,
	}
}

func (s *Scheduler) Start() error {
	log.Info().Str("timezone", s.config.Timezone).Msg("Starting scheduler with timezone")

	for _, job := range s.config.Jobs {
		_, err := s.cron.AddFunc(job.Cron, s.runJob(job))
		if err != nil {
			return err
		}
		log.Info().Str("job", job.Name).Str("cron", job.Cron).Msg("Scheduled job added")
	}

	s.cron.Start()
	log.Info().Msg("Imported all scheduled jobs and started scheduler")

	entries := s.cron.Entries()
	for _, entry := range entries {
		nextRun := entry.Next.Format(time.RFC3339)
		log.Info().Str("next_run", nextRun).Msg("Next scheduled run")
	}

	return nil
}

func (s *Scheduler) Stop() {
	s.cron.Stop()

	shutdownContext, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	go func() {
		s.wg.Wait()
		s.done <- struct{}{}
	}()

	select {
	case <-s.done:
		log.Info().Msg("Scheduler stopped gracefully")
	case <-shutdownContext.Done():
		log.Info().Msg("Scheduler shutdown timeout")
	}
}

func (s *Scheduler) runJob(job config.Job) func() {
	return func() {
		s.wg.Add(1)
		defer s.wg.Done()

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
			select {
			case <-s.done:
				log.Info().Str("job", job.Name).Msg("Aborting job, scheduler is shutting down")
				return
			default:
				if w.Execute != nil {
					_, err := conn.Execute(*w.Execute)
					if err != nil {
						log.Error().Err(err).Str("command", *w.Execute).Msg("Failed to execute command, skipping to next step")
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
}

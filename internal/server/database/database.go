package database

import (
	"rcon-scheduler-panel/internal/config"

	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Setup(cfg *config.Config) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(cfg.DatabaseDSN), &gorm.Config{
		Logger:         logger.Default.LogMode(logger.Silent),
		TranslateError: true,
	})
	if err != nil {
		return nil, err
	}
	log.Info().Msg("Connected to database")

	err = Migrate(db)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to migrate database")
		return nil, err
	}
	log.Info().Msg("Migrated database successfully")
	return db, nil
}

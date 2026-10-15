package database

import (
	"go-mc-scheduler/migrations"

	"github.com/pressly/goose/v3"
	"gorm.io/gorm"
)

func Migrate(gormDB *gorm.DB) error {
	db, err := gormDB.DB()
	if err != nil {
		return err
	}

	goose.SetBaseFS(migrations.FS)
	goose.SetDialect("pgx")

	return goose.Up(db, ".")
}

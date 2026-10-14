package server

import (
	"context"
	"fmt"
	"go-mc-scheduler/internal/config"
	"go-mc-scheduler/internal/middleware"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"

	"gorm.io/gorm"
)

type App interface {
	Start()
	Shutdown()
}

type app struct {
	DB     *gorm.DB
	Gin    *gin.Engine
	Config *config.Config
	Srv    *http.Server
}

func NewApp(cfg *config.Config) App {
	db, err := SetupDatabase(cfg)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to start database!")
	}

	r := gin.New()
	r.Use(middleware.ErrorHandler())
	r.Use(gin.Recovery())
	r.Use(cors.Default())
	r.Use(middleware.Logger())

	return &app{
		DB:     db,
		Config: cfg,
		Gin:    r,
	}
}

func (a *app) Start() {
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", a.Config.AppPort),
		Handler: a.Gin.Handler(),
	}
	go func() {
		log.Info().Msg("Starting application")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal().Msgf("listen: %s\n", err)
		}
	}()
	a.Srv = srv
}

func (a *app) Shutdown() {
	log.Info().Msg("Shutting down application")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	if err := a.Srv.Shutdown(ctx); err != nil {
		log.Panic().Err(err).Msg("Error shutting down")
	}
	log.Info().Msg("Server gracefully stopped")
}

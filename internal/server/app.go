package server

import (
	"context"
	"fmt"
	"go-rcon-panel/internal/config"
	"go-rcon-panel/internal/domain/auth"
	"go-rcon-panel/internal/middleware"
	"go-rcon-panel/internal/server/database"
	"go-rcon-panel/internal/validation"
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
	db, err := database.Setup(cfg)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to start database!")
	}

	validation.Register()

	r := gin.New()
	r.Use(middleware.Logger())
	r.Use(middleware.ErrorHandler())
	r.Use(gin.CustomRecovery(middleware.Recovery()))
	r.Use(cors.Default())

	r.HandleMethodNotAllowed = true
	r.NoRoute(middleware.NoRoute())
	r.NoMethod(middleware.NoMethod())

	api := r.Group("/api")

	authSVC := auth.NewService()
	MakeRoutes(api, auth.NewHandler(authSVC))

	return &app{
		DB:     db,
		Config: cfg,
		Gin:    r,
	}
}

func (a *app) Start() {
	srv := &http.Server{
		Addr:              fmt.Sprintf(":%d", a.Config.AppPort),
		Handler:           a.Gin.Handler(),
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      30 * time.Second,
		IdleTimeout:       60 * time.Second,
	}
	go func() {
		log.Info().Msg("Starting application")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Error().Err(err).Msg("listen failed")
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

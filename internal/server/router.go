package server

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type RouterRegistrar interface {
	Register(r *gin.RouterGroup)
}

func MakeRoutes(r *gin.RouterGroup, handlers ...RouterRegistrar) {
	for _, handler := range handlers {
		handler.Register(r)
	}
	log.Info().Msg("handlers registered")
}
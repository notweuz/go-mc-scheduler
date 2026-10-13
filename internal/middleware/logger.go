package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Debug().Str("ip", c.ClientIP()).Str("method", c.Request.Method).Str("path", c.Request.URL.Path).Msgf("Request received")
		start := time.Now()
		c.Next()

		latency := time.Since(start)
		log.Info().Msgf("HTTP %s %s completed with status %d in %v", c.Request.Method, c.Request.URL.Path, c.Request.Response.StatusCode, latency)
	}
}
package middleware

import (
	"go-mc-scheduler/internal/errs"
	"runtime/debug"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func Recovery() gin.RecoveryFunc {
	return func(c *gin.Context, recovered any) {
		log.Error().
			Interface("panic", recovered).
			Str("stack", string(debug.Stack())).
			Str("method", c.Request.Method).
			Str("path", c.Request.URL.Path).
			Msg("panic recovered")

		c.Abort()
		_ = c.Error(errs.ErrInternalServer)
	}
}

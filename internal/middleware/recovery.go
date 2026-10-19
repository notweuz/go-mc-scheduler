package middleware

import (
	"fmt"
	"go-rcon-panel/internal/errs"
	"runtime/debug"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func Recovery() gin.RecoveryFunc {
	return func(c *gin.Context, recovered any) {
		cause := fmt.Errorf("panic: %v", recovered)
		log.Error().Err(cause).Bytes("stack", debug.Stack()).Msg("panic recovered")

		c.Abort()
		_ = c.Error(errs.ErrInternalServer.WithErr(cause))
	}
}

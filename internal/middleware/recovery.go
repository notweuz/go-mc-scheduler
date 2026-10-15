package middleware

import (
	"go-mc-scheduler/internal/errs"

	"github.com/gin-gonic/gin"
)

func Recovery() gin.RecoveryFunc {
	return func(c *gin.Context, recovered any) {
		c.Abort()
		_ = c.Error(errs.ErrInternalServer)
	}
}

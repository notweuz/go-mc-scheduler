package middleware

import (
	"go-mc-scheduler/internal/errs"

	"github.com/gin-gonic/gin"
)

func NoRouteHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		_ = c.Error(errs.ErrNotFound)
	}
}
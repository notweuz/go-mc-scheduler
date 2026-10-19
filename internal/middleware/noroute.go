package middleware

import (
	"go-rcon-panel/internal/errs"

	"github.com/gin-gonic/gin"
)

func NoRoute() gin.HandlerFunc {
	return func(c *gin.Context) {
		_ = c.Error(errs.ErrNotFound)
	}
}

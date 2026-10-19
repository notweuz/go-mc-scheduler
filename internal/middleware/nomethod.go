package middleware

import (
	"go-rcon-panel/internal/errs"

	"github.com/gin-gonic/gin"
)

func NoMethod() gin.HandlerFunc {
	return func(c *gin.Context) {
		_ = c.Error(errs.ErrMethodNotAllowed)
	}
}

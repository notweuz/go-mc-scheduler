package middleware

import (
	"go-mc-scheduler/internal/errs"

	"github.com/gin-gonic/gin"
)

func NoMethod() gin.HandlerFunc {
	return func(c *gin.Context) {
		_ = c.Error(errs.ErrMethodNotAllowed)
	}
}
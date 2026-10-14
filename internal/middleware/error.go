package middleware

import (
	"errors"
	"go-mc-scheduler/internal/dto/response"
	"go-mc-scheduler/internal/errs"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) == 0 || c.Writer.Written() {
			return
		}

		var appErr *errs.AppError
		err := c.Errors.Last().Err
		if !errors.As(err, &appErr) {
			log.Error().
				Err(err).
				Str("method", c.Request.Method).
				Str("path", c.FullPath()).
				Msg("unhandled error")
			appErr = errs.ErrInternalServer
		} else if appErr.StatusCode >= 500 {
			log.Error().
				Err(appErr).
				Str("code", appErr.Status).
				Str("method", c.Request.Method).
				Str("path", c.FullPath()).
				Msg("internal app error")
		}

		c.AbortWithStatusJSON(appErr.StatusCode, response.Response{
			Success: false,
			Error: &response.ErrorInfo{
				Code: appErr.Status,
				Message: appErr.Message,
			},
		})
	}
}

package middleware

import (
	"errors"
	"go-mc-scheduler/internal/dto/response"
	"go-mc-scheduler/internal/errs"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) == 0 || c.Writer.Written() {
			return
		}

		err := c.Errors.Last().Err
		var appErr *errs.AppError
        if !errors.As(err, &appErr) {
            appErr = errs.ErrInternalServer
        }

        if appErr.StatusCode >= http.StatusInternalServerError {
            log.Error().
                Err(err).
                Str("method", c.Request.Method).
                Str("path", c.Request.URL.Path).
                Msg("request failed")
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

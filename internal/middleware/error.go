package middleware

import (
	"errors"
	"go-rcon-panel/internal/dto/response"
	"go-rcon-panel/internal/errs"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

var statusByKind = map[errs.Kind]int{
	errs.Internal:        http.StatusInternalServerError,
	errs.InvalidInput:    http.StatusBadRequest,
	errs.Unauthenticated: http.StatusUnauthorized,
	errs.Forbidden:       http.StatusForbidden,
	errs.NotFound:        http.StatusNotFound,
	errs.Conflict:        http.StatusConflict,
}

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

		statusCode, ok := statusByKind[appErr.Kind]
		if !ok {
			statusCode = http.StatusInternalServerError
		}

		if statusCode >= http.StatusInternalServerError {
			log.Error().
				Err(err).
				Str("method", c.Request.Method).
				Str("path", c.Request.URL.Path).
				Msg("request failed")
		}

		c.AbortWithStatusJSON(statusCode, response.Response{
			Success: false,
			Error: &response.ErrorInfo{
				Code:    appErr.Status,
				Message: appErr.Message,
				Fields:  appErr.Fields,
			},
		})
	}
}

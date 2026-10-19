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
	errs.BadRequest:       http.StatusBadRequest,
	errs.Unauthenticated:  http.StatusUnauthorized,
	errs.Forbidden:        http.StatusForbidden,
	errs.NotFound:         http.StatusNotFound,
	errs.MethodNotAllowed: http.StatusMethodNotAllowed,
	errs.Conflict:         http.StatusConflict,
	errs.Internal:         http.StatusInternalServerError,
}

func httpStatus(k errs.Kind) int {
	if s, ok := statusByKind[k]; ok {
		return s
	}
	return http.StatusInternalServerError
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

		status := httpStatus(appErr.Kind)
		code, message, fields := appErr.Code, appErr.Message, appErr.Fields

		if status >= http.StatusInternalServerError {
			log.Error().
				Err(err).
				Str("method", c.Request.Method).
				Str("path", c.Request.URL.Path).
				Msg("request failed")

			code = errs.ErrInternalServer.Code
			message = errs.ErrInternalServer.Message
			fields = nil
		}

		c.AbortWithStatusJSON(status, response.Response{
			Success: false,
			Error: &response.ErrorInfo{
				Code:    code,
				Message: message,
				Fields:  fields,
			},
		})
	}
}

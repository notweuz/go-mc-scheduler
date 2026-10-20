package middleware

import (
	"fmt"
	"go-rcon-panel/internal/auth"
	"go-rcon-panel/internal/config"
	"go-rcon-panel/internal/dto/header"
	"go-rcon-panel/internal/errs"
	"strings"

	"github.com/gin-gonic/gin"
)

var (
	ErrTokenMissing = errs.ErrUnauthorized.WithMessage("authorization token is missing!")
)

const userIDContextKey = "user_id"

func Auth(cfg *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		var header header.Authorization
		if err := c.ShouldBindHeader(&header); err != nil {
			_ = c.Error(ErrTokenMissing.WithErr(err))
			c.Abort()
			return
		}

		scheme, token, ok := strings.Cut(header.Token, " ")
		if !ok || !strings.EqualFold(scheme, "Bearer") || token == "" {
			_ = c.Error(errs.ErrUnauthorized.WithMessage("invalid authorization header"))
			c.Abort()
			return
		}

		userID, err := auth.ParseToken(token, cfg.JWTSecret)
		if err != nil {
			_ = c.Error(fmt.Errorf("token parse error in auth middleware: %w", err))
			c.Abort()
			return
		}

		c.Set(userIDContextKey, userID)
	}
}

func GetUserIDFromContext(c *gin.Context) (int64, error) {
	value, exists := c.Get(userIDContextKey)
	if !exists {
		return 0, errs.ErrInternalServer.WithMessage("user id is missing in context")
	}

	userID, ok := value.(int64)
	if !ok || userID <= 0 {
		return 0, errs.ErrInternalServer.WithMessage("invalid user ID in context")
	}

	return userID, nil
}

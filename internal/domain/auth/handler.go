package auth

import (
	"context"
	"go-mc-scheduler/internal/dto/request"
	"go-mc-scheduler/internal/transport/http/respond"
	"go-mc-scheduler/internal/validation"

	"github.com/gin-gonic/gin"
)

type Service interface {
	Register(ctx context.Context, req request.AuthCredentials) (string, error)
	Login(ctx context.Context, req request.AuthCredentials) (string, error)
}

type Handler struct {
	service Service
}

func NewHandler(svc Service) *Handler {
	return &Handler{service: svc}
}

func (h *Handler) Login(c *gin.Context) {
	var req request.AuthCredentials
	if err := c.ShouldBindJSON(&req); err != nil {
		respond.Error(c, validation.FromBindError(err))
		return
	}
}

func (h *Handler) Register(c *gin.Context) {
	var req request.AuthCredentials
	if err := c.ShouldBindJSON(&req); err != nil {
		respond.Error(c, validation.FromBindError(err))
		return
	}
}

func (h *Handler) RegisterRoutes(r *gin.RouterGroup) {
	auth := r.Group("/auth")
	auth.POST("/login", h.Login)
	auth.POST("/register", h.Register)
}

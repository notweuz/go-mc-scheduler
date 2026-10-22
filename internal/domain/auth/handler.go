package auth

import (
	"context"
	"rcon-scheduler-panel/internal/dto/request"
	"rcon-scheduler-panel/internal/dto/response"
	"rcon-scheduler-panel/internal/transport/http/respond"
	"rcon-scheduler-panel/internal/validation"

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

	token, err := h.service.Login(c.Request.Context(), req)
	if err != nil {
		respond.Error(c, err)
		return
	}
	respond.OK(c, response.AuthResponse{Token: token})
}

func (h *Handler) Register(c *gin.Context) {
	var req request.AuthCredentials
	if err := c.ShouldBindJSON(&req); err != nil {
		respond.Error(c, validation.FromBindError(err))
		return
	}

	token, err := h.service.Register(c.Request.Context(), req)
	if err != nil {
		respond.Error(c, err)
		return
	}
	respond.Created(c, response.AuthResponse{Token: token})
}

func (h *Handler) RegisterRoutes(r *gin.RouterGroup) {
	auth := r.Group("/auth")
	auth.POST("/login", h.Login)
	auth.POST("/register", h.Register)
}

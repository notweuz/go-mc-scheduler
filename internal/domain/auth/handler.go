package auth

import (
	"go-mc-scheduler/internal/dto/request"
	"go-mc-scheduler/internal/errs"
	"go-mc-scheduler/internal/transport/http/respond"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	// svc in future
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) Login(c *gin.Context) {
	var req request.AuthCredentials
	if err := c.ShouldBindJSON(&req); err != nil {
		respond.Error(c, errs.ErrBadRequest.WithErr(err))
		return
	}
}

func (h *Handler) Register(c *gin.Context) {
	var req request.AuthCredentials
	if err := c.ShouldBindJSON(&req); err != nil {
		respond.Error(c, err)
		return
	}
}

func (h *Handler) RegisterRoutes(r *gin.RouterGroup) {
	auth := r.Group("/auth")
	auth.POST("/login", h.Login)
	auth.POST("/register", h.Register)
}

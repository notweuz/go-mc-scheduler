package auth

import (
	"go-mc-scheduler/internal/errs"
	"go-mc-scheduler/internal/model/request"
	"go-mc-scheduler/internal/model/response"

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
		response.Fail(c, errs.ErrBadRequest, nil)
		return
	}
}

func (h *Handler) Register(c *gin.Context) {
	var req request.AuthCredentials
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, errs.ErrBadRequest, nil)
		return
	}
}

func (h *Handler) RegisterRoutes(r *gin.RouterGroup) {
	auth := r.Group("/auth")
	auth.POST("/login", h.Login)
	auth.POST("/register", h.Register)
}

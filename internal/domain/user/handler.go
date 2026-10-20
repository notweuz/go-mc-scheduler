package user

import (
	"context"
	"go-rcon-panel/internal/middleware"
	"go-rcon-panel/internal/model"
	"go-rcon-panel/internal/transport/http/respond"

	"github.com/gin-gonic/gin"
)

type Service interface {
	GetByID(ctx context.Context, id int64) (*model.User, error)
	GetByUsername(ctx context.Context, username string) (*model.User, error)
	Create(ctx context.Context, user *model.User) error
}

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) GetMe(c *gin.Context) {
	userID, err := middleware.GetUserIDFromContext(c)
	if err != nil {
		respond.Error(c, err)
		return
	}
	user, err := h.service.GetByID(c.Request.Context(), userID)
	if err != nil {
		respond.Error(c, err)
		return
	}
	respond.OK(c, user) // TODO: Make user DTO's
}

func (h *Handler) RegisterRoutes(r *gin.RouterGroup) {
	auth := r.Group("/users")
	auth.POST("/@me", h.GetMe)
}

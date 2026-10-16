package auth

import "github.com/gin-gonic/gin"

type Handler struct {
	// svc in future
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) Login(c *gin.Context) {

}

func (h *Handler) Register(c *gin.Context) {

}

func (h *Handler) RegisterRoutes(r *gin.RouterGroup) {
	auth := r.Group("/auth")
	auth.POST("/login", h.Login)
	auth.POST("/register", h.Register)
}
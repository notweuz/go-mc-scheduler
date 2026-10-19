package respond

import (
	"go-rcon-panel/internal/dto/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func OK(c *gin.Context, data any) {
	c.JSON(http.StatusOK, response.Response{
		Success: true,
		Data:    data,
	})
}

func Created(c *gin.Context, data any) {
	c.JSON(http.StatusCreated, response.Response{
		Success: true,
		Data:    data,
	})
}

func Error(c *gin.Context, err error) {
	_ = c.Error(err)
	c.Abort()
}

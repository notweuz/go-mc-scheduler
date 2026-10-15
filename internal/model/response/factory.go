package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func OK(c *gin.Context, status int, data any) {
	c.JSON(status, Response{Success: true, Data: data})
}

func Paged(c *gin.Context, data any, meta Meta) {
	c.JSON(http.StatusOK, Response{Success: true, Data: data, Meta: &meta})
}

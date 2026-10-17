package response

import (
	"go-mc-scheduler/internal/errs"
	"net/http"

	"github.com/gin-gonic/gin"
)

func OK(c *gin.Context, status int, data any) {
	c.JSON(status, Response{Success: true, Data: data})
}

func Fail(c *gin.Context, err *errs.AppError, meta *Meta) {
	c.JSON(err.StatusCode, Response{
		Success: false,
		Error: &ErrorInfo{
			Code:    err.Status,
			Message: err.Message,
		},
		Meta: meta,
	})
}

func Paged(c *gin.Context, data any, meta Meta) {
	c.JSON(http.StatusOK, Response{Success: true, Data: data, Meta: &meta})
}

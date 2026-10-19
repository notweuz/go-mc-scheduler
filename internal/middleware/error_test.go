package middleware

import (
	"errors"
	"fmt"
	"go-rcon-panel/internal/errs"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestErrorHandlerStatus(t *testing.T) {
	gin.SetMode(gin.TestMode)

	tests := []struct {
		name string
		err  error
		want int
	}{
		{"bad request", errs.ErrBadRequest, http.StatusBadRequest},
		{"method not allowed", errs.ErrMethodNotAllowed, http.StatusMethodNotAllowed},
		{"wrapped not found", fmt.Errorf("ctx: %w", errs.ErrNotFound), http.StatusNotFound},
		{"conflict", errs.ErrConflict, http.StatusConflict},
		{"unknown error", errors.New("boom"), http.StatusInternalServerError},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := gin.New()
			r.Use(ErrorHandler())
			r.GET("/", func(c *gin.Context) {_ = c.Error(tt.err)})

			w := httptest.NewRecorder()
			r.ServeHTTP(w, httptest.NewRequest(http.MethodGet, "/", nil))

			if w.Code != tt.want {
				t.Errorf("got status %d, want %d", w.Code, tt.want)
			}
		})
	}
}

package errs

import "net/http"

type AppError struct {
	StatusCode int    `json:"-"`
	Status     string `json:"status"`
	Message    string `json:"message"`
	Err        error
	Fields     map[string]string `json:"fields,omitempty"`
}

func (e *AppError) Error() string {
	if e.Err != nil {
		return e.Message + ": " + e.Err.Error()
	}
	return e.Message
}

func (e *AppError) Unwrap() error { return e.Err }

func (e *AppError) WithMessage(msg string) *AppError {
	cp := *e
	cp.Message = msg
	return &cp
}

func (e *AppError) WithErr(err error) *AppError {
	cp := *e
	cp.Err = err
	return &cp
}

func (e *AppError) WithFields(f map[string]string) *AppError {
	cp := *e
	cp.Fields = f
	return &cp
}

func NewAppError(statusCode int, status, message string) *AppError {
	return &AppError{StatusCode: statusCode, Status: status, Message: message}
}

var (
	ErrInternalServer   = NewAppError(http.StatusInternalServerError, "INTERNAL_SERVER_ERROR", "internal server error")
	ErrNotFound         = NewAppError(http.StatusNotFound, "NOT_FOUND", "resource not found")
	ErrUnauthorized     = NewAppError(http.StatusUnauthorized, "UNAUTHORIZED", "authentication required")
	ErrBadRequest       = NewAppError(http.StatusBadRequest, "BAD_REQUEST", "invalid request")
	ErrMethodNotAllowed = NewAppError(http.StatusMethodNotAllowed, "METHOD_NOT_ALLOWED", "method not allowed")
)

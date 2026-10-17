package errs

type FieldError struct {
	Code  string `json:"code"`
	Param string `json:"param,omitempty"`
}

type AppError struct {
	Kind    Kind   `json:"-"`
	Status  string `json:"status"`
	Message string `json:"message"`
	Err     error
	Fields  map[string]FieldError `json:"fields,omitempty"`
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

func (e *AppError) WithFields(f map[string]FieldError) *AppError {
	cp := *e
	cp.Fields = f
	return &cp
}

func New(kind Kind, status, message string) *AppError {
	return &AppError{Kind: kind, Status: status, Message: message}
}

var (
	ErrInternalServer   = New(Internal, "INTERNAL_SERVER_ERROR", "internal server error")
	ErrNotFound         = New(NotFound, "NOT_FOUND", "resource not found")
	ErrUnauthorized     = New(Unauthenticated, "UNAUTHORIZED", "authentication required")
	ErrBadRequest       = New(BadRequest, "BAD_REQUEST", "invalid request")
	ErrMethodNotAllowed = New(MethodNotAllowed, "METHOD_NOT_ALLOWED", "method not allowed")
)

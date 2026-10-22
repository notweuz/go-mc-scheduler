package validation

import (
	"encoding/json"
	"errors"
	"io"
	"rcon-scheduler-panel/internal/errs"

	"github.com/go-playground/validator/v10"
)

func FromBindError(err error) *errs.AppError {
	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		fields := make(map[string]errs.FieldError, len(ve))
		for _, fe := range ve {
			fields[fe.Field()] = errs.FieldError{Code: fe.Tag(), Param: fe.Param()}
		}
		return errs.ErrBadRequest.
			WithMessage("validation failed").
			WithErr(err).
			WithFields(fields)
	}

	var typeErr *json.UnmarshalTypeError
	if errors.As(err, &typeErr) {
		return errs.ErrBadRequest.
			WithMessage("invalid field type").
			WithFields(map[string]errs.FieldError{
				typeErr.Field: {Code: "type", Param: typeErr.Type.String()},
			}).
			WithErr(err)
	}

	var syntaxErr *json.SyntaxError
	if errors.As(err, &syntaxErr) || errors.Is(err, io.EOF) || errors.Is(err, io.ErrUnexpectedEOF) {
		return errs.ErrBadRequest.WithMessage("malformed JSON").WithErr(err)
	}

	return errs.ErrBadRequest.WithErr(err)
}

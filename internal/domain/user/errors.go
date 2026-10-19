package user

import "go-rcon-panel/internal/errs"

var (
	ErrUserNotFound = errs.ErrNotFound.WithMessage("user not found")
	ErrSameUserExist = errs.ErrConflict.WithMessage("user with same username already exist")
)
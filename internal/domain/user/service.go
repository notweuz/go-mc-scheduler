package user

import (
	"context"
	"go-rcon-panel/internal/model"
)

type Repository interface {
	GetByID(ctx context.Context, id int64) (*model.User, error)
	Create(ctx context.Context, user *model.User) error
}
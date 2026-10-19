package user

import (
	"context"
	"go-rcon-panel/internal/model"
)

type Service interface {
	GetByID(ctx context.Context, id int64) (*model.User, error)
	GetByUsername(ctx context.Context, username string) (*model.User, error)
	Create(ctx context.Context, user *model.User) error
}

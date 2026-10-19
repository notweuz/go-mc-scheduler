package auth

import (
	"context"
	"go-rcon-panel/internal/dto/request"
)

type service struct {
}

func NewService() Service {
	return &service{}
}

func (s *service) Register(ctx context.Context, req request.AuthCredentials) (string, error) {
	return "", nil
}

func (s *service) Login(ctx context.Context, req request.AuthCredentials) (string, error) {
	return "", nil
}

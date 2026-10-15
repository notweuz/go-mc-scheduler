package user

import (
	"go-mc-scheduler/internal/model"
	"time"
)

type userModel struct {
	ID       int64  `gorm:"primaryKey"`
	Username string `gorm:"not null"`
	Password string `gorm:"not null"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (m *userModel) toDomain() *model.User {
	return &model.User{
		ID:        m.ID,
		Username:  m.Username,
		Password:  m.Password,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}
}

func FromDomain(u *model.User) *userModel {
	return &userModel{
		ID:        u.ID,
		Username:  u.Username,
		Password:  u.Password,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}

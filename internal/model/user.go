package model

import (
	// "go-mc-scheduler/internal/domain/user"
	"time"
)

type User struct {
	ID       int64
	Username string
	Password string

	CreatedAt time.Time
	UpdatedAt time.Time
}

// func (u *User) ToModel() *user.UserModel {
// 	return &user.UserModel{
// 		ID:        u.ID,
// 		Username:  u.Username,
// 		Password:  u.Password,
// 		CreatedAt: u.CreatedAt,
// 		UpdatedAt: u.UpdatedAt,
// 	}
// }

// func (u *User) FromDomain(domain *user.UserModel) *User {
// 	return &User{
// 		ID:        domain.ID,
// 		Username:  domain.Username,
// 		Password:  domain.Password,
// 		CreatedAt: domain.CreatedAt,
// 		UpdatedAt: domain.UpdatedAt,
// 	}
// }

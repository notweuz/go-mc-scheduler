package model

import "time"

type Server struct {
	ID          int64  `gorm:"primaryKey"`
	Name        string `gorm:"not null"`
	Address     string `gorm:"not null"`
	Password    string `gorm:"not null"`
	ProtocolID  int64  `gorm:"not null"`
	Protocol    *Protocol
	Description *string

	CreatedAt time.Time `gorm:"not null"`
	UpdatedAt time.Time `gorm:"not null"`
}

func (Server) TableName() string {
	return "servers"
}

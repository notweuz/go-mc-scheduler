package protocol

import "go-mc-scheduler/internal/model"

type protocolModel struct {
	ID          int64  `gorm:"primaryKey"`
	Name        string `gorm:"not null"`
	DefaultPort *int
	Description *string
}

func (protocolModel) TableName() string {
	return "protocols"
}

func (m *protocolModel) toDomain() *model.Protocol {
	return &model.Protocol{
		ID:          m.ID,
		Name:        m.Name,
		DefaultPort: m.DefaultPort,
		Description: m.Description,
	}
}

func FromDomain(p *model.Protocol) *protocolModel {
	return &protocolModel{
		ID:          p.ID,
		Name:        p.Name,
		DefaultPort: p.DefaultPort,
		Description: p.Description,
	}
}
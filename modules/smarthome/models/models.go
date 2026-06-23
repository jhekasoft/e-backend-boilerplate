package models

import (
	"github.com/jhekasoft/e-backend/crud"

	"github.com/jackc/pgtype"
)

type SmartHomeSensorValue struct {
	crud.Model
	Name1  string       `gorm:"index"`
	Name2  string       `gorm:"index"`
	Name3  string       `gorm:"index"`
	Sensor string       `gorm:"index"`
	Value  pgtype.JSONB `gorm:"type:jsonb;default:'{}';not null"`
}

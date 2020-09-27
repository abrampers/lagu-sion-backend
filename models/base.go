package models

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type Base struct {
	Id uuid.UUID `gorm:"type:uuid;primary_key;"`
}

func (base *Base) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("Id", uuid.New())
	return nil
}

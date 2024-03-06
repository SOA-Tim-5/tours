package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Equipment struct {
	Id    uuid.UUID 
	Name string
	Description string
}

func (equipment *Equipment) BeforeCreate(scope *gorm.DB) error {
	equipment.Id = uuid.New()
	return nil
}

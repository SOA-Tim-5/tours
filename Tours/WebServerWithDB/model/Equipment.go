package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Equipment struct {
	Id    int64 
	Name string
	Description string
}

func (equipment *Equipment) BeforeCreate(scope *gorm.DB) error {
	currentTimestamp := time.Now().UnixNano() / int64(time.Microsecond)
    uniqueID := uuid.New().ID()
    equipment.Id = currentTimestamp + int64(uniqueID)
	return nil
}

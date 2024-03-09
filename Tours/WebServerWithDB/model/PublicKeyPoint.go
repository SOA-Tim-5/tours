package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PublicKeyPoint struct {
	Id    int64
	Name string
	Description string
	Longitude float64
	Latitude float64
	ImagePath string
	Order int64
	LocationAddress string
}

func (keyPoint *PublicKeyPoint) BeforeCreate(scope *gorm.DB) error {
	currentTimestamp := time.Now().UnixNano() / int64(time.Microsecond)
    uniqueID := uuid.New().ID()
    keyPoint.Id = currentTimestamp + int64(uniqueID)
	return nil
}

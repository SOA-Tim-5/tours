package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)




type TouristPosition struct {
	Id    int64
	TouristId  int64 
	Longitude float64
	Latitude float64


}

func (touristPosition *TouristPosition) BeforeCreate(scope *gorm.DB) error {
	currentTimestamp := time.Now().UnixNano() / int64(time.Microsecond)
    uniqueID := uuid.New().ID()
    touristPosition.Id = currentTimestamp + int64(uniqueID)

	return nil
}

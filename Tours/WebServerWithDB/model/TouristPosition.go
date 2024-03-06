package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)




type TouristPosition struct {
	Id    uuid.UUID 
	TouristId  int64 
	Longitude float64
	Latitude float64


}

func (touristPosition *TouristPosition) BeforeCreate(scope *gorm.DB) error {
	touristPosition.Id = uuid.New()

	return nil
}

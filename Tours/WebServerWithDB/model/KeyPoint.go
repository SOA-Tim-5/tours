package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type KeyPoint struct {
	Id    uuid.UUID 
	TourId  int64    
	Tour Tour //?
	Name string
	Description string
	Longitude float64
	Latitude float64
	LocationAddress string
	ImagePath string
	Order int64
	HaveSecret bool
	Secret KeyPointSecret `gorm:"type:jsonb;"`
	IsEncounterRequired bool
	HasEncounter bool
}

func (keyPoint *KeyPoint) BeforeCreate(scope *gorm.DB) error {
	keyPoint.Id = uuid.New()
	return nil
}

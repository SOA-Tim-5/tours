package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type KeyPoint struct {
	Id                  int64
	TourId              int64
	Tour                Tour //?
	Name                string
	Description         string
	Longitude           float64
	Latitude            float64
	LocationAddress     string
	ImagePath           string
	Order               int64
	HaveSecret          bool
	Secret              KeyPointSecret `gorm:"type:jsonb;"`
	IsEncounterRequired bool
	HasEncounter        bool
}

func (keyPoint *KeyPoint) BeforeCreate(scope *gorm.DB) error {
	currentTimestamp := time.Now().UnixNano() / int64(time.Microsecond)
	uniqueID := uuid.New().ID()
	keyPoint.Id = currentTimestamp + int64(uniqueID)
	return nil
}

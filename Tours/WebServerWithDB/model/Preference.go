package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)



type Preference struct {
	Id    int64
	UserId int64
	DifficultyLevel int
	WalkingRating int
	CyclingRating int
	CarRating int
	BoatRating int
	SelectedTags []string `gorm:"type:text[]"`



}

func (preference *Preference) BeforeCreate(scope *gorm.DB) error {
	currentTimestamp := time.Now().UnixNano() / int64(time.Microsecond)
    uniqueID := uuid.New().ID()
    preference.Id = currentTimestamp + int64(uniqueID)
	return nil
}

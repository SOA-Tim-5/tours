package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)



type Preference struct {
	Id    uuid.UUID  
	UserId int64
	DifficultyLevel int
	WalkingRating int
	CyclingRating int
	CarRating int
	BoatRating int
	SelectedTags []string `gorm:"type:text[]"`



}

func (preference *Preference) BeforeCreate(scope *gorm.DB) error {
	preference.Id = uuid.New()
	return nil
}

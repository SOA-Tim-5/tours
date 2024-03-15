package model

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Preference struct {
	Id              int64
	UserId          int64
	DifficultyLevel int
	WalkingRating   int
	CyclingRating   int
	CarRating       int
	BoatRating      int
	SelectedTags    pq.StringArray `gorm:"type:text[]"`
}

func (preference *Preference) BeforeCreate(scope *gorm.DB) error {
	if err := preference.Validate(); err != nil {
		return err
	}

	currentTimestamp := time.Now().UnixNano() / int64(time.Microsecond)
	uniqueID := uuid.New().ID()
	preference.Id = currentTimestamp + int64(uniqueID)
	return nil
}

func (preference *Preference) Validate() error {
	if preference.DifficultyLevel < 1 || preference.DifficultyLevel > 5 {
		return errors.New("invalid difficulty level")
	}
	if preference.WalkingRating < 0 || preference.WalkingRating > 3 {
		return errors.New("invalid walking rating")
	}
	if preference.CyclingRating < 0 || preference.CyclingRating > 3 {
		return errors.New("invalid cycling rating")
	}
	if preference.BoatRating < 0 || preference.BoatRating > 3 {
		return errors.New("invalid boat rating")
	}
	if preference.CarRating < 0 || preference.CarRating > 3 {
		return errors.New("invalid car rating")
	}
	return nil
}

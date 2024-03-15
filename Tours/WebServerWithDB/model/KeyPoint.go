package model

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type KeyPoint struct {
	Id     int64
	TourId int64
	//Tour                Tour //?
	Name                string
	Description         string
	Longitude           float64
	Latitude            float64
	LocationAddress     string
	ImagePath           string
	Order               int64
	HaveSecret          bool
	Secret              KeyPointSecret `gorm:"-"`
	IsEncounterRequired bool
	HasEncounter        bool
}

func (keyPoint *KeyPoint) BeforeCreate(scope *gorm.DB) error {

	if err := keyPoint.Validate(); err != nil {
		return err
	}

	currentTimestamp := time.Now().UnixNano() / int64(time.Microsecond)
	uniqueID := uuid.New().ID()
	keyPoint.Id = currentTimestamp + int64(uniqueID)
	println("okoko")
	return nil
}

func (kp *KeyPoint) Validate() error {
	if kp.Longitude < -180 || kp.Longitude > 180 {
		return errors.New("invalid Longitude")
	}
	if kp.Latitude < -90 || kp.Latitude > 90 {
		return errors.New("invalid Latitude")
	}
	if kp.ImagePath == "" {
		return errors.New("invalid ImagePath")
	}
	return nil
}

package model

import (
	"encoding/json"
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
	Secret              []byte `gorm:"column:secret;type:jsonb;"`
	IsEncounterRequired bool
	HasEncounter        bool
}

func (kp *KeyPoint) ParseSecret() (*KeyPointSecret, error) {
	var secret KeyPointSecret
	err := json.Unmarshal(kp.Secret, &secret)
	if err != nil {
		return nil, err
	}
	return &secret, nil
}

func (keyPoint *KeyPoint) BeforeCreate(scope *gorm.DB) error {
	currentTimestamp := time.Now().UnixNano() / int64(time.Microsecond)
	uniqueID := uuid.New().ID()
	keyPoint.Id = currentTimestamp + int64(uniqueID)
	return nil
}

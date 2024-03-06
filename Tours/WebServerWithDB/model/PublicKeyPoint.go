package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"fmt"
	"time"
)

type PublicKeyPoint struct {
	Id    uuid.UUID 
	Name string
	Description string
	Longitude float64
	Latitude float64
	ImagePath string
	Order int64
	LocationAddress string
}

func (keyPoint *PublicKeyPoint) BeforeCreate(scope *gorm.DB) error {
	keyPoint.Id = uuid.New()
	return nil
}

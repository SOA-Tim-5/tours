package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Subscriber struct {
	Id    uuid.UUID 
	TouristId  int64    
	EmailAddress string
	Frequency int
	LastTimeSent time.Time


}

func (subscriber *Subscriber) BeforeCreate(scope *gorm.DB) error {
	subscriber.Id = uuid.New()

	return nil
}

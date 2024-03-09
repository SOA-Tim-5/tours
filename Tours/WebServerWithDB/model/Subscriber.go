package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Subscriber struct {
	Id    int64
	TouristId  int64    
	EmailAddress string
	Frequency int
	LastTimeSent time.Time


}

func (subscriber *Subscriber) BeforeCreate(scope *gorm.DB) error {
	currentTimestamp := time.Now().UnixNano() / int64(time.Microsecond)
    uniqueID := uuid.New().ID()
    subscriber.Id = currentTimestamp + int64(uniqueID)

	return nil
}

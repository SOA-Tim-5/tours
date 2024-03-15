package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)



type PublicFacilityNotification struct {
	Id    int64
	RequestId int64
	Description string
	AuthorId int64
	Created time.Time
	IsAccepted bool
	Comment string
	IsSeen bool
	SenderPicture string
	SenderName string
	Header string
}

func (notification *PublicFacilityNotification) BeforeCreate(scope *gorm.DB) error {
	currentTimestamp := time.Now().UnixNano() / int64(time.Microsecond)
    uniqueID := uuid.New().ID()
    notification.Id = currentTimestamp + int64(uniqueID)
	return nil
}

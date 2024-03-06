package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)



type PublicKeyPointNotification struct {
	Id    uuid.UUID 
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

func (notification *PublicKeyPointNotification) BeforeCreate(scope *gorm.DB) error {
	notification.Id = uuid.New()
	return nil
}

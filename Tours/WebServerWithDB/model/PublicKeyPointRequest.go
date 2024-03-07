package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PublicStatus int

const (
    Pending PublicStatus = iota
	Accepted
	Rejected
)

type PublicKeyPointRequest struct {
	Id    int64
	KeyPointId int64
	AuthorId int64
	Status PublicStatus
	Comment string
	Created time.Time
	AuthorName string
}

func (request *PublicKeyPointRequest) BeforeCreate(scope *gorm.DB) error {
	currentTimestamp := time.Now().UnixNano() / int64(time.Microsecond)
    uniqueID := uuid.New().ID()
    request.Id = currentTimestamp + int64(uniqueID)
	return nil
}

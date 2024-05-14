package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PublicFacilityRequest struct {
	Id            int64
	FacilityId    int64
	AuthorId      int64
	Status        PublicStatus
	Comment       string
	Created       int64
	AuthorName    string
	AuthorPicture string
}

func (request *PublicFacilityRequest) BeforeCreate(scope *gorm.DB) error {
	currentTimestamp := time.Now().UnixNano() / int64(time.Microsecond)
	uniqueID := uuid.New().ID()
	request.Id = currentTimestamp + int64(uniqueID)
	return nil
}

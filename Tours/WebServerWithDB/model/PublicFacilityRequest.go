package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)


type PublicFacilityRequest struct {
	Id    uuid.UUID 
	FacilityId int64
	AuthorId int64
	Status PublicStatus
	Comment string
	Created time.Time
	AuthorName string
}

func (request *PublicFacilityRequest) BeforeCreate(scope *gorm.DB) error {
	request.Id = uuid.New()
	return nil
}

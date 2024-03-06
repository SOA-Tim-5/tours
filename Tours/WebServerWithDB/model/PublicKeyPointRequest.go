package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"fmt"
	"time"
)

type PublicStatus int

const (
    Pending PublicStatus = iota
	Accepted
	Rejected
)

type PublicKeyPointRequest struct {
	Id    uuid.UUID 
	KeyPointId int64
	AuthorId int64
	Status PublicStatus
	Comment string
	Created time.Time
	AuthorName string
}

func (request *PublicKeyPointRequest) BeforeCreate(scope *gorm.DB) error {
	request.Id = uuid.New()
	return nil
}

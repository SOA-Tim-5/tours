package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Review struct {
	Id    int64
	Rating  int   
	Comment string
	TouristId int64
	TourVisitDate time.Time //?
	CommentDate time.Time //?
	TourId int64
	Images []string `gorm:"type:text[]"`
}

func (review *Review) BeforeCreate(scope *gorm.DB) error {
	currentTimestamp := time.Now().UnixNano() / int64(time.Microsecond)
    uniqueID := uuid.New().ID()
    review.Id = currentTimestamp + int64(uniqueID)
	return nil
}

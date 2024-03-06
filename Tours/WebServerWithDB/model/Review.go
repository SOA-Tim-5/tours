package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Review struct {
	Id    uuid.UUID 
	Rating  int   
	Comment string
	TouristId int64
	TourVisitDate time.Time //?
	CommentDate time.Time //?
	TourId int64
	Images []string `gorm:"type:text[]"`
}

func (review *Review) BeforeCreate(scope *gorm.DB) error {
	review.Id = uuid.New()
	return nil
}

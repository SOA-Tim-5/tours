package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"fmt"
	"time"
)

type Review struct {
	Id    uuid.UUID 
	Rating  int   
	Comment string
	TouristId int64
	TourVisitDate time.Time //?
	CommentDate time.Time //?
	TourId int64
	Images []string
}

func (review *Review) BeforeCreate(scope *gorm.DB) error {
	review.Id = uuid.New()
	return nil
}

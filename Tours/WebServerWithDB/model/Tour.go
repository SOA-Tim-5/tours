package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"fmt"
	"time"
)

type TourStatus int

const (
    Draft TourStatus = iota
    Published
    Archived
	Ready
)

type TourCategory int

const (
    Adventure TourCategory = iota 
	FamilyTrips
	Cruise
	Cultural
)




type Tour struct {
	Id    uuid.UUID 
	AuthorId  int64    
	Name string
	Description string
	Difficulty int
	Tags []string
	Status TourStatus
	Price float64
	IsDeleted bool
	Distance float64
	PublishDate time.Time	//?
	ArchiveDate time.Time 
	EquipmentList []Equipment{}
	KeyPoints []KeyPoint{}
	Durations []TourDuration{} `gorm:"type:jsonb;"`
	Reviews []Review
	Category TourCategory

}

func (keyPoint *KeyPoint) BeforeCreate(scope *gorm.DB) error {
	keyPoint.Id = uuid.New()
	return nil
}

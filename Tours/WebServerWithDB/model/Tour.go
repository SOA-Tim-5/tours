package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
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
	Tags []string `gorm:"type:text[]"`
	Status TourStatus
	Price float64
	IsDeleted bool
	Distance float64
	PublishDate time.Time	//?
	ArchiveDate time.Time 
	EquipmentList []Equipment `gorm:"type:Equipment[]"`
	KeyPoints []KeyPoint 
	Durations []TourDuration `gorm:"type:jsonb;"`
	Reviews []Review 
	Category TourCategory

}

func (tour *Tour) BeforeCreate(scope *gorm.DB) error {
	tour.Id = uuid.New()

	return nil
}

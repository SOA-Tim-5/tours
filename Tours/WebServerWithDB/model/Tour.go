package model

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
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
	Id            int64
	AuthorId      int64
	Name          string
	Description   string
	Difficulty    int
	Tags          pq.StringArray `gorm:"type:text[]"`
	Status        TourStatus
	Price         float64
	IsDeleted     bool
	Distance      float64
	PublishDate   time.Time //?
	ArchiveDate   time.Time
	EquipmentList []Equipment `gorm:"type:Equipment[]"`
	KeyPoints     []KeyPoint
	Durations     []TourDuration `gorm:"type:jsonb;"`
	Reviews       []Review
	Category      TourCategory
}

func (tour *Tour) BeforeCreate(scope *gorm.DB) error {

	if err := tour.Validate(); err != nil {
        return err
    }


	currentTimestamp := time.Now().UnixNano() / int64(time.Microsecond)
	uniqueID := uuid.New().ID()
	tour.Id = currentTimestamp + int64(uniqueID)
	//tour.KeyPoints = make([]KeyPoint, 0)
	/*if len(tour.Tags) > 0 {
		//tour.Tags = []string{fmt.Sprintf("{%s}", strings.Join(tour.Tags, ","))}
	}*/
	//println("okoko")
	return nil

}


func (tour *Tour) Validate() error {
    if tour.Name == "" {
        return errors.New("invalid name")
    }
    if tour.Description == "" {
        return errors.New("invalid description")
    }
    if tour.Difficulty < 1 || tour.Difficulty > 5 {
        return errors.New("invalid difficulty")
    }
    //if len(tour.Tags) == 0 {
    //    return errors.New("Tags cannot be empty")
    //}
    if tour.Price < 0 {
        return errors.New("price cannot be negative")
    }
    return nil
}



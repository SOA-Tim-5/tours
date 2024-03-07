package model

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type FacilityCategory int

const (
	Restaurant FacilityCategory = iota
	ParkingLot
	Toilet
	Hospital
	Cafe
	Pharmacy
	ExchangeOffice
	BusStop
	Shop
	Other
)

type Facility struct {
	Id          uuid.UUID
	Name        string
	Description string
	ImagePath   string
	Tags        []string `gorm:"type:text[]"`
	AuthorId    int
	Category    FacilityCategory
	Longitude   float64
	Latitude    float64
}

func (facility *Facility) BeforeCreate(scope *gorm.DB) error {
	facility.Id = uuid.New()
	if len(facility.Tags) > 0 {
		facility.Tags = []string{fmt.Sprintf("{%s}", strings.Join(facility.Tags, ","))}
	}
	return nil
}

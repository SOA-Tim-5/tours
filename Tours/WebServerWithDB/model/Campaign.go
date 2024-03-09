package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)




type Campaign struct {
	Id    int64 
	TouristId  int64    
	Name string
	Description string
	Distance float64
	AverageDifficulty float64
	TourIds []int64 `gorm:"type:text[]"`
	EquipmentIds []int64 `gorm:"type:text[]"`
	KeyPointIds []int64 `gorm:"type:text[]"`
	Equipments []Equipment `gorm:"type:Equipment[]"`
	KeyPoints []KeyPoint `gorm:"many2many:campaign_key_points;"`


}

func (campaign *Campaign) BeforeCreate(scope *gorm.DB) error {
	currentTimestamp := time.Now().UnixNano() / int64(time.Microsecond)
    uniqueID := uuid.New().ID()
    campaign.Id = currentTimestamp + int64(uniqueID)
	return nil
}

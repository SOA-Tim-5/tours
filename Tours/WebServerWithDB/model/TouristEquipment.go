package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)




type TouristEquipment struct {
	Id    int64 
	TouristId  int64 
	EquipmentIds []int `gorm:"type:int[]"`
	


}

func (touristEquipment *TouristEquipment) BeforeCreate(scope *gorm.DB) error {
	currentTimestamp := time.Now().UnixNano() / int64(time.Microsecond)
    uniqueID := uuid.New().ID()
    touristEquipment.Id = currentTimestamp + int64(uniqueID)

	return nil
}

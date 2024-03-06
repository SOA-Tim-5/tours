package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)




type TouristEquipment struct {
	Id    uuid.UUID 
	TouristId  int64 
	EquipmentIds []int `gorm:"type:int[]"`
	


}

func (touristEquipment *TouristEquipment) BeforeCreate(scope *gorm.DB) error {
	touristEquipment.Id = uuid.New()

	return nil
}

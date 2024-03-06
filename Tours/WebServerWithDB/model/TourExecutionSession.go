package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)


type TourExecutionSessionStatus int

const (
    Started TourExecutionSessionStatus = iota
    Abandoned
 	Completed
)


type TourExecutionSession struct {
	Id    uuid.UUID 
	Status TourExecutionSessionStatus
	TourId  int64 
	TouristId  int64 
	NextKeyPointId  int64 
	Progress float64
	LastActivity time.Time
	IsCampaign bool  


}

func (tourExecutionSession *TourExecutionSession) BeforeCreate(scope *gorm.DB) error {
	tourExecutionSession.Id = uuid.New()

	return nil
}

package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TransportType int

const (
    Walking TransportType = iota
	Bicycle
	Car
)


type TourDuration struct {
	Duration int
	TransportType TransportType
}

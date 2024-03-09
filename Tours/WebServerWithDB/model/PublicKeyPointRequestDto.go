package model

import (
	"time"
)



type PublicKeyPointRequestDto struct {
	Id    int64
	KeyPointId int64
	AuthorId int64
	Status PublicStatus
	Comment string
	Created time.Time
	AuthorName string
	AuthorUsername string
	KeyPointName string
	AuthorPicture string
}



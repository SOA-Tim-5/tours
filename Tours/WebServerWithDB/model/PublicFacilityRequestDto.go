package model

import (
	"time"
)

type PublicFacilityRequestDto struct {
	Id            int64
	FacilityId    int64
	AuthorId      int64
	Status        PublicStatus
	Comment       string
	Created       time.Time
	AuthorName    string
	FacilityName  string
	AuthorPicture string
}

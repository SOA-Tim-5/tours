package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type KeyPointSecret struct {
	Images []string
	Description string
}

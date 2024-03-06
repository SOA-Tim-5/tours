package model

type KeyPointSecret struct {
	Images      []string `gorm:"type:text[]"`
	Description string
}

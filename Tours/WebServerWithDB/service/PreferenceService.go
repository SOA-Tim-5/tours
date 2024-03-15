package service

import (
	"database-example/model"
	"database-example/repo"
	"fmt"
)

type PreferenceService struct {
	PreferenceRepo *repo.PreferenceRepository
}



func (service *PreferenceService) Create(preference *model.Preference) error {
	err := service.PreferenceRepo.CreatePreference(preference)
	if err != nil {
		return err
	}
	return nil
}
func (service *PreferenceService) GetByTouristId(touristId int64) (*model.Preference, error) {
	storedPreference, err := service.PreferenceRepo.GetByTouristId(touristId)
	if err != nil {
		return nil, fmt.Errorf("not found")
	}
	return &storedPreference, nil
}

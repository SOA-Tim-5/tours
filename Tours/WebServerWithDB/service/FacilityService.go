package service

import (
	"database-example/model"
	"database-example/repo"
)

type FacilityService struct {
	FacilityRepo *repo.FacilityRepository
}

/*func (service *TourService) FindStudent(id string) (*model.Student, error) {
	student, err := service.StudentRepo.FindById(id)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("menu item with id %s not found", id))
	}
	return &student, nil
}*/

func (service *FacilityService) Create(facility *model.Facility) error {
	err := service.FacilityRepo.CreateFacility(facility)
	if err != nil {
		return err
	}
	return nil
}

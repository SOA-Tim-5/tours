package service

import (
	"database-example/model"
	"database-example/repo"
	"fmt"
)

type TourService struct {
	TourRepo *repo.TourRepository
}

/*func (service *TourService) FindStudent(id string) (*model.Student, error) {
	student, err := service.StudentRepo.FindById(id)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("menu item with id %s not found", id))
	}
	return &student, nil
}*/

func (service *TourService) Create(tour *model.Tour) error {
	err := service.TourRepo.CreateTour(tour)
	if err != nil {
		return err
	}
	return nil
}
func (service *TourService) GetByAuthorId(authorId int64) ([]model.Tour, error) {
	storedTours, err := service.TourRepo.GetByAuthorId(authorId)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("tours with authorId %d not found", authorId))
	}
	return storedTours, nil
}

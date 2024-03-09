package service

import (
	"database-example/model"
	"database-example/repo"
	"fmt"
)

type TourService struct {
	TourRepo *repo.TourRepository
	KeypointRepo *repo.KeyPointRepository
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

	var retStoredTours []model.Tour

	//punjenje ture keypointima rucno
	for _, storedTour := range storedTours {
		var tourKeyPoints []model.KeyPoint
		tourKeyPoints, err = service.KeypointRepo.GetKeyPoints(storedTour.Id)

		if err == nil {
			for _, id := range tourKeyPoints {
				println(id.TourId)
			}
			storedTour.KeyPoints = tourKeyPoints		
		}
		retStoredTours = append(retStoredTours, storedTour)
		println(err)
	}
	return retStoredTours, nil
}
func (service *TourService) GetById(id int64) (*model.Tour, error) {
	tour, err := service.TourRepo.GetById(id)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("menu item with id %d not found", id))
	}
	return &tour, nil
}

package service

import (
	"database-example/model"
	"database-example/repo"
)

type KeyPointService struct {
	KeyPointRepo *repo.KeyPointRepository
}

/*func (service *TourService) FindStudent(id string) (*model.Student, error) {
	student, err := service.StudentRepo.FindById(id)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("menu item with id %s not found", id))
	}
	return &student, nil
}*/

func (service *KeyPointService) Create(keyPoint *model.KeyPoint) error {
	err := service.KeyPointRepo.CreateKeyPoint(keyPoint)
	if err != nil {
		return err
	}
	return nil
}

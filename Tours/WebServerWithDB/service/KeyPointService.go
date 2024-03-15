package service

import (
	"database-example/model"
	"database-example/repo"
	"fmt"
)

type KeyPointService struct {
	KeyPointRepo *repo.KeyPointRepository
}



func (service *KeyPointService) Create(keyPoint *model.KeyPoint) error {
	err := service.KeyPointRepo.CreateKeyPoint(keyPoint)
	if err != nil {
		return err
	}
	return nil
}

func (service *KeyPointService) GetKeyPoints(tourId int64) ([]model.KeyPoint, error) {
	storedKeyPoints, err := service.KeyPointRepo.GetKeyPoints(tourId)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("tours with authorId %d not found", tourId))
	}
	println("usaoo")
	return storedKeyPoints, nil
}



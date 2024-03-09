package service

import (
	"database-example/model"
	"database-example/repo"
	"fmt"
)

type PublicKeyPointRequestService struct {
	PublicKeyPointRequestRepo *repo.PublicKeyPointRequestRepository
	KeypointRepo *repo.KeyPointRepository
}


func (service *PublicKeyPointRequestService) Create(request *model.PublicKeyPointRequest) error {
	err := service.PublicKeyPointRequestRepo.CreatePublicKeyPointRequest(request)
	if err != nil {
		return err
	}
	return nil
}


func (service *PublicKeyPointRequestService) GetAll() ([]model.PublicKeyPointRequestDto, error) {
	storedRequests, err := service.PublicKeyPointRequestRepo.GetAll()
	if err != nil {
		return nil, fmt.Errorf("not found")
	}

	var requestsDtos []model.PublicKeyPointRequestDto

	for _, storedRequest := range storedRequests {
		var keyPoint model.KeyPoint
		keyPoint, err = service.KeypointRepo.GetById(storedRequest.KeyPointId)

		if err == nil {
			var requestDto model.PublicKeyPointRequestDto
			requestDto.AuthorId = storedRequest.AuthorId
			requestDto.KeyPointId = storedRequest.KeyPointId
			requestDto.Status = storedRequest.Status
			requestDto.Comment = storedRequest.Comment
			requestDto.Created = storedRequest.Created
			requestDto.AuthorName = storedRequest.AuthorName
			requestDto.KeyPointName = keyPoint.Name
			requestDto.AuthorPicture = storedRequest.AuthorPicture
			requestsDtos = append(requestsDtos, requestDto)
		}
	}

	return requestsDtos, nil
}

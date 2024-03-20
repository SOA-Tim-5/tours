package service

import (
	"database-example/model"
	"database-example/repo"
	"fmt"
)

type PublicFacilityRequestService struct {
	PublicFacilityRequestRepo *repo.PublicFacilityRequestRepository
	FacilityRepo              *repo.FacilityRepository
}

func (service *PublicFacilityRequestService) Create(request *model.PublicFacilityRequest) error {
	err := service.PublicFacilityRequestRepo.CreatePublicFacilityRequest(request)
	if err != nil {
		return err
	}
	return nil
}

func (service *PublicFacilityRequestService) GetAllFacilityRequest() ([]model.PublicFacilityRequestDto, error) {
	storedRequests, err := service.PublicFacilityRequestRepo.GetAllFacilityRequest()
	if err != nil {
		return nil, fmt.Errorf("not found")
	}

	var requestsDtos []model.PublicFacilityRequestDto

	for _, storedRequest := range storedRequests {
		var facility model.Facility
		facility, err = service.FacilityRepo.GetById(storedRequest.FacilityId)

		if err == nil {
			var requestDto model.PublicFacilityRequestDto
			requestDto.AuthorId = storedRequest.AuthorId
			requestDto.FacilityId = storedRequest.FacilityId
			requestDto.Status = storedRequest.Status
			requestDto.Comment = storedRequest.Comment
			requestDto.Created = storedRequest.Created
			requestDto.AuthorName = storedRequest.AuthorName
			requestDto.AuthorPicture = storedRequest.AuthorPicture
			requestDto.FacilityName = facility.Name
			requestsDtos = append(requestsDtos, requestDto)
		}
	}

	return requestsDtos, nil
}

package repo

import (
	"database-example/model"

	"gorm.io/gorm"
)

type PublicFacilityRequestRepository struct {
	DatabaseConnection *gorm.DB
}

func (repo *PublicFacilityRequestRepository) CreatePublicFacilityRequest(request *model.PublicFacilityRequest) error {
	dbResult := repo.DatabaseConnection.Create(request)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	println("Rows affected: ", dbResult.RowsAffected)
	return nil
}

func (repo *PublicFacilityRequestRepository) GetAllFacilityRequest() ([]model.PublicFacilityRequest, error) {
	var storedRequests []model.PublicFacilityRequest
	dbResult := repo.DatabaseConnection.Find(&storedRequests)
	if dbResult.Error != nil {
		return nil, dbResult.Error
	}

	return storedRequests, nil
}

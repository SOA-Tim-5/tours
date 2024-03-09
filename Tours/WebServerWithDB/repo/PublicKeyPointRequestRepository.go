package repo

import (
	"database-example/model"

	"gorm.io/gorm"
)

type PublicKeyPointRequestRepository struct {
	DatabaseConnection *gorm.DB
}


func (repo *PublicKeyPointRequestRepository) CreatePublicKeyPointRequest(request *model.PublicKeyPointRequest) error {
	dbResult := repo.DatabaseConnection.Create(request)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	println("Rows affected: ", dbResult.RowsAffected)
	return nil
}

func (repo *PublicKeyPointRequestRepository) GetAll() ([]model.PublicKeyPointRequest, error) {
	var storedRequests []model.PublicKeyPointRequest
	dbResult := repo.DatabaseConnection.Find(&storedRequests)
	if dbResult.Error != nil {
		return nil, dbResult.Error
	}

	return storedRequests, nil
}

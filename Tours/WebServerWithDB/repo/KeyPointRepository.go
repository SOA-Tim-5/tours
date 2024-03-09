package repo

import (
	"database-example/model"

	"gorm.io/gorm"
)

type KeyPointRepository struct {
	DatabaseConnection *gorm.DB
}

/*func (repo *StudentRepository) FindById(id string) (model.Student, error) {
	student := model.Student{}
	dbResult := repo.DatabaseConnection.First(&student, "id = ?", id)
	if dbResult != nil {
		return student, dbResult.Error
	}
	return student, nil
}*/

func (repo *KeyPointRepository) CreateKeyPoint(keyPoint *model.KeyPoint) error {
	dbResult := repo.DatabaseConnection.Create(keyPoint)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	println("Rows affected: ", dbResult.RowsAffected)
	return nil
}

func (repo *KeyPointRepository) GetKeyPoints(tourId int64) ([]model.KeyPoint, error) {
	var storedKeyPoints []model.KeyPoint
	dbResult := repo.DatabaseConnection.Find(&storedKeyPoints, "tour_id = ?", tourId)
	
	
	if dbResult.Error != nil {
		return nil, dbResult.Error
	}

	return storedKeyPoints, nil
}

func (repo *KeyPointRepository) GetById(id int64) (model.KeyPoint, error) {
	keyPoint := model.KeyPoint{}
	dbResult := repo.DatabaseConnection.First(&keyPoint, "id = ?", id)
	if dbResult != nil {
		return keyPoint, dbResult.Error
	}
	return keyPoint, nil
}
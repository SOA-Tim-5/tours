package repo

import (
	"database-example/model"

	"gorm.io/gorm"
)

type FacilityRepository struct {
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

func (repo *FacilityRepository) CreateFacility(facility *model.Facility) error {
	dbResult := repo.DatabaseConnection.Create(facility)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	println("Rows affected: ", dbResult.RowsAffected)
	return nil
}

func (repo *FacilityRepository) GetByAuthorId(authorId int64) ([]model.Facility, error) {
	var storedFacilites []model.Facility
	dbResult := repo.DatabaseConnection.Find(&storedFacilites, "author_id = ?", authorId)
	if dbResult.Error != nil {
		return nil, dbResult.Error
	}

	return storedFacilites, nil
}
func (repo *FacilityRepository) GetById(id int64) (model.Facility, error) {
	facility := model.Facility{}
	dbResult := repo.DatabaseConnection.First(&facility, "id = ?", id)
	if dbResult != nil {
		return facility, dbResult.Error
	}
	return facility, nil
}

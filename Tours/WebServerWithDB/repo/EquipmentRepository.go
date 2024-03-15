package repo

import (
	"database-example/model"

	"gorm.io/gorm"
)

type EquipmentRepository struct {
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

func (repo *EquipmentRepository) CreateEquipment(equipment *model.Equipment) error {
	dbResult := repo.DatabaseConnection.Create(equipment)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	println("Rows affected: ", dbResult.RowsAffected)
	return nil
}

func (repo *EquipmentRepository) GetAll() ([]model.Equipment, error) {
	var storedEquipments []model.Equipment
	dbResult := repo.DatabaseConnection.Find(&storedEquipments)
	if dbResult.Error != nil {
		return nil, dbResult.Error
	}

	return storedEquipments, nil
}

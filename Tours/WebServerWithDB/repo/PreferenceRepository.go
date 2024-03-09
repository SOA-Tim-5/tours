package repo

import (
	"database-example/model"

	"gorm.io/gorm"
)

type PreferenceRepository struct {
	DatabaseConnection *gorm.DB
}


func (repo *PreferenceRepository) CreatePreference(preference *model.Preference) error {
	dbResult := repo.DatabaseConnection.Create(preference)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	println("Rows affected: ", dbResult.RowsAffected)
	return nil
}

func (repo *PreferenceRepository) GetByTouristId(touristId int64) (model.Preference, error) {
	var storedPreference model.Preference
	dbResult := repo.DatabaseConnection.First(&storedPreference, "user_id = ?", touristId)
	if dbResult.Error != nil {
		return storedPreference, dbResult.Error
	}

	return storedPreference, nil
}

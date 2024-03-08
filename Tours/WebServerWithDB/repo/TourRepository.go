package repo

import (
	"database-example/model"

	"gorm.io/gorm"
)

type TourRepository struct {
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

func (repo *TourRepository) CreateTour(tour *model.Tour) error {
	dbResult := repo.DatabaseConnection.Create(tour)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	println("Rows affected: ", dbResult.RowsAffected)
	return nil
}

func (repo *TourRepository) GetByAuthorId(authorId int64) ([]model.Tour, error) {
	var storedTours []model.Tour
	dbResult := repo.DatabaseConnection.Find(&storedTours, "author_id = ?", authorId)
	if dbResult.Error != nil {
		return nil, dbResult.Error
	}

	// Skeniranje tagova za svaku turu
	for _, i := range storedTours {
		//var tags pq.StringArray
		//err := repo.DatabaseConnection.Model(&storedTours[i]).Select("tags").Where("id = ?", storedTours[i].Id).Scan(&storedTours[i].Tags).Error
		/*if err != nil {
		            return nil, err
		        }
		        //storedTours[i].Tags = tags
				var tags []string
				for _, tag := range storedTours[i].Tags {
					tags = append(tags, tag)
				}
				storedTours[i].Tags = tags
				println(tags[0])*/
		if i.KeyPoints == nil {
			i.KeyPoints = make([]model.KeyPoint, 0)
		}
	}

	return storedTours, nil
}
func (repo *TourRepository) GetById(id int64) (model.Tour, error) {
	tour := model.Tour{}
	dbResult := repo.DatabaseConnection.First(&tour, "id = ?", id)
	if dbResult != nil {
		return tour, dbResult.Error
	}
	return tour, nil
}

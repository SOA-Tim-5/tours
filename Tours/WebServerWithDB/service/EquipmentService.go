package service

import (
	"database-example/model"
	"database-example/repo"
	"fmt"
)

type EquipmentService struct {
	EquipmentRepo *repo.EquipmentRepository
}

/*func (service *TourService) FindStudent(id string) (*model.Student, error) {
	student, err := service.StudentRepo.FindById(id)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("menu item with id %s not found", id))
	}
	return &student, nil
}*/

func (service *EquipmentService) Create(equipment *model.Equipment) error {
	err := service.EquipmentRepo.CreateEquipment(equipment)
	if err != nil {
		return err
	}
	return nil
}
func (service *EquipmentService) GetAll() ([]model.Equipment, error) {
	storedEquipments, err := service.EquipmentRepo.GetAll()
	if err != nil {
		return nil, fmt.Errorf("not found")
	}
	return storedEquipments, nil
}

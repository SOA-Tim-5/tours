package handler

import (
	"database-example/model"
	"database-example/service"
	"encoding/json"
	"net/http"
)

type EquipmentHandler struct {
	EquipmentService *service.EquipmentService
}

/*
func (handler *TourHandler) Get(writer http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	log.Printf("Student sa id-em %s", id)
	// student, err := handler.StudentService.FindStudent(id)
	// writer.Header().Set("Content-Type", "application/json")
	// if err != nil {
	// 	writer.WriteHeader(http.StatusNotFound)
	// 	return
	// }
	writer.WriteHeader(http.StatusOK)
	// json.NewEncoder(writer).Encode(student)
}*/

func (handler *EquipmentHandler) Create(writer http.ResponseWriter, req *http.Request) {
	var equipment model.Equipment
	//println(tour.Tags[0], tour.Tags[1])
	err := json.NewDecoder(req.Body).Decode(&equipment)
	if err != nil {
		println("Error while parsing json")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	err = handler.EquipmentService.Create(&equipment)
	if err != nil {
		println("Error while creating a new equipment")
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")

	json.NewEncoder(writer).Encode(equipment)
}
func (handler *EquipmentHandler) GetAll(writer http.ResponseWriter, req *http.Request) {

	storedFacilites, err := handler.EquipmentService.GetAll()
	writer.Header().Set("Content-Type", "application/json")
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}

	err = json.NewEncoder(writer).Encode(storedFacilites)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
}

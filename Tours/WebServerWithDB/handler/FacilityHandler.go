package handler

import (
	"database-example/model"
	"database-example/service"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type FacilityHandler struct {
	FacilityService *service.FacilityService
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

func (handler *FacilityHandler) Create(writer http.ResponseWriter, req *http.Request) {
	var facility model.Facility
	//println(tour.Tags[0], tour.Tags[1])
	err := json.NewDecoder(req.Body).Decode(&facility)
	if err != nil {
		println("Error while parsing json")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	err = handler.FacilityService.Create(&facility)
	if err != nil {
		println("Error while creating a new facility")
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")
}

func (handler *FacilityHandler) GetByAuthorId(writer http.ResponseWriter, req *http.Request) {
	authorIdStr := mux.Vars(req)["authorId"]
	authorId, err := strconv.ParseInt(authorIdStr, 10, 64)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	storedFacilites, err := handler.FacilityService.GetByAuthorId(authorId)
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

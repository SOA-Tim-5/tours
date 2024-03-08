package handler

import (
	"database-example/model"
	"database-example/service"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type TourHandler struct {
	TourService *service.TourService
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

func (handler *TourHandler) Create(writer http.ResponseWriter, req *http.Request) {
	var tour model.Tour
	//println(tour.Tags[0], tour.Tags[1])
	err := json.NewDecoder(req.Body).Decode(&tour)
	if err != nil {
		println("Error while parsing json")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	err = handler.TourService.Create(&tour)
	if err != nil {
		println("Error while creating a new tour")
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(tour)
}

func (handler *TourHandler) GetByAuthorId(writer http.ResponseWriter, req *http.Request) {
	authorIdStr := mux.Vars(req)["authorId"]
	authorId, err := strconv.ParseInt(authorIdStr, 10, 64)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	storedTours, err := handler.TourService.GetByAuthorId(authorId)
	writer.Header().Set("Content-Type", "application/json")
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}

	err = json.NewEncoder(writer).Encode(storedTours)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (handler *TourHandler) GetById(writer http.ResponseWriter, req *http.Request) {
	tourIdStr := mux.Vars(req)["tourId"]
	tourId, err := strconv.ParseInt(tourIdStr, 10, 64)

	tour, err := handler.TourService.GetById(tourId)
	writer.Header().Set("Content-Type", "application/json")
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(tour)
}

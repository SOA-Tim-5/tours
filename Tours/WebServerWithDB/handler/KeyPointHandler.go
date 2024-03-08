package handler

import (
	"database-example/model"
	"database-example/service"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type KeyPointHandler struct {
	KeyPointService *service.KeyPointService
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

func (handler *KeyPointHandler) Create(writer http.ResponseWriter, req *http.Request) {
	var keyPoint model.KeyPoint
	//println(tour.Tags[0], tour.Tags[1])
	err := json.NewDecoder(req.Body).Decode(&keyPoint)
	if err != nil {
		println("Error while parsing json")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	err = handler.KeyPointService.Create(&keyPoint)
	if err != nil {
		println("Error while creating a new keypoint")
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(keyPoint)

}
func (handler *KeyPointHandler) GetKeyPoints(writer http.ResponseWriter, req *http.Request) {
	tourIdStr := mux.Vars(req)["tourId"]
	tourId, err := strconv.ParseInt(tourIdStr, 10, 64)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	storedKeyPoints, err := handler.KeyPointService.GetKeyPoints(tourId)
	writer.Header().Set("Content-Type", "application/json")
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}

	err = json.NewEncoder(writer).Encode(storedKeyPoints)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	println("usaoo")
}

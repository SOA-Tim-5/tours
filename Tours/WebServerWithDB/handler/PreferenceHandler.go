package handler

import (
	"database-example/model"
	"database-example/service"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type PreferenceHandler struct {
	PreferenceService *service.PreferenceService
}


func (handler *PreferenceHandler) Create(writer http.ResponseWriter, req *http.Request) {
	var preference model.Preference
	//println(tour.Tags[0], tour.Tags[1])
	err := json.NewDecoder(req.Body).Decode(&preference)
	if err != nil {
		println("Error while parsing json")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	err = handler.PreferenceService.Create(&preference)
	if err != nil {
		println("Error while creating a new preference")
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")

	json.NewEncoder(writer).Encode(preference)
}
func (handler *PreferenceHandler) GetByTouristId(writer http.ResponseWriter, req *http.Request) {

	touristIdStr := mux.Vars(req)["touristId"]
	touristId, _ := strconv.ParseInt(touristIdStr, 10, 64)
	storedPreference, err := handler.PreferenceService.GetByTouristId(touristId)
	writer.Header().Set("Content-Type", "application/json")
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}

	err = json.NewEncoder(writer).Encode(storedPreference)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
}

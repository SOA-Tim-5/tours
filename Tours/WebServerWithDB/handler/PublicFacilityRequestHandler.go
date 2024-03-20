package handler

import (
	"database-example/model"
	"database-example/service"
	"encoding/json"
	"net/http"
)

type PublicFacilityRequestHandler struct {
	PublicFacilityRequestService *service.PublicFacilityRequestService
}

func (handler *PublicFacilityRequestHandler) Create(writer http.ResponseWriter, req *http.Request) {
	var request model.PublicFacilityRequest
	//println(tour.Tags[0], tour.Tags[1])
	err := json.NewDecoder(req.Body).Decode(&request)
	if err != nil {
		println("Error while parsing json")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	err = handler.PublicFacilityRequestService.Create(&request)
	if err != nil {
		println("Error while creating a new request")
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")

	json.NewEncoder(writer).Encode(request)
}

func (handler *PublicFacilityRequestHandler) GetAllFacilityRequest(writer http.ResponseWriter, req *http.Request) {

	storedRequests, err := handler.PublicFacilityRequestService.GetAllFacilityRequest()
	writer.Header().Set("Content-Type", "application/json")
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}

	err = json.NewEncoder(writer).Encode(storedRequests)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
}

package main

import (
	"database-example/handler"
	"database-example/model"
	"database-example/repo"
	"database-example/service"
	"log"
	"net/http"

	"gorm.io/driver/postgres"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func initDB() *gorm.DB {
	dsn := "user=postgres password=super dbname=explorer host=database port=5432 sslmode=disable search_path=tours"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		print(err)
		return nil
	}

	database = database.Exec("create schema if not exists tours")
	if database.Error != nil {
		print(database.Error)
		return nil
	}

	err = database.AutoMigrate(&model.Equipment{}, &model.Facility{}, &model.KeyPointSecret{}, &model.TourDuration{}, &model.Tour{},
		&model.KeyPoint{},
		&model.PublicKeyPoint{}, &model.PublicKeyPointNotification{}, &model.PublicKeyPointRequest{},
		&model.Review{}, &model.Campaign{}, &model.Coordinate{},
		&model.Preference{}, &model.PublicFacilityNotification{}, &model.PublicFacilityRequest{}, &model.Subscriber{},
		&model.TourExecutionSession{}, &model.TouristEquipment{}, &model.TouristPosition{})
	if err != nil {
		log.Fatalf("Error migrating models: %v", err)
	}

	return database
}

func startTourServer(handler *handler.TourHandler, keyPointHandler *handler.KeyPointHandler, facilityHandler *handler.FacilityHandler,
	equipmentHandler *handler.EquipmentHandler, preferenceHandler *handler.PreferenceHandler, publicKeyPointRequestHandler *handler.PublicKeyPointRequestHandler,
	publicFacilityRequestHandler *handler.PublicFacilityRequestHandler) {
	router := mux.NewRouter().StrictSlash(true)

	//POST
	router.HandleFunc("/tour/create", handler.Create).Methods("POST")
	router.HandleFunc("/keypoint/create", keyPointHandler.Create).Methods("POST")
	router.HandleFunc("/facility/create", facilityHandler.Create).Methods("POST")
	router.HandleFunc("/equipment/create", equipmentHandler.Create).Methods("POST")
	router.HandleFunc("/preference/create", preferenceHandler.Create).Methods("POST")
	router.HandleFunc("/publicKeyPointRequest/create", publicKeyPointRequestHandler.Create).Methods("POST")
	router.HandleFunc("/publicFacilityRequest/create", publicFacilityRequestHandler.Create).Methods("POST")

	//GET
	router.HandleFunc("/tours/get/{authorId}", handler.GetByAuthorId).Methods("GET")
	router.HandleFunc("/tours/getTour/{tourId}", handler.GetById).Methods("GET")
	router.HandleFunc("/tours/getKeypoints/{tourId}", keyPointHandler.GetKeyPoints).Methods("GET")
	router.HandleFunc("/facility/get/{authorId}", facilityHandler.GetByAuthorId).Methods("GET")
	router.HandleFunc("/equipment/get/", equipmentHandler.GetAll).Methods("GET")
	router.HandleFunc("/preferences/get/{touristId}", preferenceHandler.GetByTouristId).Methods("GET")
	router.HandleFunc("/publicKeyPointRequest/get/", publicKeyPointRequestHandler.GetAll).Methods("GET")
	router.HandleFunc("/publicFacilityRequest/get/", publicFacilityRequestHandler.GetAllFacilityRequest).Methods("GET")

	println("Server starting")
	log.Fatal(http.ListenAndServe(":88", router))
}

func main() {
	database := initDB()
	if database == nil {
		print("FAILED TO CONNECT TO DB")
		return
	}

	keyPointRepo := &repo.KeyPointRepository{DatabaseConnection: database}
	keyPointService := &service.KeyPointService{KeyPointRepo: keyPointRepo}
	keyPointHandler := &handler.KeyPointHandler{KeyPointService: keyPointService}

	tourRepo := &repo.TourRepository{DatabaseConnection: database}
	tourService := &service.TourService{TourRepo: tourRepo, KeypointRepo: keyPointRepo}
	tourHandler := &handler.TourHandler{TourService: tourService}

	facilityRepo := &repo.FacilityRepository{DatabaseConnection: database}
	facilityService := &service.FacilityService{FacilityRepo: facilityRepo}
	facilityHandler := &handler.FacilityHandler{FacilityService: facilityService}

	equipmentRepo := &repo.EquipmentRepository{DatabaseConnection: database}
	equipmentService := &service.EquipmentService{EquipmentRepo: equipmentRepo}
	equipmentHandler := &handler.EquipmentHandler{EquipmentService: equipmentService}

	preferenceRepo := &repo.PreferenceRepository{DatabaseConnection: database}
	preferenceService := &service.PreferenceService{PreferenceRepo: preferenceRepo}
	preferenceHandler := &handler.PreferenceHandler{PreferenceService: preferenceService}

	publicKeyPointRequestRepo := &repo.PublicKeyPointRequestRepository{DatabaseConnection: database}
	publicKeyPointRequestService := &service.PublicKeyPointRequestService{PublicKeyPointRequestRepo: publicKeyPointRequestRepo, KeypointRepo: keyPointRepo}
	publicKeyPointRequestHandler := &handler.PublicKeyPointRequestHandler{PublicKeyPointRequestService: publicKeyPointRequestService}

	publicFacilityRequestRepo := &repo.PublicFacilityRequestRepository{DatabaseConnection: database}
	publicFacilityRequestService := &service.PublicFacilityRequestService{PublicFacilityRequestRepo: publicFacilityRequestRepo, FacilityRepo: facilityRepo}
	publicFacilityRequestHandler := &handler.PublicFacilityRequestHandler{PublicFacilityRequestService: publicFacilityRequestService}

	startTourServer(tourHandler, keyPointHandler, facilityHandler, equipmentHandler, preferenceHandler, publicKeyPointRequestHandler, publicFacilityRequestHandler)

}

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
	dsn := "user=postgres password=super dbname=explorer-v1 host=localhost port=5432 sslmode=disable search_path=tours"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		print(err)
		return nil
	}

	err = database.AutoMigrate(&model.Equipment{}, &model.Facility{}, &model.KeyPointSecret{}, &model.TourDuration{},
		&model.KeyPoint{},
		&model.PublicKeyPoint{}, &model.PublicKeyPointNotification{}, &model.PublicKeyPointRequest{},
		&model.Review{}, &model.Tour{}, &model.Campaign{}, &model.Coordinate{},
		&model.Preference{}, &model.PublicFacilityNotification{}, &model.PublicFacilityRequest{}, &model.Subscriber{},
		&model.TourExecutionSession{}, &model.TouristEquipment{}, &model.TouristPosition{})
	if err != nil {
		log.Fatalf("Error migrating models: %v", err)
	}

	return database
}

func startTourServer(handler *handler.TourHandler, keyPointHandler *handler.KeyPointHandler, facilityHandler *handler.FacilityHandler, equipmentHandler *handler.EquipmentHandler) {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/tour/create", handler.Create).Methods("POST")
	router.HandleFunc("/keypoint/create", keyPointHandler.Create).Methods("POST")
	router.HandleFunc("/facility/create", facilityHandler.Create).Methods("POST")
	router.HandleFunc("/equipment/create", equipmentHandler.Create).Methods("POST")

	println("Server starting")
	log.Fatal(http.ListenAndServe(":88", router))
}

func main() {
	database := initDB()
	if database == nil {
		print("FAILED TO CONNECT TO DB")
		return
	}
	tourRepo := &repo.TourRepository{DatabaseConnection: database}
	tourService := &service.TourService{TourRepo: tourRepo}
	tourHandler := &handler.TourHandler{TourService: tourService}

	keyPointRepo := &repo.KeyPointRepository{DatabaseConnection: database}
	keyPointService := &service.KeyPointService{KeyPointRepo: keyPointRepo}
	keyPointHandler := &handler.KeyPointHandler{KeyPointService: keyPointService}

	facilityRepo := &repo.FacilityRepository{DatabaseConnection: database}
	facilityService := &service.FacilityService{FacilityRepo: facilityRepo}
	facilityHandler := &handler.FacilityHandler{FacilityService: facilityService}

	equipmentRepo := &repo.EquipmentRepository{DatabaseConnection: database}
	equipmentService := &service.EquipmentService{EquipmentRepo: equipmentRepo}
	equipmentHandler := &handler.EquipmentHandler{EquipmentService: equipmentService}
	startTourServer(tourHandler, keyPointHandler, facilityHandler, equipmentHandler)
}

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

	//POST
	router.HandleFunc("/tour/create", handler.Create).Methods("POST")
	router.HandleFunc("/keypoint/create", keyPointHandler.Create).Methods("POST")
	router.HandleFunc("/facility/create", facilityHandler.Create).Methods("POST")
	router.HandleFunc("/equipment/create", equipmentHandler.Create).Methods("POST")

	//GET
	router.HandleFunc("/tours/get/{authorId}", handler.GetByAuthorId).Methods("GET")
	router.HandleFunc("/tours/getTour/{tourId}", handler.GetById).Methods("GET")
	router.HandleFunc("/tours/getKeypoints/{tourId}", keyPointHandler.GetKeyPoints).Methods("GET")
	router.HandleFunc("/facility/get/{authorId}", facilityHandler.GetByAuthorId).Methods("GET")

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
	startTourServer(tourHandler, keyPointHandler, facilityHandler, equipmentHandler)
}

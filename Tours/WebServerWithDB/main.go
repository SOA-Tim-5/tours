package main

import (
	"context"
	"database-example/handler"
	"database-example/model"
	"database-example/proto/tour"
	"database-example/repo"
	"database-example/service"
	"log"
	"net"
	"net/http"
	"strconv"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"gorm.io/driver/postgres"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func initDB() *gorm.DB {
	// host database dbname explorer-v1
	dsn := "user=postgres password=super dbname=explorer-v1 host=localhost port=5432 sslmode=disable search_path=tours"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		print(err)
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

	keyPointRepo := repo.KeyPointRepository{DatabaseConnection: database}
	/*
		tourRepo := &repo.TourRepository{DatabaseConnection: database}
		facilityRepo := &repo.FacilityRepository{DatabaseConnection: database}
		equipmentRepo := &repo.EquipmentRepository{DatabaseConnection: database}
		preferenceRepo := &repo.PreferenceRepository{DatabaseConnection: database}
		publicFacilityRequestRepo := &repo.PublicFacilityRequestRepository{DatabaseConnection: database}
		publicKeyPointRequestRepo := &repo.PublicKeyPointRequestRepository{DatabaseConnection: database}


		type Server struct {
			tour.UnimplementedEncounterServer
			EncounterRepo         *repo.EncounterRepository
			EncounterInstanceRepo *repo.EncounterInstanceRepository
			TouristProgressRepo   *repo.TouristProgressRepository
		}

	*/
	//keyPointService := &service.KeyPointService{KeyPointRepo: keyPointRepo}
	//keyPointHandler := &handler.KeyPointHandler{KeyPointService: keyPointService}

	tourRepo := repo.TourRepository{DatabaseConnection: database}
	//tourService := &service.TourService{TourRepo: tourRepo, KeypointRepo: keyPointRepo}
	/*
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
	*/
	//	startTourServer(tourHandler, keyPointHandler, facilityHandler, equipmentHandler, preferenceHandler, publicKeyPointRequestHandler, publicFacilityRequestHandler)

	lis, err := net.Listen("tcp", "localhost:88")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)

	tour.RegisterTourServiceServer(grpcServer, Server{TourRepo: &tourRepo, KeyPointRepo: &keyPointRepo})
	reflection.Register(grpcServer)
	grpcServer.Serve(lis)

}

type Server struct {
	tour.UnimplementedTourServiceServer
	TourRepo     *repo.TourRepository
	KeyPointRepo *repo.KeyPointRepository
}

func (s Server) Create(ctx context.Context, request *tour.TourCreateDto) (*tour.TourResponseDto, error) {
	t := model.Tour{
		AuthorId: request.AuthorId, Name: request.Name, Description: request.Description, Difficulty: int(request.Difficulty),
		Tags: request.Tags, Status: model.TourStatus(request.Status), Price: request.Price, IsDeleted: request.IsDeleted,
		Distance: request.Distance, Category: model.TourCategory(request.Category),
	}

	tourService := service.TourService{TourRepo: s.TourRepo, KeypointRepo: s.KeyPointRepo}
	err := tourService.Create(&t)
	if err != nil {
		println("Error while creating a new tour")
		return nil, nil
	}

	println(t.Name)
	return &tour.TourResponseDto{
		Id: 0, AuthorId: &request.AuthorId, Name: request.Name, Description: request.Description, Difficulty: request.Difficulty,
		Tags: request.Tags, Status: tour.TourResponseDto_TourStatus(request.Status), Price: request.Price, IsDeleted: request.IsDeleted,
		Distance: request.Distance, Category: tour.TourResponseDto_TourCategory(request.Category),
	}, nil
}

func (s Server) GetAuthorsTours(ctx context.Context, request *tour.GetParams) (*tour.TourListResponse, error) {

	tourService := service.TourService{TourRepo: s.TourRepo, KeypointRepo: s.KeyPointRepo}
	authorId, err := strconv.ParseInt(request.Id, 10, 64)
	if err != nil {
		return nil, nil
	}
	tours, err := tourService.GetByAuthorId(authorId)
	if err != nil {
		println("Error while getting")
		return nil, nil
	}

	println(request.Id)

	tourResponses := []*tour.TourResponseDto{}
	for i := 0; i < len(tours); i++ {
		tourResponses = append(tourResponses,
			&tour.TourResponseDto{
				Id: tours[i].Id, AuthorId: &tours[i].AuthorId, Name: tours[i].Name, Description: tours[i].Description, Difficulty: int32(tours[i].Difficulty),
				Tags: tours[i].Tags, Status: tour.TourResponseDto_TourStatus(tours[i].Status), Price: tours[i].Price, IsDeleted: tours[i].IsDeleted,
				Distance: tours[i].Distance, Category: tour.TourResponseDto_TourCategory(tours[i].Category),
			},
		)
	}

	return &tour.TourListResponse{
		TourResponses: tourResponses,
	}, nil
}

func (s Server) CreateKeyPoint(ctx context.Context, request *tour.KeyPointCreateDto) (*tour.KeyPointResponseDto, error) {

	keyPointService := &service.KeyPointService{s.KeyPointRepo}
	tourId, err := strconv.ParseInt(request.TourId, 10, 64)
	if err != nil {
		return nil, nil
	}
	keyPoint := model.KeyPoint{
		TourId: tourId, Name: request.Name, Description: request.Description, Longitude: request.Longitude,
		Latitude: request.Latitude, LocationAddress: request.LocationAddress, ImagePath: request.ImagePath,
		Order: request.Order, IsEncounterRequired: request.IsEncounterRequired, HasEncounter: request.HasEncounter,
	}
	err = keyPointService.Create(&keyPoint)

	println(request.TourId)

	return &tour.KeyPointResponseDto{
		TourId: tourId, Name: request.Name, Description: request.Description, Longitude: request.Longitude,
		Latitude: request.Latitude, LocationAddress: request.LocationAddress, ImagePath: request.ImagePath,
		Order: request.Order,
	}, nil
}

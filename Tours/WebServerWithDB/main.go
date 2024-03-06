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

func initDB() gorm.DB {
    dsn := "user=postgres password=super dbname=explorer-v1 host=localhost port=5432 sslmode=disable search_path=tours"
    database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

    if err != nil {
        print(err)
        return nil
    }

    err = database.AutoMigrate(&model.Equipment{}, &model.Facility{}, &model.KeyPoint{},
        &model.PublicKeyPoint{}, &model.PublicKeyPointNotification{}, &model.PublicKeyPointRequest{},
		&model.Review{}, &model.Tour{})
    if err != nil {
        log.Fatalf("Error migrating models: %v", err)
    }

    return database
}

func startTourServer(handlerhandler.TourHandler) {
    router := mux.NewRouter().StrictSlash(true)

    router.HandleFunc("/tour/misc", handler.CreateMiscEncounter).Methods("POST")

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
    startTourServer(tourHandler)
}

package main

import (
	"database-example/handler"
	"database-example/model"
	"database-example/repo"
	"database-example/service"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func initDB() *gorm.DB {
	connectionStr := "user=postgres password=super dbname=explorer-v1 host=localhost port=5432 sslmode=disable search_path=tours"

	database, err := gorm.Open(mysql.Open(connectionStr), &gorm.Config{})
	if err != nil {
		print(err)
		return nil
	}

	database.AutoMigrate(&model.Student{})
	database.Exec("INSERT IGNORE INTO students VALUES ('aec7e123-233d-4a09-a289-75308ea5b7e6', 'Marko Markovic', 'Graficki dizajn')")
	return database
}

func startServer(handler *handler.StudentHandler) {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/students/{id}", handler.Get).Methods("GET")
	router.HandleFunc("/students", handler.Create).Methods("POST")

	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static")))
	println("Server starting")
	log.Fatal(http.ListenAndServe(":8088", router))
}

func main() {
	database := initDB()
	if database == nil {
		print("FAILED TO CONNECT TO DB")
		return
	}
	repo := &repo.StudentRepository{DatabaseConnection: database}
	service := &service.StudentService{StudentRepo: repo}
	handler := &handler.StudentHandler{StudentService: service}

	startServer(handler)
}
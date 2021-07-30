package main

import (
	"net/http"
	"os"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

func initAPIHandler() (Handler, error) {
	address := os.Getenv("REDIS_URL")
	if address == "" {
		log.Info("Failed to read REDIS_URL from .env file, using a default value instead")
		address = "redis:6379" //todo Czy pozwalać programowi kontynuować? Czy lepiej przerwać?
	}

	database := NewDatabaseConnection(address,os.Getenv("REDIS_PASSWORD"))
	err := database.Connect()
	if err != nil {
		return Handler{}, err
	}
	apiHandler := NewHandler(database)
	return apiHandler, nil
}

func main() {
	router := mux.NewRouter()

	handler, err := initAPIHandler()
	if err != nil {
		log.Fatalf("Error connecting to database: %s", err)
	}

	router.HandleFunc("/v1/images", handler.DBGetAllHandler).Methods("GET")
	router.HandleFunc("/v1/images/{id}", handler.DBGetHandler).Methods("GET")
	router.HandleFunc("/v1/images/{id}", handler.DBPostHandler).Methods("POST")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
		log.Info("Failed to read PORT from .env file, using a default value instead")
		//todo return?? To samo co z REDIS_URL
	}

	log.Info("Starting server at port %s\n", port)
	err = http.ListenAndServe(":"+port, router)
	if err != nil {
		log.Fatalf("Starting server at port %s failed!", port)
	}
}

package main

import (
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
)

// initAPIHandler initializes a handler for the API.
func initAPIHandler() (Handler, error) {
	address := os.Getenv("REDIS_URL")
	if address == "" {
		log.Fatal("Failed to read REDIS_URL from .env file")
	}

	database := NewDatabaseConnection(address, os.Getenv("REDIS_PASSWORD"))
	err := database.Connect()
	if err != nil {
		return Handler{}, err
	}
	apiHandler := NewHandler(database)
	return apiHandler, nil
}

// main contains all the function handlers and initializes the database connection.
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
		log.Fatal("Failed to read PORT from .env file")
	}

	log.Info("Starting server at port " + port + "\n")
	err = http.ListenAndServe(":"+port, router)
	if err != nil {
		log.Fatalf("Starting server at port %s failed!", port)
	}
	log.Printf("Starting server at port %s\n", port)
	err = http.ListenAndServe(":"+port, router)
}

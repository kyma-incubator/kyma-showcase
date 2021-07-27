package main

import (
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
)

func initAPIHandler() (Handler, error) {
	address := os.Getenv("REDIS_URL")
	if address == "" {
		address = "redis:6379"
	}

	database := NewDatabaseConnection(address,os.Getenv("REDIS_PASSWORD"))
	err := database.Connect()
	if err != nil {
		log.Println(err)
		return Handler{}, err
	}

	apiHandler := NewHandler(database)
	return apiHandler, nil
}

func main() {
	router := mux.NewRouter()

	handler, err := initAPIHandler()
	if err != nil {
		os.Exit(1)
	}

	///v1/health --> dla k8s
	// sciezka /wersjonowane[v1]/rejestracja-metadanych[metadata]/services

	//zaprojektowanie API na sucho przed implementajca

	router.HandleFunc("/v1/photo/{id}", handler.DBGetHandler).Methods("GET") // get certain image
	router.HandleFunc("/v1/photo", handler.DBGetAllHandler).Methods("GET") //get all images
	router.HandleFunc("/v1/photo/{id}", handler.DBPostHandler).Methods("POST") // add photo to database

	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}
	log.Printf("Starting server at port %s\n", port)
	err = http.ListenAndServe(":"+port, router)
}

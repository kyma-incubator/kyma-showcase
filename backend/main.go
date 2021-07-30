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
		//router.HandleFunc("/",ErrorHandler) //500
	}

	///v1/health --> dla k8s
	// sciezka /wersjonowane[v1]/rejestracja-metadanych[metadata]/services

	//zaprojektowanie API na sucho przed implementajca

	router.HandleFunc("/get/{id}", handler.DBGetHandler).Methods("GET")
	router.HandleFunc("/post/{id}", handler.DBPostHandler).Methods("POST")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}
	log.Printf("Starting server at port %s\n", port)
	err = http.ListenAndServe(":"+port, router)
}

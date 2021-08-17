package main

import (
	"github.com/gorilla/mux"
	"github.com/kyma-incubator/Kyma-Showcase/internal/api"
	db "github.com/kyma-incubator/Kyma-Showcase/internal/database"
	"github.com/kyma-incubator/Kyma-Showcase/internal/utils"
	"github.com/rs/cors"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
)

// initAPIHandler initializes a handler for the API.
func initAPIHandler() (api.Handler, error) {
	address := os.Getenv("REDIS_URL")
	if address == "" {
		log.Fatal("Failed to read REDIS_URL from .env file")
	}

	database := db.NewDatabaseConnection(address, os.Getenv("REDIS_PASSWORD"))
	err := database.Connect()
	if err != nil {
		return api.Handler{}, err
	}
	apiHandler := api.NewHandler(database, utils.NewIdGenerator())
	return apiHandler, nil
}

// main contains all the function handlers and initializes the database connection.
func main() {
	mux := mux.NewRouter()

	APIhandler, err := initAPIHandler()
	if err != nil {
		log.Fatalf("Error connecting to database: %s", err)
	}

	APIhandler.EndpointInitialize(mux)

	handler := cors.Default().Handler(mux)

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("Failed to read PORT from .env file")
	}

	err = http.ListenAndServe(":"+port, handler)
	if err != nil {
		log.Fatalf("Starting server at port %s failed!", port)
	}
}

package main

import (
	"github.com/gorilla/mux"
	"github.com/kyma-incubator/Kyma-Showcase/internal/api"
	db "github.com/kyma-incubator/Kyma-Showcase/internal/database"
	"github.com/kyma-incubator/Kyma-Showcase/internal/events"
	"github.com/rs/cors"
	log "github.com/sirupsen/logrus"
	"github.com/vrischmann/envconfig"
	"net/http"
)

// Configuration struct containing environmental variables
type Configuration struct {
	Redis struct {
		URL      string
		Password string
	}
	Backend struct {
		Port string `envconfig:"default=8081"`
	}
	Event struct {
		URL string
	}
}

// initEnvConfiguration initialize environmental variables
func initEnvConfiguration() (Configuration, error) {
	configuration := Configuration{}
	if err := envconfig.Init(&configuration); err != nil {
		return Configuration{}, err
	}

	return configuration, nil
}

// initAPIHandler initializes a handler for the API.
func initAPIHandler(conf Configuration) (api.Handler, error) {

	database := db.NewDatabaseConnection(conf.Redis.URL, conf.Redis.Password)
	err := database.Connect()
	if err != nil {
		return api.Handler{}, err
	}

	eventHandler := events.NewEventHandler(conf.Event.URL)

	apiHandler := api.NewHandler(database, api.NewIdGenerator(), eventHandler)
	return apiHandler, nil
}

// main contains all the function handlers and initializes the database connection.
func main() {
	conf, err := initEnvConfiguration()
	if err != nil {
		log.Fatal("Error when getting environmental variables: " + err.Error())
	}

	router := mux.NewRouter()

	apiHandler, err := initAPIHandler(conf)
	if err != nil {
		log.Fatalf("Error connecting to database: %s", err)
	}
	apiHandler.EndpointInitialize(router)

	handler := cors.Default().Handler(router)
	port := conf.Backend.Port
	err = http.ListenAndServe(":"+port, handler)
	if err != nil {
		log.Fatalf("Starting server at port %s failed!", port)
	}
}

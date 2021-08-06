package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
)

//go:generate mockery --name=DBManager
// DBManager defines a contract between api and the database.
type DBManager interface {
	InsertToDB(key string, value string) error
	GetFromDB(key string) (interface{}, error)
	GetAllKeys() ([]string, error)
}

//go:generate mockery --name=IdGenerator
//TODO documentation idGenerator
type IdGenerator interface {
	NewID() (string, error)
}

// Handler for database manager.
type Handler struct {
	dbManager   DBManager
	idGenerator IdGenerator
}

// NewHandler returns handler for database manager.
func NewHandler(dbManager DBManager, idGenerator IdGenerator) Handler {
	return Handler{
		dbManager:   dbManager,
		idGenerator: idGenerator,
	}
}

func accessControl(w http.ResponseWriter, r *http.Request) {
	if origin := r.Header.Get("Origin"); origin != "" {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*") //todo ip frontend instead of '*', k8s?
	}
}

// DBGetHandler processes a request and passes request ID to the GetFromDB function, returns the value of the given ID.
func (h Handler) DBGetHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/v1/images/{id}" {
		err := errors.New("DBGETHANDLER: 404 not found")
		log.Error(err)
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	accessControl(w, r)
	params := mux.Vars(r)
	var img Image

	key := params["id"]
	fromDB, err := h.dbManager.GetFromDB(key)
	if err != nil {
		err = errors.New("DBGETHANDLER: failed to get data from db: " + err.Error())
		log.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.Unmarshal([]byte(fromDB.(string)), &img)
	if err != nil {
		err = errors.New("DBGETHANDLER: failed to convert marshal to json: " + err.Error())
		log.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s", fromDB)
}

// DBGetAllHandler processes a request and gets all keys using GetAllKeys function, returns all values from database as a string with JSON array.
func (h Handler) DBGetAllHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/v1/images" {
		err := errors.New("DBGETALLHANDLER: 404 not found")
		log.Error(err)
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	accessControl(w, r)
	keys, err := h.dbManager.GetAllKeys()
	if err != nil {
		err = errors.New("DBGETALL: failed to get all keys from db: " + err.Error())
		log.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var img Image
	var result []Image
	for _, key := range keys {
		fromDB, err := h.dbManager.GetFromDB(key)
		if err != nil {
			err = errors.New("DBGETALL: failed to get value from db " + err.Error())
			log.Error(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		dataStr, ok := fromDB.(string)

		if !ok {
			err = errors.New("DBGETALL: failed to assert a type")
			log.Error(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err = json.Unmarshal([]byte(dataStr), &img)

		if err != nil {
			err = errors.New("DBGETALL: fail to unmarshal " + err.Error())
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		result = append(result, img)
	}
	allImages, err := json.Marshal(result)

	if err != nil {
		err = errors.New("DBGETALL: fail to marshal" + err.Error())
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, string(allImages))
}

// DBPostHandler processes a request and passes the parsed data to the InsertToDB function.
func (h Handler) DBPostHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/v1/images" {
		err := errors.New("POST: 404 not found")
		log.Error(err)
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	accessControl(w, r)

	var img Image

	headerContentType := r.Header.Get("Content-Type")
	if headerContentType != "application/json" {
		err := errors.New("POST: invalid content type")
		log.Error(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&img)
	if err != nil {
		err = errors.New("POST: invalid input: " + err.Error())
		log.Error(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	jsonImg, err := json.Marshal(img)
	if err != nil {
		err = errors.New("POST: failed to convert json into marshal: " + err.Error())
		log.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	id, err := h.idGenerator.NewID()
	if err != nil {
		log.Error(err)
		return
	}

	err = h.dbManager.InsertToDB(id, string(jsonImg))
	if err != nil {
		err = errors.New("POST: failed to insert values to database: " + err.Error())
		log.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

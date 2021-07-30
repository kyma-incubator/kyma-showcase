package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type DBManager interface {
	InsertToDB(key string, value string) error
	GetFromDB(key string) (interface{}, error)
	GetAllKeys() ([]string, error)
}

type Handler struct {
	dbManager DBManager
}

func NewHandler(dbManager DBManager) Handler {
	return Handler{
		dbManager: dbManager,
	}
}

func (h Handler) DBPostHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var img Image

	err := json.NewDecoder(r.Body).Decode(&img)
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
	err = h.dbManager.InsertToDB(params["id"], string(jsonImg))
	if err != nil {
		err = errors.New("POST: failed to insert values to database: " + err.Error())
		log.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h Handler) DBGetHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var img Image

	key := params["id"]
	fromDB, err := h.dbManager.GetFromDB(key)
	if err != nil {
		err = errors.New("DBGETHANDLER: failed to get data from db: " + err.Error()) //Nazwa do zmiany
		log.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.Unmarshal([]byte(fromDB.(string)), &img)
	if err != nil {
		err = errors.New("DBGETHANDLER: failed to convert marshal to json: " + err.Error()) //Nazwa do zmiany
		log.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "key: %s, value: %s\n", key, fromDB)
	fmt.Fprintf(w, "key: %s, json URL = %s,  json GCP = %s, json image = %s\n", key, img.URL, img.GCP, img.IMG)

}

func (h Handler) DBGetAllHandler(w http.ResponseWriter, r *http.Request){

	keys, err := h.dbManager.GetAllKeys()
	if err != nil {
		err = errors.New("DBGETALLHANDLER: failed to get all keys from db: " + err.Error())
		log.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	for _, key := range keys {
		fromDB, err := h.dbManager.GetFromDB(key)
		if err != nil {
			err = errors.New("DBGETALL: failed to get value from db " + err.Error())
			log.Error(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		} else {
			fmt.Fprintf(w, "key: %s, value: %s\n", key, fromDB)
		}
	}
	w.WriteHeader(http.StatusOK)
}

package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"strings"

	//"encoding/json"
	"fmt"
	//_ "image/jpeg"
	//_ "image/png"
	"net/http"
	//"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
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

	if r.Method != http.MethodPost {
		w.WriteHeader(405)
		return
	}

	params := mux.Vars(r)

	if params["id"] == "GetAllKeys" {
		log.Println("method POST is not supported for getting all keys!")
	} else {
		var img Image

		err := json.NewDecoder(r.Body).Decode(&img)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		jsonImg, err := json.Marshal(img)
		if err != nil {
			log.Println(err)
		}
		h.dbManager.InsertToDB(params["id"], string(jsonImg))
	}
}

func (h Handler) DBGetHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(405)
		return
	}

	params := mux.Vars(r)
	var img Image

	if params["id"] == "GetAllKeys" {
		keys, err := h.dbManager.GetAllKeys()
		if err != nil {
			log.Println(err)
		}
		for _, key := range keys {
			fromDB, err := h.dbManager.GetFromDB(key)
			if err != nil {
				log.Println(err)
			} else {
				fmt.Fprintf(w, "key: %s, value: %s\n", key, fromDB)
			}
		}
	} else {
		key := params["id"]
		fromDB, err := h.dbManager.GetFromDB(key)
		if err != nil {
			log.Println(err)
		}

		err = json.Unmarshal([]byte(fromDB.(string)), &img)
		if err != nil {
			log.Println(err)
		}

		fmt.Fprintf(w, "key: %s, value: %s\n", key, fromDB)
		fmt.Fprintf(w, "key: %s, json URL = %s,  json GCP = %s, json image = %s\n", key,  img.URL, img.GCP, img.IMG)
	}
}

func (h Handler) DBGetAllHandler (w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodGet {
		w.WriteHeader(405)
		return
	}

	result := "["

	keys, err := h.dbManager.GetAllKeys()
	if err != nil{
		log.Println(err)
	}

	for _,key := range keys{
		fromDB, err := h.dbManager.GetFromDB(key)
		if err != nil{
			log.Println(err)
		}
		result += fromDB.(string)
		result+=","
	}
	result = strings.TrimSuffix(result,",")
	result+="]"
	/*err = json.NewDecoder(r.Body).Decode(&result)
	if err != nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resultJSON, err := json.Marshal(result)
	fmt.Fprintf(w,string(resultJSON))*/
	fmt.Fprintf(w,result)
}

package main

import (
	//"encoding/json"
	"fmt"
	//_ "image/jpeg"
	//_ "image/png"
	"net/http"
	//"github.com/gorilla/mux"
)

// type Image struct {
// 	URL string `json:"url"`
// 	GCP string `json:"GCP"`
// 	Img string `json:"Img"`
// }

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

	h.dbManager.InsertToDB("key", "value")
	// //var img Image

	// err := json.NewDecoder(r.Body).Decode(&img)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusBadRequest)
	// 	return
	// }

	// params := mux.Vars(r)

	// if /*r.Method == "POST" */ params["id"] == "999" {
	// 	fmt.Println("json =method POST is not supported for all keys!")
	// } else {
	// 	j, err := json.Marshal(img) //key: url | value: {url:"...", GCP:"..."}
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}

	// 	h.dbManager.InsertToDB(params["id"], string(j))
	// }

}

func (h Handler) DBGetHandler(w http.ResponseWriter, r *http.Request) {

	fromDB, err := h.dbManager.GetFromDB("key")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Fprintf(w, "json = %s\n", fromDB)
	// params := mux.Vars(r)
	// var img Image

	// if params["id"] == "999" {
	// 	keys, err := h.dbManager.GetAllKeys()
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}
	// 	for _, key := range keys {
	// 		fromDB, err := h.dbManager.GetFromDB(key)
	// 		if err != nil {
	// 			fmt.Println(err)
	// 		} else {
	// 			fmt.Fprintf(w, "json = %s\n", fromDB)
	// 		}
	// 	}
	// }

	// fromDB, err := h.dbManager.GetFromDB(params["id"])
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// err = json.Unmarshal([]byte(fromDB.(string)), &img)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// fmt.Fprintf(w, "json = %s\n", fromDB)
	// fmt.Fprintf(w, "json URL = %s,  json GCP = %s, json image = %s\n", img.URL, img.GCP, img.Img)

}

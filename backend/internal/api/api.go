package api

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/kyma-incubator/Kyma-Showcase/internal/model"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
	"strings"
)

//go:generate mockery --name=DBManager --output=../mocks
// DBManager defines a contract between api and the database.
type DBManager interface {
	InsertToDB(key string, value string) error
	GetFromDB(key string) (interface{}, error)
	GetAllKeys() ([]string, error)
}

// Handler for database manager.
type Handler struct {
	dbManager      DBManager
	idGenerator    IdGenerator
	eventBus       EventBus
	getEndpoint    string
	getAllEndpoint string
	postEndpoint   string
}

// EndpointInitialize adds api endpoints to the mux router
func (h Handler) EndpointInitialize(mux *mux.Router) {
	mux.HandleFunc(h.getAllEndpoint, h.DBGetAllHandler).Methods("GET")
	mux.HandleFunc(h.getEndpoint, h.DBGetHandler).Methods("GET")
	mux.HandleFunc(h.postEndpoint, h.DBPostHandler).Methods("POST")

}

//go:generate mockery --name=EventBus --output=../mocks
// EventBus defines a contract between api and events.
type EventBus interface {
	SendNewImage(id string, img model.Image) error
}

// NewHandler returns handler for database manager.
func NewHandler(dbManager DBManager, idGenerator IdGenerator, eventBus EventBus) Handler {
	return Handler{
		dbManager:      dbManager,
		idGenerator:    idGenerator,
		eventBus:       eventBus,
		getEndpoint:    "/v1/images/{id}",
		getAllEndpoint: "/v1/images",
		postEndpoint:   "/v1/images",
	}
}

// accessControl sets headers that allow browser to pass data to frontend
func accessControl(w http.ResponseWriter, r *http.Request) {
	if origin := r.Header.Get("Origin"); origin != "" {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*") //todo ip frontend instead of '*', k8s?
	}
}

// DBGetHandler processes a request and passes request ID to the GetFromDB function, returns the value of the given ID.
func (h Handler) DBGetHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	key := params["id"]
	url := strings.Replace(h.getEndpoint, "{id}", key, 1)
	if r.URL.Path != url {
		log.Error(h.getEndpoint)
		err := errors.New("DBGETHANDLER: 404 not found")
		log.Error(err)
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	accessControl(w, r)
	var img model.Image

	fromDB, err := h.dbManager.GetFromDB(key)

	if err != nil {
		if err.Error() == "GETFROMDB:key "+key+" does not exist" {
			err = errors.New("DBGETHANDLER: failed to get data from db: " + err.Error())
			http.Error(w, err.Error(), http.StatusNotFound)
		} else {
			err = errors.New("DBGETHANDLER: failed to get data from db: " + err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		log.Error(err)
		return
	}

	err = json.Unmarshal([]byte(fromDB.(string)), &img)
	if err != nil {
		err = errors.New("DBGETHANDLER: failed to convert marshal to json: " + err.Error())
		log.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "%s", fromDB)
}

// DBGetAllHandler processes a request and gets all keys using GetAllKeys function, returns all values from database as a string with JSON array.
func (h Handler) DBGetAllHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != h.getAllEndpoint {
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

	var img model.Image
	var result []model.Image
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
	fmt.Fprint(w, string(allImages))
}

// DBPostHandler processes a request and passes the parsed data to the InsertToDB function.
func (h Handler) DBPostHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != h.postEndpoint {
		err := errors.New("POST: 404 not found")
		log.Error(err)
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	accessControl(w, r)
	var img model.Image

	headerContentType := r.Header.Get("Content-Type")
	if headerContentType != "application/json" {
		err := errors.New("POST: invalid content type")
		log.Error(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	decoder := json.NewDecoder(r.Body)
	for {
		if err := decoder.Decode(&img); err == io.EOF {
			break
		} else if err != nil {
			err = errors.New("POST: invalid input: " + err.Error())
			log.Error(err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}

	id, err := h.idGenerator.NewID()
	if err != nil {
		log.Error(err)
		return
	}

	img.ID = id
	jsonImg, err := json.Marshal(img)
	if err != nil {
		err = errors.New("POST: failed to convert json into marshal: " + err.Error())
		log.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = h.dbManager.InsertToDB(id, string(jsonImg))
	if err != nil {
		err = errors.New("POST: failed to insert values to database: " + err.Error())
		log.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var idStruct model.ID
	idStruct.ID = id

	jsonID, err := json.Marshal(idStruct)
	if err != nil {
		err = errors.New("POST: failed to convert json into marshal: " + err.Error())
		log.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	eventId, err := h.idGenerator.NewID()
	if err != nil {
		err = errors.New("POST: failed to generate id for event: " + err.Error())
		log.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = h.eventBus.SendNewImage(eventId, img)
	if err != nil {
		err = errors.New("POST: failed to send an event" + err.Error())
		log.Error(err)
		http.Error(w, err.Error(), http.StatusBadGateway)
		return
	}

	fmt.Fprint(w, string(jsonID))

}

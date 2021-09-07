package api

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"time"

	"github.com/gabriel-vasile/mimetype"
	"github.com/gorilla/mux"
	"github.com/kyma-incubator/Kyma-Showcase/internal/model"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

//go:generate mockery --name=DBManager
// DBManager defines a contract between api and the database.
type DBManager interface {
	Insert(key string, value string) error
	Get(key string) (interface{}, error)
	GetAll() ([]string, error)
}

// Handler for database manager.
type Handler struct {
	dbManager      DBManager
	idGenerator    IdGenerator
	eventBus       EventBus
	getEndpoint    string
	getAllEndpoint string
	postEndpoint   string
	putEndpoint    string
}

// EndpointInitialize adds api endpoints to the mux router
func (h Handler) EndpointInitialize(mux *mux.Router) {
	mux.HandleFunc(h.getAllEndpoint, h.GetAll).Methods("GET")
	mux.HandleFunc(h.getEndpoint, h.Get).Methods("GET")
	mux.HandleFunc(h.postEndpoint, h.Create).Methods("POST")
	mux.HandleFunc(h.putEndpoint, h.Update).Methods("PUT")
}

//go:generate mockery --name=EventBus
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
		putEndpoint:    "/v1/images/{id}",
	}
}

// accessControl sets headers that allow browser to pass data to frontend
func accessControl(w http.ResponseWriter, r *http.Request) {
	if origin := r.Header.Get("Origin"); origin != "" {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
	}
}

////requestContent check if request body contains image saved in base64
//func requestContent(w http.ResponseWriter, r *http.Request) error {
//	bodyContent, err := ioutil.ReadAll(r.Body)
//
//	if err != nil {
//		err := errors.New("requestContent: failed to read body content")
//		log.Error(err)
//		return err
//	}
//	_, err = base64.StdEncoding.DecodeString(string(bodyContent))
//
//	if err != nil {
//		err := errors.New("requestContent: body content is not an image")
//		log.Error(err)
//		return err
//	}
//	return nil
//}

// Get processes a request and passes request ID to the GetFromDB function, returns the value of the given ID.
func (h Handler) Get(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	key := params["id"]
	url := strings.Replace(h.getEndpoint, "{id}", key, 1)
	if r.URL.Path != url {
		log.Error(h.getEndpoint)
		err := errors.New("GET handler: 404 not found")
		log.Error(err)
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	accessControl(w, r)
	var img model.Image

	fromDB, err := h.dbManager.Get(key)

	if err != nil {
		if err.Error() == "GET from db: key "+key+" does not exist" {
			err = errors.New("GET handler: failed to get data from db: " + err.Error())
			http.Error(w, err.Error(), http.StatusNotFound)
		} else {
			err = errors.New("GET handler: failed to get data from db: " + err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		log.Error(err)
		return
	}

	err = json.Unmarshal([]byte(fromDB.(string)), &img)
	if err != nil {
		err = errors.New("GET handler: failed to convert marshal to json: " + err.Error())
		log.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "%s", fromDB)
}

//GetAll processes a request and gets all keys using GetAllKeys function, returns all values from database as a string with JSON array.
func (h Handler) GetAll(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != h.getAllEndpoint {
		err := errors.New("GETALL handler: 404 not found")
		log.Error(err)
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	accessControl(w, r)

	keys, err := h.dbManager.GetAll()
	if err != nil {
		err = errors.New("GETALL handler: failed to get all keys from db: " + err.Error())
		log.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var img model.Image
	var result []model.Image
	for _, key := range keys {
		fromDB, err := h.dbManager.Get(key)
		if err != nil {
			err = errors.New("GETALL handler: failed to get value from db " + err.Error())
			log.Error(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		dataStr, ok := fromDB.(string)

		if !ok {
			err = errors.New("GETALL handler: failed to assert a type")
			log.Error(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err = json.Unmarshal([]byte(dataStr), &img)

		if err != nil {
			err = errors.New("GETALL handler: fail to unmarshal " + err.Error())
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		result = append(result, img)
	}
	allImages, err := json.Marshal(result)

	if err != nil {
		err = errors.New("GETALL handler: fail to marshal" + err.Error())
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprint(w, string(allImages))
}

// getExtension returns extension of given image in url
func getExtension(bytes []byte) string {
	mType := mimetype.Detect(bytes)
	if mType.String() == "image/png" || mType.String() == "image/jpg" ||
		mType.String() == "image/jpeg" || mType.String() == "image/gif" {
		return mType.String()
	} else {
		return ""
	}
}

// calculateSize returns size of image given in base64, calculated based on equation:
// ((3/4) * length of base64 string) - number of equal signs at the end
func calculateSize(imgBase64 string) int {
	if imgBase64 == "" {
		return 0
	} else {
		l := len(imgBase64)
		eq := "="
		e := strings.Count(imgBase64[len(imgBase64)-2:], eq)
		return (3 * l / 4) - e
	}
}

// Create processes a request and passes the parsed data to the InsertToDB function.
func (h Handler) Create(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != h.postEndpoint {
		err := errors.New("CREATE handler: 404 not found")
		log.Error(err)
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	accessControl(w, r)
	var img model.Image

	headerContentType := r.Header.Get("Content-Type")
	if headerContentType != "application/json" {
		err := errors.New("CREATE handler: invalid content type")
		log.Error(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	decoder := json.NewDecoder(r.Body)
	for {
		if err := decoder.Decode(&img); err == io.EOF {
			break
		} else if err != nil {
			err = errors.New("CREATE handler: invalid input: " + err.Error())
			log.Error(err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}

	_, err := url.ParseRequestURI(img.Content)
	u, err := url.Parse(img.Content)
	if err == nil && u.Scheme != "" && u.Host != "" {
		resp, err := http.Get(img.Content)
		if err != nil {
			err = errors.New("CREATE handler: could not get image from: " + img.Content + err.Error())
			log.Error(err)
			http.Error(w, err.Error(), http.StatusNotAcceptable)
			return
		}
		defer resp.Body.Close()

		imgByte, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			err = errors.New("CREATE handler: failed to read request body" + err.Error())
			log.Error(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		imgBase64 := base64.StdEncoding.EncodeToString(imgByte)

		if calculateSize(imgBase64) > 5000000 {
			err = errors.New("image from url is too large")
			log.Error(err)
			http.Error(w, err.Error(), http.StatusTeapot)
			return
		}

		ext := getExtension(imgByte)
		if ext == "" {
			err = errors.New("extension is not supported")
			log.Error(err)
			http.Error(w, err.Error(), http.StatusUnsupportedMediaType)
			return
		}

		img.Content = fmt.Sprintf("data:%s;base64,", ext)
		img.Content += imgBase64
	} else {
		r := regexp.MustCompile("data:.*?base64,")
		contentBase64 := r.ReplaceAllString(img.Content, "")
		_, err = base64.StdEncoding.DecodeString(contentBase64)

		if err != nil {
			err = errors.New("CREATE handler: content is not an image" + err.Error())
			log.Error(err)
			http.Error(w, err.Error(), http.StatusNotAcceptable)
			return
		}
	}

	if img.Content == "" {
		err = errors.New("CREATE handler: content is empty")
		log.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	id, err := h.idGenerator.NewID()
	if err != nil {
		log.Error(err)
		return
	}

	img.ID = id
	imgTime := time.Now()
	img.Time = imgTime.Format(time.RFC3339)

	_, err = base64.StdEncoding.DecodeString(img.Content)

	if err != nil {
		err := errors.New("CREATE handler: content is not an image" + err.Error())
		log.Error(err)
		http.Error(w, err.Error(), http.StatusNotAcceptable)
		return
	}

	jsonImg, err := json.Marshal(img)
	if err != nil {
		err = errors.New("CREATE handler: failed to convert json into marshal: " + err.Error())
		log.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = h.dbManager.Insert(id, string(jsonImg))
	if err != nil {
		err = errors.New("CREATE handler: failed to insert values to database: " + err.Error())
		log.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var idStruct model.ID
	idStruct.ID = id

	jsonID, err := json.Marshal(idStruct)
	if err != nil {
		err = errors.New("CREATE handler: failed to convert json into marshal: " + err.Error())
		log.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	eventId, err := h.idGenerator.NewID()
	if err != nil {
		err = errors.New("CREATE handler: failed to generate id for event: " + err.Error())
		log.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = h.eventBus.SendNewImage(eventId, img)
	if err != nil {
		err = errors.New("CREATE handler: failed to send an event" + err.Error())
		log.Error(err)
		http.Error(w, err.Error(), http.StatusBadGateway)
		return
	}

	fmt.Fprint(w, string(jsonID))
	log.Info("create: succeeded")
}

// Update processes a request, that modify values in database with given JSON from GCP API
func (h Handler) Update(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	key := params["id"]
	url := strings.Replace(h.putEndpoint, "{id}", key, 1)
	if r.URL.Path != url {
		log.Error(h.putEndpoint)
		err := errors.New("update: 404 not found")
		log.Error(err)
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	accessControl(w, r)

	var img model.Image

	fromDB, err := h.dbManager.Get(key)

	if err != nil {
		if err.Error() == "GET from db:key "+key+" does not exist" {
			err = errors.New("update: failed to get data from db: " + err.Error())
			http.Error(w, err.Error(), http.StatusNotFound)
		} else {
			err = errors.New("update: failed to get data from db: " + err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		log.Error(err)
		return
	}

	err = json.Unmarshal([]byte(fromDB.(string)), &img)
	if err != nil {
		err = errors.New("update: failed to convert marshal to json: " + err.Error())
		log.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	headerContentType := r.Header.Get("Content-Type")
	if headerContentType != "application/json" {
		err := errors.New("PUT: invalid content type")
		log.Error(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	value, err := ioutil.ReadAll(r.Body)
	if err != nil {
		err = errors.New("update: failed to read request body" + err.Error())
		log.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	img.GCP = append(img.GCP, string(value))

	jsonImg, err := json.Marshal(img)
	if err != nil {
		err = errors.New("update: failed to convert json into marshal: " + err.Error())
		log.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = h.dbManager.Insert(img.ID, string(jsonImg))
	if err != nil {
		err = errors.New("update: failed to insert values to database: " + err.Error())
		log.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var idStruct model.ID
	idStruct.ID = img.ID

	jsonID, err := json.Marshal(idStruct)
	if err != nil {
		err = errors.New("update: failed to convert json id into marshal: " + err.Error())
		log.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprint(w, string(jsonID))
	log.Info("update: succeeded")
}

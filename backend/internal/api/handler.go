package api

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/gabriel-vasile/mimetype"
	"github.com/gorilla/mux"
	"github.com/kyma-incubator/Kyma-Showcase/internal/model"
	"github.com/kyma-project/kyma/common/logging/logger"
	"github.com/pkg/errors"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
	"runtime"
	"strings"
	"time"
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
	log			   *logger.Logger
	getEndpoint    string
	getAllEndpoint string
	postEndpoint   string
	putEndpoint    string
}

type LogHandler struct{
		Log			*logger.Logger
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

func NewLogHandler(logHandler *logger.Logger) LogHandler{
	return LogHandler{
		Log: logHandler,
	}
}

// NewHandler returns handler for database manager.
func NewHandler(dbManager DBManager, idGenerator IdGenerator, eventBus EventBus, log *logger.Logger) Handler {
	return Handler{
		dbManager:      dbManager,
		idGenerator:    idGenerator,
		eventBus:       eventBus,
		log:			log,
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

// handleError wraps error formatting, logging and response sending
func (h Handler)handleError(w http.ResponseWriter, code int, format string, a ...interface{}) {
	err := errors.New(fmt.Sprintf(format, a...))
	h.log.WithContext().Error(err)
	http.Error(w, err.Error(), code)
}

// getFuncName returns current function name in lowercase, used for logs
func(lh LogHandler) getFuncName() string {
	pc, _, _, ok := runtime.Caller(1)
	if !ok {
		return ""
	}
	return strings.ToLower(runtime.FuncForPC(pc).Name()[strings.LastIndex(runtime.FuncForPC(pc).Name(), ".")+1:])
}

// Get processes a request and passes request ID to the GetFromDB function, returns the value of the given ID.
func (h Handler) Get(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	key := params["id"]
	url := strings.Replace(h.getEndpoint, "{id}", key, 1)
	if r.URL.Path != url {
		h.handleError(w, http.StatusNotFound, "%s: 404 not found", getFuncName())
		return
	}
	accessControl(w, r)
	var img model.Image

	fromDB, err := h.dbManager.Get(key)

	if err != nil {
		if err.Error() == "key "+key+" does not exist" {
			h.handleError(w, http.StatusNotFound, "%s: failed to get data from db: %s", getFuncName(), err)
		} else {
			h.handleError(w, http.StatusInternalServerError, "%s: failed to get data from db: %s", getFuncName(), err)
		}
		return
	}

	err = json.Unmarshal([]byte(fromDB.(string)), &img)
	if err != nil {
		h.handleError(w, http.StatusInternalServerError, "%s: failed to convert marshal to json: %s", getFuncName(), err)
		return
	}

	fmt.Fprintf(w, "%s", fromDB)
}

//GetAll processes a request and gets all keys using GetAllKeys function, returns all values from database as a string with JSON array.
func (h Handler) GetAll(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != h.getAllEndpoint {
		h.handleError(w, http.StatusNotFound, "%s: 404 not found", getFuncName())
		return
	}
	accessControl(w, r)

	keys, err := h.dbManager.GetAll()
	if err != nil {
		h.handleError(w, http.StatusInternalServerError, "%s: failed to get all keys from db: %s", getFuncName(), err)
		return
	}

	var img model.Image
	var result []model.Image
	for _, key := range keys {
		fromDB, err := h.dbManager.Get(key)
		if err != nil {
			h.handleError(w, http.StatusInternalServerError, "%s: failed to get value from db: %s", getFuncName(), err)
			return
		}
		dataStr, ok := fromDB.(string)

		if !ok {
			h.handleError(w, http.StatusInternalServerError, "%s: failed to assert a type", getFuncName())
			return
		}
		err = json.Unmarshal([]byte(dataStr), &img)

		if err != nil {
			h.handleError(w, http.StatusInternalServerError, "%s: fail to unmarshal: %s", getFuncName(), err)
			return
		}

		result = append(result, img)
	}
	allImages, err := json.Marshal(result)

	if err != nil {
		h.handleError(w, http.StatusInternalServerError, "%s: fail to marshal: %s", getFuncName(), err)
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
		h.handleError(w, http.StatusNotFound, "%s: 404 not found", getFuncName())
		return
	}

	accessControl(w, r)
	var img model.Image

	headerContentType := r.Header.Get("Content-Type")
	if headerContentType != "application/json" {
		h.handleError(w, http.StatusBadRequest, "%s: invalid content type", getFuncName())
		return
	}

	decoder := json.NewDecoder(r.Body)
	for {
		if err := decoder.Decode(&img); err == io.EOF {
			break
		} else if err != nil {
			h.handleError(w, http.StatusBadRequest, "%s: invalid input: %s", getFuncName(), err)
			return
		}
	}

	_, err := url.ParseRequestURI(img.Content)
	u, err := url.Parse(img.Content)
	if err == nil && u.Scheme != "" && u.Host != "" {
		resp, err := http.Get(img.Content)
		if err != nil {
			h.handleError(w, http.StatusNotAcceptable, "%s: could not get image from: %s", getFuncName(), err)
			return
		}
		defer resp.Body.Close()

		imgByte, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			h.handleError(w, http.StatusInternalServerError, "%s: failed to read request body: %s", getFuncName(), err)
			return
		}

		imgBase64 := base64.StdEncoding.EncodeToString(imgByte)

		if calculateSize(imgBase64) > 5000000 {
			h.handleError(w, http.StatusTeapot, "%s: image from url is too large", getFuncName())
			return
		}

		ext := getExtension(imgByte)
		if ext == "" {
			h.handleError(w, http.StatusUnsupportedMediaType, "%s: extension is not supported", getFuncName())
			return
		}

		img.Content = fmt.Sprintf("data:%s;base64,", ext)
		img.Content += imgBase64
	} else {
		r := regexp.MustCompile("data:.*?base64,")
		contentBase64 := r.ReplaceAllString(img.Content, "")
		_, err = base64.StdEncoding.DecodeString(contentBase64)

		if err != nil {
			h.handleError(w, http.StatusNotAcceptable, "%s: content is not an image: %s", getFuncName(), err)
			return
		}
	}

	if img.Content == "" {
		h.handleError(w, http.StatusInternalServerError, "%s: content is empty", getFuncName())
		return
	}

	id, err := h.idGenerator.NewID()
	if err != nil {
		h.handleError(w, http.StatusInternalServerError, "%s: failed to generate id for image: %s", getFuncName(), err)
		return
	}

	img.ID = id

	imgTime := time.Now()
	img.Time = imgTime.Format(time.RFC3339)

	jsonImg, err := json.Marshal(img)
	if err != nil {
		h.handleError(w, http.StatusInternalServerError, "%s: failed to convert json into marshal: %s", getFuncName(), err)
		return
	}

	err = h.dbManager.Insert(id, string(jsonImg))
	if err != nil {
		h.handleError(w, http.StatusInternalServerError, "%s: failed to insert values to database: %s", getFuncName(), err)
		return
	}

	var idStruct model.ID
	idStruct.ID = id

	jsonID, err := json.Marshal(idStruct)
	if err != nil {
		h.handleError(w, http.StatusInternalServerError, "%s: failed to convert json into marshal: %s", getFuncName(), err)
		return
	}

	eventId, err := h.idGenerator.NewID()
	if err != nil {
		h.handleError(w, http.StatusInternalServerError, "%s: failed to generate id for event: %s", getFuncName(), err)
		return
	}
	err = h.eventBus.SendNewImage(eventId, img)
	if err != nil {
		h.handleError(w, http.StatusBadGateway, "%s: failed to send an event: %s", getFuncName(), err)
		return
	}

	fmt.Fprint(w, string(jsonID))
	h.log.WithContext().Info(getFuncName() + ": succeeded")
}

// Update processes a request, that modify values in database with given JSON from GCP API
func (h Handler) Update(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	key := params["id"]
	url := strings.Replace(h.putEndpoint, "{id}", key, 1)
	if r.URL.Path != url {
		h.handleError(w, http.StatusNotFound, "%s: 404 not found", getFuncName())
		return
	}
	accessControl(w, r)

	var img model.Image

	fromDB, err := h.dbManager.Get(key)

	if err != nil {
		if err.Error() == "GET from db:key "+key+" does not exist" {
			h.handleError(w, http.StatusNotFound, "%s: failed to get data from db: %s", getFuncName(), err)
		} else {
			h.handleError(w, http.StatusInternalServerError, "%s: failed to get data from db: %s", getFuncName(), err)
		}
		return
	}

	err = json.Unmarshal([]byte(fromDB.(string)), &img)
	if err != nil {
		h.handleError(w, http.StatusInternalServerError, "%s: failed to convert marshal to json: %s", getFuncName(), err)
		return
	}

	headerContentType := r.Header.Get("Content-Type")
	if headerContentType != "application/json" {
		h.handleError(w, http.StatusBadRequest, "%s: invalid content type", getFuncName())
		return
	}
	value, err := ioutil.ReadAll(r.Body)
	if err != nil {
		h.handleError(w, http.StatusInternalServerError, "%s: failed to read request body: %s", getFuncName(), err)
		return
	}
	defer r.Body.Close()

	img.GCP = append(img.GCP, string(value))

	jsonImg, err := json.Marshal(img)
	if err != nil {
		h.handleError(w, http.StatusInternalServerError, "%s: failed to convert json into marshal: %s", getFuncName(), err)
		return
	}

	err = h.dbManager.Insert(img.ID, string(jsonImg))
	if err != nil {
		h.handleError(w, http.StatusInternalServerError, "%s: failed to insert values to database: %s", getFuncName(), err)

		return
	}

	var idStruct model.ID
	idStruct.ID = img.ID

	jsonID, err := json.Marshal(idStruct)
	if err != nil {
		h.handleError(w, http.StatusInternalServerError, "%s: failed to convert json id into marshal: %s", getFuncName(), err)
		return
	}

	fmt.Fprint(w, string(jsonID))
	h.log.WithContext().Info(getFuncName() + ": succeeded")
}

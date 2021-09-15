package api

import (
	"bou.ke/monkey"
	"bytes"
	"encoding/base64"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/kyma-incubator/Kyma-Showcase/internal/api/mocks"
	"github.com/kyma-incubator/Kyma-Showcase/internal/model"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"testing/iotest"
	"time"
)

const fixedID = "FEA98D88-0669-4FFD-B17A-8F80BB97C381"

func fixedTime() string {
	monkey.Patch(time.Now, func() time.Time {
		return time.Date(2009, 11, 10, 20, 34, 58, 651387237, time.UTC)
	})
	return time.Now().Format(time.RFC3339)
}

func TestGet(t *testing.T) {
	img := model.Image{
		ID:      fixedID,
		Content: "base64",
		GCP:     []string{"{labels:[labels,moods]}"},
		Time:    fixedTime(),
	}
	jsonImg, err := json.Marshal(img)
	assert.NoError(t, err)

	tests := []struct {
		testMessage   string
		requestURL    string
		getReturned   string
		getError      error
		assertNoOfGet int
		bodyContains  string
		statusCode    int
	}{
		{
			testMessage:   "should return error with status code 404 when url is wrong",
			requestURL:    "/v1/images/{id}/wrong",
			assertNoOfGet: 0,
			bodyContains:  "get handler: 404 not found",
			statusCode:    http.StatusNotFound,
		},
		{
			testMessage:   "should return error with status code 500 when database does not respond",
			requestURL:    "/v1/images/" + fixedID,
			getError:      errors.New("GET handler: database not respond error"),
			assertNoOfGet: 1,
			bodyContains:  "GET handler: database not respond error",
			statusCode:    http.StatusInternalServerError,
		},
		{
			testMessage:   "should return error with status code 404 when key does not exist in database",
			requestURL:    "/v1/images/" + fixedID,
			getError:      errors.New("GET from db: key " + fixedID + " does not exist"),
			assertNoOfGet: 1,
			bodyContains:  "GET from db: key " + fixedID + " does not exist",
			statusCode:    http.StatusNotFound,
		},
		{
			testMessage:   "should return error with status code 500 when key has no value assigned",
			requestURL:    "/v1/images/" + fixedID,
			getReturned:   "",
			getError:      errors.New("GET handler:for key " + fixedID + " value is empty"),
			assertNoOfGet: 1,
			bodyContains:  "GET handler:for key " + fixedID + " value is empty",
			statusCode:    http.StatusInternalServerError,
		},
		{
			testMessage:   "should return error with status code 500 when key exist but value is not json",
			requestURL:    "/v1/images/" + fixedID,
			getReturned:   "not json",
			assertNoOfGet: 1,
			bodyContains:  "GET handler: failed to convert marshal to json:",
			statusCode:    http.StatusInternalServerError,
		},
		{
			testMessage:   "should return status code 200 when key exists in database and value is correct",
			requestURL:    "/v1/images/" + fixedID,
			getReturned:   string(jsonImg),
			assertNoOfGet: 1,
			bodyContains:  string(jsonImg),
			statusCode:    http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.testMessage, func(t *testing.T) {
			//given
			req, err := http.NewRequest("GET", tt.requestURL, nil)
			vars := map[string]string{
				"id": fixedID,
			}
			req = mux.SetURLVars(req, vars)
			assert.NoError(t, err)
			recorder := httptest.NewRecorder()
			dbManagerMock := mocks.DBManager{}
			idMock := mocks.IdGenerator{}
			idMock.On("NewID").Return(fixedID, nil)
			eventBusMock := mocks.EventBus{}
			testSubject := NewHandler(&dbManagerMock, &idMock, &eventBusMock)
			dbManagerMock.On("Get", fixedID).Return(tt.getReturned, tt.getError)

			//when
			testSubject.Get(recorder, req)

			//then
			dbManagerMock.AssertNumberOfCalls(t, "Get", tt.assertNoOfGet)
			assert.Contains(t, recorder.Body.String(), tt.bodyContains)
			assert.Equal(t, tt.statusCode, recorder.Code)
		})
	}
}

func TestGetAll(t *testing.T) {
	img := model.Image{
		ID:      fixedID,
		Content: "base64",
		GCP:     []string{"{labels:[labels,moods]}"},
		Time:    fixedTime(),
	}
	jsonImg, err := json.Marshal(img)
	expected := []model.Image{img, img}
	jsonExpected, err := json.Marshal(expected)
	assert.NoError(t, err)

	tests := []struct {
		testMessage       string
		requestURL        string
		getAllReturned    []string
		getAllError       error
		getFirstReturned  string
		getSecondReturned string
		getFirstError     error
		getSecondError    error
		assertNoOfGetAll  int
		assertNoOfGet     int
		bodyContains      string
		statusCode        int
	}{
		{
			testMessage:      "should return error with status code 404 when url is wrong",
			requestURL:       "/v1/images/wrong",
			getAllReturned:   nil,
			assertNoOfGetAll: 0,
			assertNoOfGet:    0,
			bodyContains:     "GETALL handler: 404 not found",
			statusCode:       http.StatusNotFound,
		},
		{
			testMessage:      "should return empty value when database is empty",
			requestURL:       "/v1/images",
			getAllReturned:   nil,
			assertNoOfGetAll: 1,
			assertNoOfGet:    0,
			bodyContains:     "",
			statusCode:       http.StatusOK,
		},
		{
			testMessage:       "should return 500 code when all values are empty",
			requestURL:        "/v1/images",
			getAllReturned:    []string{"1", "2"},
			getFirstReturned:  "",
			getFirstError:     errors.New("value is empty"),
			getSecondReturned: "",
			getSecondError:    errors.New("value is empty"),
			assertNoOfGetAll:  1,
			assertNoOfGet:     1,
			bodyContains:      "GETALL handler: failed to get value from db value is empty",
			statusCode:        http.StatusInternalServerError,
		},
		{
			testMessage:       "should return 500 code when one of the values is empty",
			requestURL:        "/v1/images",
			getAllReturned:    []string{"1", "2"},
			getFirstReturned:  string(jsonImg),
			getSecondReturned: "",
			getSecondError:    errors.New("value is empty"),
			assertNoOfGetAll:  1,
			assertNoOfGet:     2,
			bodyContains:      "GETALL handler: failed to get value from db value is empty",
			statusCode:        http.StatusInternalServerError,
		},
		{
			testMessage:       "should return JSON array compatible with given data",
			requestURL:        "/v1/images",
			getAllReturned:    []string{"1", "2"},
			getFirstReturned:  string(jsonImg),
			getSecondReturned: string(jsonImg),
			assertNoOfGetAll:  1,
			assertNoOfGet:     2,
			bodyContains:      string(jsonExpected),
			statusCode:        http.StatusOK,
		},
		{
			testMessage:    "should return error while is error during unmarshal",
			requestURL:     "/v1/images",
			getAllReturned: []string{"1", "2"},
			getAllError:    nil,
			getFirstReturned: `
					{
						id":` + fixedID + `,
						"content":"base64",
					}`,
			getSecondReturned: string(jsonImg),
			assertNoOfGetAll:  1,
			assertNoOfGet:     1,
			bodyContains:      "GETALL handler: fail to unmarshal",
			statusCode:        http.StatusInternalServerError,
		},
		{
			testMessage:       "should return 200 while function can create valid JSON array",
			requestURL:        "/v1/images",
			getAllReturned:    []string{"1", "2"},
			getFirstReturned:  string(jsonImg),
			getSecondReturned: string(jsonImg),
			assertNoOfGetAll:  1,
			assertNoOfGet:     2,
			bodyContains:      string(jsonExpected),
			statusCode:        http.StatusOK,
		},
	}
	for _, tt := range tests {
		t.Run(tt.testMessage, func(t *testing.T) {
			//given
			req, err := http.NewRequest("GET", tt.requestURL, nil)
			assert.NoError(t, err)
			recorder := httptest.NewRecorder()
			dbManagerMock := mocks.DBManager{}
			idMock := mocks.IdGenerator{}
			idMock.On("NewID").Return(fixedID, nil)
			eventBusMock := mocks.EventBus{}
			testSubject := NewHandler(&dbManagerMock, &idMock, &eventBusMock)
			dbManagerMock.On("GetAll").Return(tt.getAllReturned, nil)

			dbManagerMock.On("Get", "1").Return(tt.getFirstReturned, tt.getFirstError)
			dbManagerMock.On("Get", "2").Return(tt.getSecondReturned, tt.getSecondError)

			//when
			testSubject.GetAll(recorder, req)

			//then
			dbManagerMock.AssertNumberOfCalls(t, "GetAll", tt.assertNoOfGetAll)
			dbManagerMock.AssertNumberOfCalls(t, "Get", tt.assertNoOfGet)
			assert.Contains(t, recorder.Body.String(), tt.bodyContains)
			assert.Equal(t, tt.statusCode, recorder.Code)
		})
	}
}

func TestCreate(t *testing.T) {

	img := model.Image{
		ID:      fixedID,
		Content: "data:image/png;base64," + base64.StdEncoding.EncodeToString([]byte("data")),
		Time:    fixedTime(),
	}
	jsonImg, err := json.Marshal(img)
	assert.NoError(t, err)

	tests := []struct {
		testMessage      string
		requestURL       string
		body             io.Reader
		contentType      string
		insertArg        string
		insertError      error
		eventError       error
		assertNoOfInsert int
		bodyContains     string
		logContains      string
		statusCode       int
	}{
		{
			testMessage:      "should return error with status code 404 when url is wrong",
			requestURL:       "/v1/images/wrong",
			body:             bytes.NewBuffer(jsonImg),
			contentType:      "application/json",
			assertNoOfInsert: 0,
			bodyContains:     "CREATE handler: 404 not found",
			logContains:      "CREATE handler: 404 not found",
			statusCode:       http.StatusNotFound,
		},
		{
			testMessage:      "should return 400 when Content-Type is incorrect",
			requestURL:       "/v1/images",
			body:             bytes.NewBuffer(jsonImg),
			contentType:      "application/golang",
			assertNoOfInsert: 0,
			bodyContains:     "CREATE handler: invalid content type",
			logContains:      "CREATE handler: invalid content type",
			statusCode:       http.StatusBadRequest,
		},
		{
			testMessage:      "should return 400 error when request body is not json",
			requestURL:       "/v1/images",
			body:             bytes.NewBuffer([]byte("string")),
			contentType:      "application/json",
			assertNoOfInsert: 0,
			bodyContains:     "CREATE handler: invalid input:",
			logContains:      "CREATE handler: invalid input:",
			statusCode:       http.StatusBadRequest,
		},
		{
			testMessage:      "should return 500 error when unable to insert json to db",
			requestURL:       "/v1/images",
			body:             bytes.NewBuffer(jsonImg),
			contentType:      "application/json",
			insertArg:        string(jsonImg),
			insertError:      errors.New("failed to insert json to db"),
			assertNoOfInsert: 1,
			bodyContains:     "CREATE handler: failed to insert values to database:",
			logContains:      "CREATE handler: failed to insert values to database:",
			statusCode:       http.StatusInternalServerError,
		},
		{
			testMessage:      "should log proper error when sending event failed",
			requestURL:       "/v1/images",
			body:             bytes.NewBuffer(jsonImg),
			contentType:      "application/json",
			insertArg:        string(jsonImg),
			eventError:       errors.New("SENDEVENT: error"),
			assertNoOfInsert: 1,
			bodyContains:     "SENDEVENT: error",
			logContains:      "SENDEVENT: error",
			statusCode:       http.StatusBadGateway,
		},
		{
			testMessage:      "should return 200 code when request, data and connection with database are correct",
			requestURL:       "/v1/images",
			body:             bytes.NewBuffer(jsonImg),
			contentType:      "application/json",
			insertArg:        string(jsonImg),
			assertNoOfInsert: 1,
			bodyContains:     fixedID,
			logContains:      "succeeded",
			statusCode:       http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.testMessage, func(t *testing.T) {

			//given
			hook := test.NewGlobal()
			req, err := http.NewRequest("POST", tt.requestURL, tt.body)
			assert.NoError(t, err)
			req.Header.Set("Content-Type", tt.contentType)
			recorder := httptest.NewRecorder()
			dbManagerMock := mocks.DBManager{}
			idMock := mocks.IdGenerator{}
			idMock.On("NewID").Return(fixedID, nil)
			eventBusMock := mocks.EventBus{}
			testSubject := NewHandler(&dbManagerMock, &idMock, &eventBusMock)
			dbManagerMock.On("Insert", fixedID, tt.insertArg).Return(tt.insertError)
			eventBusMock.On("SendNewImage", fixedID, img).Return(tt.eventError)

			//when
			testSubject.Create(recorder, req)

			//then
			dbManagerMock.AssertNumberOfCalls(t, "Insert", tt.assertNoOfInsert)
			assert.Contains(t, recorder.Body.String(), tt.bodyContains)
			assert.Contains(t, hook.LastEntry().Message, tt.logContains)
			assert.Equal(t, tt.statusCode, recorder.Code)
		})
	}

	t.Run("should return error with status code 406 when request content is corrupted", func(t *testing.T) {

		//given
		img := model.Image{
			ID:      fixedID,
			Content: "notAnImage",
			Time:    time.Now().Format(time.RFC3339),
		}
		jsonImg, err := json.Marshal(img)
		assert.NoError(t, err)
		hook := test.NewGlobal()
		req, err := http.NewRequest("POST", "/v1/images", bytes.NewBuffer(jsonImg))
		assert.NoError(t, err)
		req.Header.Set("Content-Type", "application/json")
		recorder := httptest.NewRecorder()
		dbManagerMock := mocks.DBManager{}
		idMock := mocks.IdGenerator{}
		idMock.On("NewID").Return(fixedID, nil)
		eventBusMock := mocks.EventBus{}
		testSubject := NewHandler(&dbManagerMock, &idMock, &eventBusMock)
		dbManagerMock.On("Insert", fixedID, string(jsonImg)).Return(nil)
		eventBusMock.On("SendNewImage", fixedID, img).Return(nil)

		//when
		testSubject.Create(recorder, req)

		//then
		dbManagerMock.AssertNotCalled(t, "Insert")
		assert.Contains(t, hook.LastEntry().Message, "CREATE handler: content is not an image")
		assert.Equal(t, http.StatusNotAcceptable, recorder.Code)
	})

	t.Run("should return error with status code 500 when content in unmarshaled struct is empty", func(t *testing.T) {

		//given
		hook := test.NewGlobal()
		img := model.Image{
			ID:      fixedID,
			Content: "",
			Time:    time.Now().Format(time.RFC3339),
		}
		jsonImg, err := json.Marshal(img)
		assert.NoError(t, err)
		req, err := http.NewRequest("POST", "/v1/images", bytes.NewBuffer(jsonImg))
		assert.NoError(t, err)
		req.Header.Set("Content-Type", "application/json")
		recorder := httptest.NewRecorder()
		dbManagerMock := mocks.DBManager{}
		idMock := mocks.IdGenerator{}
		idMock.On("NewID").Return(fixedID, nil)
		eventBusMock := mocks.EventBus{}
		testSubject := NewHandler(&dbManagerMock, &idMock, &eventBusMock)

		//when
		testSubject.Create(recorder, req)

		//then
		dbManagerMock.AssertNotCalled(t, "Insert")
		assert.Contains(t, hook.LastEntry().Message, "CREATE handler: content is empty")
		assert.Equal(t, http.StatusInternalServerError, recorder.Code)
	})
}

func TestUpdate(t *testing.T) {
	imgTime := fixedTime()
	body := `{` +
		`labels:[labels,moods]` +
		`}`
	img := model.Image{
		ID: fixedID,

		Content: "base64",
		Time:    imgTime,
	}
	imgWithGCP := model.Image{
		ID:      fixedID,
		Content: "base64",
		GCP:     []string{"{labels:[labels,moods]}"},
		Time:    imgTime,
	}
	jsonImg, err := json.Marshal(img)
	assert.NoError(t, err)
	jsonImgWithGCP, err := json.Marshal(imgWithGCP)
	assert.NoError(t, err)
	idStruct := model.ID{ID: fixedID}
	jsonID, err := json.Marshal(idStruct)
	assert.NoError(t, err)

	tests := []struct {
		testMessage      string
		requestURL       string
		body             io.Reader
		contentType      string
		getReturned      string
		getError         error
		insertArg        string
		insertError      error
		assertNoOfInsert int
		assertNoOfGet    int
		bodyContains     string
		logContains      string
		statusCode       int
	}{
		{
			testMessage:      "should return error with status code 404 when url is wrong",
			requestURL:       "/v1/images/{id}/wrong",
			contentType:      "application/json",
			assertNoOfGet:    0,
			assertNoOfInsert: 0,
			bodyContains:     "update: 404 not found",
			logContains:      "update: 404 not found",
			statusCode:       http.StatusNotFound,
		},
		{
			testMessage:      "should return error when key does not exists",
			requestURL:       "/v1/images/" + fixedID,
			contentType:      "application/json",
			getError:         errors.New("GET from db:key " + fixedID + " does not exist"),
			assertNoOfGet:    1,
			assertNoOfInsert: 0,
			bodyContains:     "GET from db:key " + fixedID + " does not exist",
			logContains:      "GET from db:key " + fixedID + " does not exist",
			statusCode:       http.StatusNotFound,
		},
		{
			testMessage:      "should return error with status code 500 when database does not respond",
			requestURL:       "/v1/images/" + fixedID,
			contentType:      "application/json",
			getError:         errors.New("GET from db: database not respond error"),
			assertNoOfGet:    1,
			assertNoOfInsert: 0,
			bodyContains:     "GET from db: database not respond error",
			logContains:      "GET from db: database not respond error",
			statusCode:       http.StatusInternalServerError,
		},
		{
			testMessage: "should return error while there is an error during unmarshal",
			requestURL:  "/v1/images/" + fixedID,
			contentType: "application/json",
			getReturned: `
			{
				id":` + fixedID + `,
				"content":"base64",
			}`,
			assertNoOfGet:    1,
			assertNoOfInsert: 0,
			bodyContains:     "update: failed to convert marshal to json",
			logContains:      "update: failed to convert marshal to json",
			statusCode:       http.StatusInternalServerError,
		},
		{
			testMessage:      "should return 400 when Content-Type is incorrect",
			requestURL:       "/v1/images/" + fixedID,
			contentType:      "application/golang",
			getReturned:      string(jsonImgWithGCP),
			assertNoOfGet:    1,
			assertNoOfInsert: 0,
			bodyContains:     "PUT: invalid content type",
			logContains:      "PUT: invalid content type",
			statusCode:       http.StatusBadRequest,
		},
		{
			testMessage:      "should return 500 when is an error during reading request body",
			requestURL:       "/v1/images/" + fixedID,
			body:             iotest.ErrReader(errors.New("failed to read")),
			contentType:      "application/json",
			getReturned:      string(jsonImg),
			assertNoOfGet:    1,
			assertNoOfInsert: 0,
			bodyContains:     "failed to read",
			logContains:      "update: failed to read request body",
			statusCode:       http.StatusInternalServerError,
		},
		{
			testMessage:      "should return 500 when failed to insert data to database",
			requestURL:       "/v1/images/" + fixedID,
			body:             bytes.NewBuffer([]byte(body)),
			contentType:      "application/json",
			getReturned:      string(jsonImg),
			insertArg:        string(jsonImgWithGCP),
			insertError:      errors.New("failed to insert json to db"),
			assertNoOfGet:    1,
			assertNoOfInsert: 1,
			bodyContains:     "failed to insert json to db",
			logContains:      "failed to insert json to db",
			statusCode:       http.StatusInternalServerError,
		},
		{
			testMessage:      "should return 200 code when updating data in database is correct",
			requestURL:       "/v1/images/" + fixedID,
			body:             bytes.NewBuffer([]byte(body)),
			contentType:      "application/json",
			getReturned:      string(jsonImg),
			insertArg:        string(jsonImgWithGCP),
			assertNoOfGet:    1,
			assertNoOfInsert: 1,
			bodyContains:     string(jsonID),
			logContains:      "succeeded",
			statusCode:       http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.testMessage, func(t *testing.T) {
			//given
			hook := test.NewGlobal()
			req, err := http.NewRequest("PUT", tt.requestURL, tt.body)
			vars := map[string]string{
				"id": fixedID,
			}
			req = mux.SetURLVars(req, vars)
			assert.NoError(t, err)
			req.Header.Set("Content-Type", tt.contentType)
			recorder := httptest.NewRecorder()
			dbManagerMock := mocks.DBManager{}
			idMock := mocks.IdGenerator{}
			idMock.On("NewID").Return(fixedID, nil)
			eventBusMock := mocks.EventBus{}
			testSubject := NewHandler(&dbManagerMock, &idMock, &eventBusMock)

			dbManagerMock.On("Get", fixedID).Return(tt.getReturned, tt.getError)
			dbManagerMock.On("Insert", fixedID, tt.insertArg).Return(tt.insertError)

			//when
			testSubject.Update(recorder, req)

			//then
			dbManagerMock.AssertNumberOfCalls(t, "Get", tt.assertNoOfGet)
			dbManagerMock.AssertNumberOfCalls(t, "Insert", tt.assertNoOfInsert)
			assert.Contains(t, recorder.Body.String(), tt.bodyContains)
			assert.Equal(t, tt.statusCode, recorder.Code)
			assert.Contains(t, hook.LastEntry().Message, tt.logContains)
		})
	}
}

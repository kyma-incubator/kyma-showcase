package api

import (
	"bytes"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/kyma-incubator/Kyma-Showcase/internal/api/mocks"
	"github.com/kyma-incubator/Kyma-Showcase/internal/model"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"testing/iotest"
)

const fixedID = "FEA98D88-0669-4FFD-B17A-8F80BB97C381"

func TestGet(t *testing.T) {
	const key = fixedID
	t.Run("should return error with status code 404 when url is wrong", func(t *testing.T) {

		//given
		req, err := http.NewRequest("GET", "/v1/images/{id}/wrong", nil)
		assert.NoError(t, err)
		recorder := httptest.NewRecorder()
		dbManagerMock := mocks.DBManager{}
		idMock := mocks.IdGenerator{}
		idMock.On("NewID").Return(fixedID, nil)
		eventBusMock := mocks.EventBus{}
		testSubject := NewHandler(&dbManagerMock, &idMock, &eventBusMock)
		err = errors.New("GET handler: 404 not found")

		//when
		testSubject.Get(recorder, req)

		//then
		dbManagerMock.AssertNumberOfCalls(t, "Get", 0)
		assert.Contains(t, recorder.Body.String(), err.Error())
		assert.Equal(t, http.StatusNotFound, recorder.Code)
	})

	t.Run("should return error with status code 500 when database does not respond", func(t *testing.T) {

		//given
		req, err := http.NewRequest("GET", "/v1/images/"+key, nil)
		vars := map[string]string{
			"id": key,
		}
		req = mux.SetURLVars(req, vars)
		assert.NoError(t, err)
		recorder := httptest.NewRecorder()
		dbManagerMock := mocks.DBManager{}
		idMock := mocks.IdGenerator{}
		idMock.On("NewID").Return(fixedID, nil)
		eventBusMock := mocks.EventBus{}
		testSubject := NewHandler(&dbManagerMock, &idMock, &eventBusMock)
		err = errors.New("GET handler: database not respond error")
		dbManagerMock.On("Get", key).Return(nil, err)
		//when
		testSubject.Get(recorder, req)

		//then
		dbManagerMock.AssertNumberOfCalls(t, "Get", 1)
		assert.Contains(t, recorder.Body.String(), err.Error())
		assert.Equal(t, http.StatusInternalServerError, recorder.Code)
	})

	t.Run("should return error with status code 404 when key does not exist in database", func(t *testing.T) {

		//given
		req, err := http.NewRequest("GET", "/v1/images/"+key, nil)
		vars := map[string]string{
			"id": key,
		}
		req = mux.SetURLVars(req, vars)
		assert.NoError(t, err)
		recorder := httptest.NewRecorder()
		dbManagerMock := mocks.DBManager{}
		idMock := mocks.IdGenerator{}
		idMock.On("NewID").Return(fixedID, nil)
		eventBusMock := mocks.EventBus{}
		testSubject := NewHandler(&dbManagerMock, &idMock, &eventBusMock)
		err = errors.New("GET from db: key " + key + " does not exist")
		dbManagerMock.On("Get", key).Return(nil, err)

		//when
		testSubject.Get(recorder, req)

		//then
		dbManagerMock.AssertNumberOfCalls(t, "Get", 1)
		assert.Contains(t, recorder.Body.String(), err.Error())
		assert.Equal(t, http.StatusNotFound, recorder.Code)
	})

	t.Run("should return error with status code 500 when key has no value assigned", func(t *testing.T) {

		//given
		req, err := http.NewRequest("GET", "/v1/images/"+key, nil)
		vars := map[string]string{
			"id": key,
		}
		req = mux.SetURLVars(req, vars)
		assert.NoError(t, err)
		recorder := httptest.NewRecorder()
		dbManagerMock := mocks.DBManager{}
		idMock := mocks.IdGenerator{}
		idMock.On("NewID").Return(fixedID, nil)
		eventBusMock := mocks.EventBus{}
		testSubject := NewHandler(&dbManagerMock, &idMock, &eventBusMock)
		err = errors.New("GET handler:for key " + key + " value is empty")
		dbManagerMock.On("Get", key).Return("", err)

		//when
		testSubject.Get(recorder, req)

		//then
		dbManagerMock.AssertNumberOfCalls(t, "Get", 1)
		assert.Equal(t, http.StatusInternalServerError, recorder.Code)
		assert.Contains(t, recorder.Body.String(), err.Error())
	})

	t.Run("should return error with status code 500 when key exist but value is not json", func(t *testing.T) {

		//given
		req, err := http.NewRequest("GET", "/v1/images/"+key, nil)
		value := "not json"
		vars := map[string]string{
			"id": key,
		}
		req = mux.SetURLVars(req, vars)
		assert.NoError(t, err)
		recorder := httptest.NewRecorder()
		dbManagerMock := mocks.DBManager{}
		idMock := mocks.IdGenerator{}
		idMock.On("NewID").Return(fixedID, nil)
		eventBusMock := mocks.EventBus{}
		testSubject := NewHandler(&dbManagerMock, &idMock, &eventBusMock)
		dbManagerMock.On("Get", key).Return(value, nil)

		//when
		testSubject.Get(recorder, req)

		//then
		dbManagerMock.AssertNumberOfCalls(t, "Get", 1)
		assert.Contains(t, recorder.Body.String(), "GET handler: failed to convert marshal to json:")
		assert.Equal(t, http.StatusInternalServerError, recorder.Code)
	})

	t.Run("should return status code 200 when key exists in database and value is correct", func(t *testing.T) {

		//given
		req, err := http.NewRequest("GET", "/v1/images/"+key, nil)
		img := model.Image{
			ID:      fixedID,
			Content: "base64",
			GCP:     []string{"{labels:[labels,moods]}"},
			Status:  false,
		}
		jsonImg, err := json.Marshal(img)
		vars := map[string]string{
			"id": key,
		}
		req = mux.SetURLVars(req, vars)
		assert.NoError(t, err)
		recorder := httptest.NewRecorder()
		dbManagerMock := mocks.DBManager{}
		idMock := mocks.IdGenerator{}
		idMock.On("NewID").Return(fixedID, nil)
		eventBusMock := mocks.EventBus{}
		testSubject := NewHandler(&dbManagerMock, &idMock, &eventBusMock)
		dbManagerMock.On("Get", key).Return(string(jsonImg), nil)

		//when
		testSubject.Get(recorder, req)

		//then
		dbManagerMock.AssertNumberOfCalls(t, "Get", 1)
		assert.Equal(t, string(jsonImg), recorder.Body.String())
		assert.Equal(t, http.StatusOK, recorder.Code)
	})
}

func TestGetAll(t *testing.T) {
	t.Run("should return error with status code 404 when url is wrong", func(t *testing.T) {

		//given
		req, err := http.NewRequest("GET", "/v1/images/wrong", nil)
		assert.NoError(t, err)
		recorder := httptest.NewRecorder()
		dbManagerMock := mocks.DBManager{}
		idMock := mocks.IdGenerator{}
		idMock.On("NewID").Return(fixedID, nil)
		eventBusMock := mocks.EventBus{}
		testSubject := NewHandler(&dbManagerMock, &idMock, &eventBusMock)
		err = errors.New("GETALL handler: 404 not found")

		//when
		testSubject.GetAll(recorder, req)

		//then
		dbManagerMock.AssertNumberOfCalls(t, "Get", 0)
		assert.Contains(t, recorder.Body.String(), err.Error())
		assert.Equal(t, http.StatusNotFound, recorder.Code)
	})

	t.Run("should return empty value when database is empty", func(t *testing.T) {

		//given
		req, err := http.NewRequest("GET", "/v1/images", nil)

		assert.NoError(t, err)
		recorder := httptest.NewRecorder()
		dbManagerMock := mocks.DBManager{}
		idMock := mocks.IdGenerator{}
		idMock.On("NewID").Return(fixedID, nil)
		eventBusMock := mocks.EventBus{}
		testSubject := NewHandler(&dbManagerMock, &idMock, &eventBusMock)
		dbManagerMock.On("GetAll").Return(nil, nil)

		//when
		testSubject.GetAll(recorder, req)

		//then
		dbManagerMock.AssertNumberOfCalls(t, "GetAll", 1)
		assert.Equal(t, http.StatusOK, recorder.Code)

	})

	t.Run("should return 500 code when all values are empty", func(t *testing.T) {

		//given
		req, err := http.NewRequest("GET", "/v1/images", nil)

		assert.NoError(t, err)
		recorder := httptest.NewRecorder()
		dbManagerMock := mocks.DBManager{}
		idMock := mocks.IdGenerator{}
		idMock.On("NewID").Return(fixedID, nil)
		eventBusMock := mocks.EventBus{}
		testSubject := NewHandler(&dbManagerMock, &idMock, &eventBusMock)
		dbManagerMock.On("GetAll").Return([]string{"1", "2", "3"}, nil)
		dbManagerMock.On("Get", "1").Return("", errors.New("value is empty"))
		dbManagerMock.On("Get", "2").Return("", errors.New("value is empty"))
		dbManagerMock.On("Get", "3").Return("", errors.New("value is empty"))

		//when
		testSubject.GetAll(recorder, req)

		//then
		dbManagerMock.AssertNumberOfCalls(t, "GetAll", 1)
		dbManagerMock.AssertCalled(t, "Get", "1")
		dbManagerMock.AssertNotCalled(t, "Get", "2")
		dbManagerMock.AssertNotCalled(t, "Get", "3")
		dbManagerMock.AssertNumberOfCalls(t, "Get", 1)
		assert.Equal(t, http.StatusInternalServerError, recorder.Code)

	})

	t.Run("should return 500 code when one of the values is empty", func(t *testing.T) {

		//given
		req, err := http.NewRequest("GET", "/v1/images", nil)
		assert.NoError(t, err)
		img := model.Image{
			ID:      fixedID,
			Content: "base64",
			Status:  false,
		}
		jsonImg, err := json.Marshal(img)
		assert.NoError(t, err)
		recorder := httptest.NewRecorder()
		dbManagerMock := mocks.DBManager{}
		idMock := mocks.IdGenerator{}
		idMock.On("NewID").Return(fixedID, nil)
		eventBusMock := mocks.EventBus{}
		testSubject := NewHandler(&dbManagerMock, &idMock, &eventBusMock)
		dbManagerMock.On("GetAll").Return([]string{"1", "2", "3"}, nil)
		dbManagerMock.On("Get", "1").Return(string(jsonImg), nil)
		dbManagerMock.On("Get", "2").Return(string(jsonImg), nil)
		dbManagerMock.On("Get", "3").Return("", errors.New("value is empty"))

		//when
		testSubject.GetAll(recorder, req)

		//then
		dbManagerMock.AssertNumberOfCalls(t, "GetAll", 1)
		dbManagerMock.AssertCalled(t, "Get", "1")
		dbManagerMock.AssertCalled(t, "Get", "2")
		dbManagerMock.AssertCalled(t, "Get", "3")
		dbManagerMock.AssertNumberOfCalls(t, "Get", 3)
		assert.Equal(t, http.StatusInternalServerError, recorder.Code)

	})

	t.Run("should return 500 code when error occurred during type assertion", func(t *testing.T) {

		//given
		req, err := http.NewRequest("GET", "/v1/images", nil)
		assert.NoError(t, err)
		recorder := httptest.NewRecorder()
		dbManagerMock := mocks.DBManager{}
		idMock := mocks.IdGenerator{}
		idMock.On("NewID").Return(fixedID, nil)
		eventBusMock := mocks.EventBus{}
		testSubject := NewHandler(&dbManagerMock, &idMock, &eventBusMock)
		dbManagerMock.On("GetAll").Return([]string{"1"}, nil)
		dbManagerMock.On("Get", "1").Return(100, nil)

		//when
		testSubject.GetAll(recorder, req)

		//then
		dbManagerMock.AssertNumberOfCalls(t, "GetAll", 1)
		dbManagerMock.AssertCalled(t, "Get", "1")
		dbManagerMock.AssertNumberOfCalls(t, "Get", 1)
		assert.Equal(t, http.StatusInternalServerError, recorder.Code)
	})

	t.Run("should return JSON array compatible with given data ", func(t *testing.T) {

		//given
		img := model.Image{
			ID:      fixedID,
			Content: "base64",
			Status:  false,
		}
		jsonImg, err := json.Marshal(img)
		expected := []model.Image{img, img}
		jsonExpected, err := json.Marshal(expected)
		assert.NoError(t, err)
		req, err := http.NewRequest("GET", "/v1/images", nil)
		assert.NoError(t, err)
		recorder := httptest.NewRecorder()
		dbManagerMock := mocks.DBManager{}
		idMock := mocks.IdGenerator{}
		idMock.On("NewID").Return(fixedID, nil)
		eventBusMock := mocks.EventBus{}
		testSubject := NewHandler(&dbManagerMock, &idMock, &eventBusMock)
		dbManagerMock.On("GetAll").Return([]string{"1", "2"}, nil)
		dbManagerMock.On("Get", "1").Return(string(jsonImg), nil)
		dbManagerMock.On("Get", "2").Return(string(jsonImg), nil)

		//when
		testSubject.GetAll(recorder, req)

		//then
		dbManagerMock.AssertNumberOfCalls(t, "GetAll", 1)
		dbManagerMock.AssertCalled(t, "Get", "1")
		dbManagerMock.AssertCalled(t, "Get", "2")
		dbManagerMock.AssertNumberOfCalls(t, "Get", 2)
		assert.Equal(t, string(jsonExpected), recorder.Body.String())
	})

	t.Run("should return error while is error during unmarshal", func(t *testing.T) {

		//given
		req, err := http.NewRequest("GET", "/v1/images", nil)
		assert.NoError(t, err)
		recorder := httptest.NewRecorder()
		dbManagerMock := mocks.DBManager{}
		idMock := mocks.IdGenerator{}
		idMock.On("NewID").Return(fixedID, nil)
		eventBusMock := mocks.EventBus{}
		testSubject := NewHandler(&dbManagerMock, &idMock, &eventBusMock)
		dbManagerMock.On("GetAll").Return([]string{"1", "2"}, nil)

		firstReturn := `
			{
				id":"abcd1234",
				"content":"base64",
				"gcp":"JSON1",
				"status":false
			}`

		secondReturn := `
			{
				"id":"zaqwsx"
				"content":"base64_2",
				"gcp":"JSON2"
				"status":false
			}
			`

		dbManagerMock.On("Get", "1").Return(firstReturn, nil)
		dbManagerMock.On("Get", "2").Return(secondReturn, nil)

		//when
		testSubject.GetAll(recorder, req)

		//then
		dbManagerMock.AssertNumberOfCalls(t, "GetAll", 1)
		dbManagerMock.AssertCalled(t, "Get", "1")
		dbManagerMock.AssertNotCalled(t, "Get", "2")
		dbManagerMock.AssertNumberOfCalls(t, "Get", 1)
		assert.Equal(t, http.StatusInternalServerError, recorder.Code)
	})

	t.Run("should return 200 while function can create valid JSON array", func(t *testing.T) {

		//given
		req, err := http.NewRequest("GET", "/v1/images", nil)
		assert.NoError(t, err)
		recorder := httptest.NewRecorder()
		dbManagerMock := mocks.DBManager{}
		idMock := mocks.IdGenerator{}
		idMock.On("NewID").Return(fixedID, nil)
		eventBusMock := mocks.EventBus{}
		testSubject := NewHandler(&dbManagerMock, &idMock, &eventBusMock)
		dbManagerMock.On("GetAll").Return([]string{"1", "2"}, nil)

		img := model.Image{
			ID:      fixedID,
			Content: "base64",
			Status:  false,
		}
		jsonImg, err := json.Marshal(img)
		expected := []model.Image{img, img}
		jsonExpected, err := json.Marshal(expected)
		dbManagerMock.On("Get", "1").Return(string(jsonImg), nil)
		dbManagerMock.On("Get", "2").Return(string(jsonImg), nil)

		//when
		testSubject.GetAll(recorder, req)

		//then
		dbManagerMock.AssertNumberOfCalls(t, "GetAll", 1)
		dbManagerMock.AssertCalled(t, "Get", "1")
		dbManagerMock.AssertCalled(t, "Get", "2")
		dbManagerMock.AssertNumberOfCalls(t, "Get", 2)
		assert.Equal(t, string(jsonExpected), recorder.Body.String())
		assert.Equal(t, http.StatusOK, recorder.Code)
	})
}

func TestCreate(t *testing.T) {
	t.Run("should return error with status code 404 when url is wrong", func(t *testing.T) {

		//given
		req, err := http.NewRequest("POST", "/v1/images/wrong", nil)
		assert.NoError(t, err)
		recorder := httptest.NewRecorder()
		dbManagerMock := mocks.DBManager{}
		idMock := mocks.IdGenerator{}
		idMock.On("NewID").Return(fixedID, nil)
		eventBusMock := mocks.EventBus{}
		testSubject := NewHandler(&dbManagerMock, &idMock, &eventBusMock)
		err = errors.New("CREATE handler: 404 not found")

		//when
		testSubject.Create(recorder, req)

		//then
		dbManagerMock.AssertNumberOfCalls(t, "Get", 0)
		assert.Contains(t, recorder.Body.String(), err.Error())
		assert.Equal(t, http.StatusNotFound, recorder.Code)
	})

	t.Run("should return 400 when Content-Type is incorrect", func(t *testing.T) {

		//given
		img := model.Image{
			ID:      fixedID,
			Content: "base64",
			GCP:     []string{"labels", "moods"},
			Status:  false,
		}
		jsonImg, err := json.Marshal(img)
		assert.NoError(t, err)
		req, err := http.NewRequest("POST", "/v1/images", bytes.NewBuffer(jsonImg))
		assert.NoError(t, err)
		req.Header.Set("Content-Type", "application/golang")
		recorder := httptest.NewRecorder()
		dbManagerMock := mocks.DBManager{}
		idMock := mocks.IdGenerator{}
		idMock.On("NewID").Return(fixedID, nil)
		eventBusMock := mocks.EventBus{}
		testSubject := NewHandler(&dbManagerMock, &idMock, &eventBusMock)

		//when
		testSubject.Create(recorder, req)

		//then
		assert.Equal(t, http.StatusBadRequest, recorder.Code)
	})

	t.Run("should return 400 error when request body is not json", func(t *testing.T) {

		//given
		var jsonStr = "string"
		req, err := http.NewRequest("POST", "/v1/images", bytes.NewBuffer([]byte(jsonStr)))
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
		dbManagerMock.AssertNumberOfCalls(t, "Create", 0)
		assert.Contains(t, recorder.Body.String(), "CREATE handler: invalid input:")
		assert.Equal(t, http.StatusBadRequest, recorder.Code)
	})

	t.Run("should return 500 error when unable to insert json to db", func(t *testing.T) {

		//given
		img := model.Image{
			ID:      fixedID,
			Content: "base64",
			Status:  false,
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
		err = errors.New("failed to insert json to db")
		dbManagerMock.On("Insert", fixedID, string(jsonImg)).Return(err)

		//when
		testSubject.Create(recorder, req)

		//then
		dbManagerMock.AssertNumberOfCalls(t, "Insert", 1)
		assert.Contains(t, recorder.Body.String(), "CREATE handler: failed to insert values to database:")
		assert.Equal(t, http.StatusInternalServerError, recorder.Code)
	})

	t.Run("should return 200 code when request, data and connection with database are correct", func(t *testing.T) {

		//given
		img := model.Image{
			ID:      fixedID,
			Content: "base64",
			GCP:     []string{"labels", "moods"},
			Status:  false,
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
		dbManagerMock.On("Insert", fixedID, string(jsonImg)).Return(nil)
		eventBusMock.On("SendNewImage", fixedID, img).Return(nil)
		//when
		testSubject.Create(recorder, req)

		//then
		dbManagerMock.AssertNumberOfCalls(t, "Insert", 1)
		assert.Contains(t, recorder.Body.String(), fixedID)
		assert.Equal(t, http.StatusOK, recorder.Code)
		eventBusMock.AssertExpectations(t)

	})

	t.Run("should log proper error when sending event failed", func(t *testing.T) {
		//given
		hook := test.NewGlobal()
		img := model.Image{
			ID:      fixedID,
			Content: "base64",
			GCP:     []string{"labels", "moods"},
			Status:  false,
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
		dbManagerMock.On("Insert", fixedID, string(jsonImg)).Return(nil)
		eventBusMock.On("SendNewImage", fixedID, img).Return(errors.New("SENDEVENT: error"))

		//when
		testSubject.Create(recorder, req)

		//then
		assert.Contains(t, hook.LastEntry().Message, "SENDEVENT")
		assert.Equal(t, http.StatusBadGateway, recorder.Code)
		eventBusMock.AssertExpectations(t)
	})
}

func TestUpdate(t *testing.T) {
	const key = fixedID
	t.Run("should return error with status code 404 when url is wrong", func(t *testing.T) {

		//given
		req, err := http.NewRequest("PUT", "/v1/images/{id}/wrong", nil)
		recorder := httptest.NewRecorder()
		dbManagerMock := mocks.DBManager{}
		idMock := mocks.IdGenerator{}
		eventBusMock := mocks.EventBus{}
		testSubject := NewHandler(&dbManagerMock, &idMock, &eventBusMock)
		err = errors.New("update: 404 not found")

		//when
		testSubject.Update(recorder, req)

		//then
		dbManagerMock.AssertNumberOfCalls(t, "GetFromDB", 0)
		dbManagerMock.AssertNumberOfCalls(t, "InsertToDB", 0)
		assert.Contains(t, recorder.Body.String(), err.Error())
		assert.Equal(t, http.StatusNotFound, recorder.Code)
	})

	t.Run("should return error when key does not exists", func(t *testing.T) {
		//given
		req, err := http.NewRequest("PUT", "/v1/images/"+key, nil)
		vars := map[string]string{
			"id": key,
		}
		req = mux.SetURLVars(req, vars)
		assert.NoError(t, err)
		recorder := httptest.NewRecorder()
		dbManagerMock := mocks.DBManager{}
		idMock := mocks.IdGenerator{}
		eventBusMock := mocks.EventBus{}
		testSubject := NewHandler(&dbManagerMock, &idMock, &eventBusMock)
		err = errors.New("GETFROMDB:key " + key + " does not exist")
		dbManagerMock.On("GetFromDB", key).Return(nil, err)

		//when
		testSubject.Update(recorder, req)

		//then
		dbManagerMock.AssertNumberOfCalls(t, "GetFromDB", 1)
		dbManagerMock.AssertNumberOfCalls(t, "InsertToDB", 0)
		assert.Contains(t, recorder.Body.String(), err.Error())
		assert.Equal(t, http.StatusNotFound, recorder.Code)
	})

	t.Run("should return error with status code 500 when database does not respond", func(t *testing.T) {

		//given
		req, err := http.NewRequest("PUT", "/v1/images/"+key, nil)
		vars := map[string]string{
			"id": key,
		}
		req = mux.SetURLVars(req, vars)
		assert.NoError(t, err)
		recorder := httptest.NewRecorder()
		dbManagerMock := mocks.DBManager{}
		idMock := mocks.IdGenerator{}
		eventBusMock := mocks.EventBus{}
		testSubject := NewHandler(&dbManagerMock, &idMock, &eventBusMock)
		err = errors.New("GETFROMDB: database not respond error")
		dbManagerMock.On("GetFromDB", key).Return(nil, err)

		//when
		testSubject.Update(recorder, req)

		//then
		dbManagerMock.AssertNumberOfCalls(t, "GetFromDB", 1)
		dbManagerMock.AssertNumberOfCalls(t, "InsertToDB", 0)
		assert.Contains(t, recorder.Body.String(), err.Error())
		assert.Equal(t, http.StatusInternalServerError, recorder.Code)
	})

	t.Run("should return error while there is an error during unmarshal", func(t *testing.T) {

		//given
		hook := test.NewGlobal()
		returnValue := `
			{
				id":"abcd1234",
				"content":"base64",
				"gcp":"JSON1",
				"status":false
			}`
		req, err := http.NewRequest("PUT", "/v1/images/"+key, nil)
		assert.NoError(t, err)
		vars := map[string]string{
			"id": key,
		}
		req = mux.SetURLVars(req, vars)
		assert.NoError(t, err)
		recorder := httptest.NewRecorder()
		dbManagerMock := mocks.DBManager{}
		idMock := mocks.IdGenerator{}
		eventBusMock := mocks.EventBus{}
		testSubject := NewHandler(&dbManagerMock, &idMock, &eventBusMock)
		dbManagerMock.On("GetFromDB", key).Return(returnValue, nil)

		//when
		testSubject.Update(recorder, req)

		//then
		dbManagerMock.AssertNumberOfCalls(t, "InsertToDB", 0)
		dbManagerMock.AssertNumberOfCalls(t, "GetFromDB", 1)
		assert.Equal(t, http.StatusInternalServerError, recorder.Code)
		assert.Contains(t, hook.LastEntry().Message, "update: failed to convert marshal to json")
	})

	t.Run("should return 400 when Content-Type is incorrect", func(t *testing.T) {

		//given
		img := model.Image{
			ID:      fixedID,
			Content: "base64",
			GCP:     []string{"labels", "moods"},
			Status:  false,
		}
		jsonImg, err := json.Marshal(img)
		assert.NoError(t, err)
		req, err := http.NewRequest("PUT", "/v1/images/"+key, nil)
		assert.NoError(t, err)
		vars := map[string]string{
			"id": key,
		}
		req = mux.SetURLVars(req, vars)
		assert.NoError(t, err)
		req.Header.Set("Content-Type", "application/golang")
		recorder := httptest.NewRecorder()
		dbManagerMock := mocks.DBManager{}
		idMock := mocks.IdGenerator{}
		eventBusMock := mocks.EventBus{}
		testSubject := NewHandler(&dbManagerMock, &idMock, &eventBusMock)
		dbManagerMock.On("GetFromDB", key).Return(string(jsonImg), nil)

		//when
		testSubject.Update(recorder, req)

		//then
		assert.Equal(t, http.StatusBadRequest, recorder.Code)
		dbManagerMock.AssertNumberOfCalls(t, "InsertToDB", 0)
		dbManagerMock.AssertNumberOfCalls(t, "GetFromDB", 1)
		assert.Contains(t, recorder.Body.String(), "PUT: invalid content type")
	})

	t.Run("should return 500 when is an error during reading request body", func(t *testing.T) {
		//given
		hook := test.NewGlobal()
		img := model.Image{
			ID:      fixedID,
			Content: "base64",
			GCP:     []string{"labels", "moods"},
			Status:  false,
		}
		jsonImg, err := json.Marshal(img)
		assert.NoError(t, err)
		req, err := http.NewRequest("PUT", "/v1/images/"+key, iotest.ErrReader(errors.New("failed to read")))
		assert.NoError(t, err)
		vars := map[string]string{
			"id": key,
		}
		req = mux.SetURLVars(req, vars)
		assert.NoError(t, err)
		req.Header.Set("Content-Type", "application/json")
		recorder := httptest.NewRecorder()
		dbManagerMock := mocks.DBManager{}
		idMock := mocks.IdGenerator{}
		eventBusMock := mocks.EventBus{}
		testSubject := NewHandler(&dbManagerMock, &idMock, &eventBusMock)
		dbManagerMock.On("GetFromDB", key).Return(string(jsonImg), nil)

		//when
		testSubject.Update(recorder, req)

		//then
		assert.Equal(t, http.StatusInternalServerError, recorder.Code)
		dbManagerMock.AssertNumberOfCalls(t, "InsertToDB", 0)
		dbManagerMock.AssertNumberOfCalls(t, "GetFromDB", 1)
		assert.Contains(t, recorder.Body.String(), "failed to read")
		assert.Contains(t, hook.LastEntry().Message, "update: failed to read request body")
	})

	t.Run("should return 500 when failed to insert data to database", func(t *testing.T) {

		//give
		body := `{` +
			`labels:[labels,moods]` +
			`}`
		hook := test.NewGlobal()
		img := model.Image{
			ID:      fixedID,
			Content: "base64",
			Status:  false,
		}
		imgWithGCP := model.Image{
			ID:      fixedID,
			Content: "base64",
			GCP:     []string{"{labels:[labels,moods]}"},
			Status:  false,
		}
		jsonImg, err := json.Marshal(img)
		jsonImgWithGCP, err := json.Marshal(imgWithGCP)
		assert.NoError(t, err)
		req, err := http.NewRequest("PUT", "/v1/images/"+key, bytes.NewBuffer([]byte(body)))
		vars := map[string]string{
			"id": key,
		}
		req = mux.SetURLVars(req, vars)
		assert.NoError(t, err)
		req.Header.Set("Content-Type", "application/json")
		recorder := httptest.NewRecorder()
		dbManagerMock := mocks.DBManager{}
		idMock := mocks.IdGenerator{}
		eventBusMock := mocks.EventBus{}
		testSubject := NewHandler(&dbManagerMock, &idMock, &eventBusMock)
		dbManagerMock.On("GetFromDB", key).Return(string(jsonImg), nil)
		err = errors.New("failed to insert json to db")
		dbManagerMock.On("InsertToDB", key, string(jsonImgWithGCP)).Return(err)

		//when
		testSubject.Update(recorder, req)

		//then
		assert.Equal(t, http.StatusInternalServerError, recorder.Code)
		dbManagerMock.AssertNumberOfCalls(t, "InsertToDB", 1)
		dbManagerMock.AssertNumberOfCalls(t, "GetFromDB", 1)
		assert.Error(t, err)
		assert.Contains(t, recorder.Body.String(), err.Error())
		assert.Contains(t, hook.LastEntry().Message, err.Error())
	})

	t.Run("should return 200 code whewasupdating data in databse is correct", func(t *testing.T) {

		//given
		body := `{` +
			`labels:[labels,moods]` +
			`}`
		img := model.Image{
			ID:      fixedID,
			Content: "base64",
			Status:  false,
		}
		imgWithGCP := model.Image{
			ID:      fixedID,
			Content: "base64",
			GCP:     []string{"{labels:[labels,moods]}"},
			Status:  false,
		}
		jsonImg, err := json.Marshal(img)
		jsonImgWithGCP, err := json.Marshal(imgWithGCP)
		req, err := http.NewRequest("PUT", "/v1/images/"+fixedID, bytes.NewBuffer([]byte(body)))
		vars := map[string]string{
			"id": fixedID,
		}
		req.Header.Set("Content-Type", "application/json")
		req = mux.SetURLVars(req, vars)
		assert.NoError(t, err)
		recorder := httptest.NewRecorder()
		dbManagerMock := mocks.DBManager{}
		idMock := mocks.IdGenerator{}
		idMock.On("NewID").Return(fixedID, nil)
		eventBusMock := mocks.EventBus{}
		testSubject := NewHandler(&dbManagerMock, &idMock, &eventBusMock)
		dbManagerMock.On("GetFromDB", key).Return(string(jsonImg), nil)
		dbManagerMock.On("InsertToDB", key, string(jsonImgWithGCP)).Return(nil)

		//when
		testSubject.Update(recorder, req)

		//then
		dbManagerMock.AssertNumberOfCalls(t, "GetFromDB", 1)
		dbManagerMock.AssertNumberOfCalls(t, "InsertToDB", 1)
		assert.Equal(t, http.StatusOK, recorder.Code)
	})

}

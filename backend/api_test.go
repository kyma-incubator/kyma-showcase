package main

import (
	"bytes"
	"github.com/gorilla/mux"
	"github.com/kyma-incubator/Kyma-Showcase/mocks"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDBGetAllHandler(t *testing.T) {
	t.Run("should not return error when database is empty", func(t *testing.T) {

		//given
		req, err := http.NewRequest("GET", "v1/images", nil)

		require.NoError(t, err)
		recorder := httptest.NewRecorder()
		dbManagerMock := mocks.DBManager{}
		testSubject := NewHandler(&dbManagerMock)
		dbManagerMock.On("GetAllKeys").Return(nil, nil)

		//when
		testSubject.DBGetAllHandler(recorder, req)

		//then
		dbManagerMock.AssertNumberOfCalls(t, "GetAllKeys",1)
		assert.Equal(t, http.StatusOK, recorder.Code)

	})

	t.Run("should return 500 code when all values are empty", func(t *testing.T) {

		//given
		req, err := http.NewRequest("GET", "v1/images", nil)

		require.NoError(t, err)
		recorder := httptest.NewRecorder()
		dbManagerMock := mocks.DBManager{}
		testSubject := NewHandler(&dbManagerMock)
		dbManagerMock.On("GetAllKeys").Return([]string{"1", "2", "3"}, nil)
		dbManagerMock.On("GetFromDB", "1").Return("", errors.New("value is empty"))
		dbManagerMock.On("GetFromDB", "2").Return("", errors.New("value is empty"))
		dbManagerMock.On("GetFromDB", "3").Return("", errors.New("value is empty"))
		//handler := http.HandleFunc(DBGetAllHandler)

		//when
		testSubject.DBGetAllHandler(recorder, req)

		//then
		dbManagerMock.AssertNumberOfCalls(t, "GetAllKeys",1)
		dbManagerMock.AssertCalled(t, "GetFromDB","1")
		dbManagerMock.AssertNotCalled(t, "GetFromDB","2")
		dbManagerMock.AssertNotCalled(t, "GetFromDB","3")
		dbManagerMock.AssertNumberOfCalls(t, "GetFromDB",1) // 1 poniewaz jest return po pierwszej pustej wartosci
		assert.Equal(t, http.StatusInternalServerError, recorder.Code)

	})

	t.Run("should return 500 code when one of the values is empty", func(t *testing.T) {

		//given
		req, err := http.NewRequest("GET", "v1/images", nil)
		firstReturn := `
			{
				"url":"raccoon.com",
				"gcp":"image.png",
				"img":"image.png"
			}`

		secondReturn := `
			{
				"url":"raccoon.com",
				"gcp":"image2.png",
				"img":"image2.png"
			}
			`

		require.NoError(t, err)
		recorder := httptest.NewRecorder()
		dbManagerMock := mocks.DBManager{}
		testSubject := NewHandler(&dbManagerMock)
		dbManagerMock.On("GetAllKeys").Return([]string{"1", "2", "3"}, nil)
		dbManagerMock.On("GetFromDB", "1").Return(firstReturn, nil)
		dbManagerMock.On("GetFromDB", "2").Return(secondReturn, nil)
		dbManagerMock.On("GetFromDB", "3").Return("", errors.New("value is empty"))
		//handler := http.HandleFunc(DBGetAllHandler)

		//when
		testSubject.DBGetAllHandler(recorder, req)

		//then
		dbManagerMock.AssertNumberOfCalls(t, "GetAllKeys",1)
		dbManagerMock.AssertCalled(t, "GetFromDB","1")
		dbManagerMock.AssertCalled(t, "GetFromDB","2")
		dbManagerMock.AssertCalled(t, "GetFromDB","3")
		dbManagerMock.AssertNumberOfCalls(t, "GetFromDB",3)
		assert.Equal(t, http.StatusInternalServerError, recorder.Code)

	})

	t.Run("should return 500 code when is error during type assertion", func(t *testing.T) {

		//given
		req, err := http.NewRequest("GET", "v1/images", nil)
		require.NoError(t, err)
		recorder := httptest.NewRecorder()
		dbManagerMock := mocks.DBManager{}
		testSubject := NewHandler(&dbManagerMock)
		dbManagerMock.On("GetAllKeys").Return([]string{"1"}, nil)
		dbManagerMock.On("GetFromDB", "1").Return(100, nil)
		
		//when
		testSubject.DBGetAllHandler(recorder,req)
		
		//then
		dbManagerMock.AssertNumberOfCalls(t, "GetAllKeys",1)
		dbManagerMock.AssertCalled(t, "GetFromDB","1")
		dbManagerMock.AssertNumberOfCalls(t, "GetFromDB",1)
		assert.Equal(t,http.StatusInternalServerError,recorder.Code)
	})

	t.Run("should return JSON array compatible with given data ", func(t *testing.T) {

		//given
		expected := `[{` +
			`"url":"raccoon.com",` +
			`"gcp":"image.png",`+
			`"img":"image.png"` +
			`},`+
			`{`+
			`"url":"raccoon.com",`+
			`"gcp":"image2.png",`+
			`"img":"image2.png"`+
			`}]`

		req, err := http.NewRequest("GET", "v1/images", nil)
		require.NoError(t, err)
		recorder := httptest.NewRecorder()
		dbManagerMock := mocks.DBManager{}
		testSubject := NewHandler(&dbManagerMock)
		dbManagerMock.On("GetAllKeys").Return([]string{"1","2"}, nil)

		firstReturn := `
			{
				"url":"raccoon.com",
				"gcp":"image.png",
				"img":"image.png"
			}`

		secondReturn := `
			{
				"url":"raccoon.com",
				"gcp":"image2.png",
				"img":"image2.png"
			}
			`

		dbManagerMock.On("GetFromDB", "1").Return(firstReturn, nil)
		dbManagerMock.On("GetFromDB", "2").Return(secondReturn, nil)

		//when
		testSubject.DBGetAllHandler(recorder,req)

		//then
		dbManagerMock.AssertNumberOfCalls(t, "GetAllKeys",1)
		dbManagerMock.AssertCalled(t, "GetFromDB","1")
		dbManagerMock.AssertCalled(t, "GetFromDB","2")
		dbManagerMock.AssertNumberOfCalls(t, "GetFromDB",2)
		assert.Equal(t,expected,recorder.Body.String())
	})

	t.Run("should return error while is error during unmarshal", func(t *testing.T) {

		//given
		req, err := http.NewRequest("GET", "v1/images", nil)
		require.NoError(t, err)
		recorder := httptest.NewRecorder()
		dbManagerMock := mocks.DBManager{}
		testSubject := NewHandler(&dbManagerMock)
		dbManagerMock.On("GetAllKeys").Return([]string{"1","2"}, nil)

		firstReturn := `
			{
				url":"raccoon.com",
				"gcp":"image.png",
				"img":"image.png"
			}`

		secondReturn := `
			{
				"url":"raccoon.com"
				"gcp":"image2.png",
				"img":"image2.png"
			}
			`

		dbManagerMock.On("GetFromDB", "1").Return(firstReturn, nil)
		dbManagerMock.On("GetFromDB", "2").Return(secondReturn, nil)

		//when
		testSubject.DBGetAllHandler(recorder,req)

		//then
		dbManagerMock.AssertNumberOfCalls(t, "GetAllKeys",1)
		dbManagerMock.AssertCalled(t, "GetFromDB","1")
		dbManagerMock.AssertNotCalled(t, "GetFromDB","2")
		dbManagerMock.AssertNumberOfCalls(t, "GetFromDB",1) //return po pierwszym blednym odczycie
		assert.Equal(t,http.StatusInternalServerError,recorder.Code)
	})

	t.Run("should not return error while everything is ok", func(t *testing.T) {

		//given

		req, err := http.NewRequest("GET", "v1/images", nil)
		require.NoError(t, err)
		recorder := httptest.NewRecorder()
		dbManagerMock := mocks.DBManager{}
		testSubject := NewHandler(&dbManagerMock)
		dbManagerMock.On("GetAllKeys").Return([]string{"1","2"}, nil)

		firstReturn := `
			{
				"url":"raccoon.com",
				"gcp":"image.png",
				"img":"image.png"
			}`

		secondReturn := `
			{
				"url":"raccoon.com",
				"gcp":"image2.png",
				"img":"image2.png"
			}
			`

		dbManagerMock.On("GetFromDB", "1").Return(firstReturn, nil)
		dbManagerMock.On("GetFromDB", "2").Return(secondReturn, nil)

		//when
		testSubject.DBGetAllHandler(recorder,req)

		//then
		dbManagerMock.AssertNumberOfCalls(t, "GetAllKeys",1)
		dbManagerMock.AssertCalled(t, "GetFromDB","1")
		dbManagerMock.AssertCalled(t, "GetFromDB","2")
		dbManagerMock.AssertNumberOfCalls(t, "GetFromDB",2)
		assert.Equal(t,http.StatusOK,recorder.Code)
	})
}

func TestDBPostHandler(t *testing.T) {
	t.Run("should return 400 error when unable to decode json from request", func(t *testing.T) {
		//TODO fix json decoding (all json fields must be compatible with Image struct)
		//given
		var jsonStr = []byte(`{test}`)
		req, err := http.NewRequest("POST", "v1/images/1", bytes.NewBuffer(jsonStr))
		require.NoError(t, err)
		vars := map[string]string{
			"id": "1",
		}
		req = mux.SetURLVars(req, vars)
		recorder := httptest.NewRecorder()
		dbManagerMock := mocks.DBManager{}
		testSubject := NewHandler(&dbManagerMock)

		//when
		testSubject.DBPostHandler(recorder, req)

		//then
		assert.Equal(t, http.StatusBadRequest, recorder.Code)

	})
	t.Run("should return 500 error when unable to insert json to db", func(t *testing.T) {

		//given
		var jsonStr = `{"url":"raccoon.com","gcp":"image.png","img":"image.png"}`
		req, err := http.NewRequest("POST", "/v1/images/1", bytes.NewBuffer([]byte(jsonStr)))
		require.NoError(t, err)
		vars := map[string]string{
			"id": "1",
		}
		req = mux.SetURLVars(req, vars)
		recorder := httptest.NewRecorder()
		dbManagerMock := mocks.DBManager{}
		testSubject := NewHandler(&dbManagerMock)
		dbManagerMock.On("InsertToDB", "1",jsonStr).Return(errors.New("failed to insert json to db"))

		//when
		testSubject.DBPostHandler(recorder, req)

		//then
		assert.Equal(t, http.StatusInternalServerError, recorder.Code)

	})

	//TODO change to id checking
	t.Run("should return 200 when request is correct", func(t *testing.T) {

		//given
		var jsonStr = `{"url":"raccoon.com","gcp":"image.png","img":"image.png"}`
		req, err := http.NewRequest("POST", "/v1/images/1", bytes.NewBuffer([]byte(jsonStr)))
		require.NoError(t, err)
		vars := map[string]string{
			"id": "1",
		}
		req = mux.SetURLVars(req, vars)
		recorder := httptest.NewRecorder()
		dbManagerMock := mocks.DBManager{}
		testSubject := NewHandler(&dbManagerMock)
		dbManagerMock.On("InsertToDB", "1",jsonStr).Return(nil)

		//when
		testSubject.DBPostHandler(recorder, req)

		//then
		assert.Equal(t, http.StatusOK, recorder.Code)

	})

}

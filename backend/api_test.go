package main

import (
	"github.com/kyma-incubator/Kyma-Showcase/mocks"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDBGetAllHandler(t *testing.T) {
	t.Run("should return error when database is empty", func(t *testing.T) {

		//given
		req, err := http.NewRequest("GET", "v1/images", nil)

		require.NoError(t, err)
		recorder := httptest.NewRecorder()
		dbManagerMock := mocks.DBManager{}
		testSubject := NewHandler(&dbManagerMock)
		dbManagerMock.On("GetAllKeys").Return(nil, errors.New("Empty database, returned nil"))

		//when
		testSubject.DBGetAllHandler(recorder, req)

		//then
		assert.Equal(t, http.StatusNotFound, recorder.Code)

	})

	t.Run("should return 404 code when value is empty", func(t *testing.T) {

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
		assert.Equal(t, http.StatusNotFound, recorder.Code)
	})

	//TODO
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
		assert.Equal(t,http.StatusInternalServerError,recorder.Code)
	})
}

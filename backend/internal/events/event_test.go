package events

import (
	"github.com/kyma-incubator/Kyma-Showcase/internal/model"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

const fixedID = "FEA98D88-0669-4FFD-B17A-8F80BB97C381"

func TestSendNewImage(t *testing.T) {
	t.Run("should send event", func(t *testing.T) {

		img := model.Image{}

		handler := func(rw http.ResponseWriter, r *http.Request) {
			rw.WriteHeader(http.StatusNoContent)
		}
		testServer := httptest.NewServer(http.HandlerFunc(handler))
		defer testServer.Close()
		eventHandler := NewEventHandler(testServer.URL)
		err := eventHandler.SendNewImage(fixedID, img)
		assert.NoError(t, err)

	})
	t.Run("should return error when failed to access event's url", func(t *testing.T) {

		img := model.Image{}

		eventHandler := NewEventHandler("notexists.com")
		err := eventHandler.SendNewImage(fixedID, img)
		assert.Error(t, err)

	})
	t.Run("should return error when server returns unexpected status code", func(t *testing.T) {

		img := model.Image{}

		handler := func(rw http.ResponseWriter, r *http.Request) {
			rw.WriteHeader(http.StatusInternalServerError)
		}
		testServer := httptest.NewServer(http.HandlerFunc(handler))
		defer testServer.Close()
		eventHandler := NewEventHandler(testServer.URL)
		err := eventHandler.SendNewImage(fixedID, img)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "sending event returned unexpected status")
	})
}

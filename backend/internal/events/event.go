package events

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/kyma-incubator/Kyma-Showcase/internal/model"
	"github.com/kyma-incubator/Kyma-Showcase/internal/utils"
	"net/http"
	"os"
)

//go:generate mockery --name=EventFactory
// EventFactory defines all event methods
type EventFactory interface {
	NewEvent(id string) model.Event
}

// ImgEvent fulfills the EventFactory interface
type ImgEvent struct {
	Img model.Image
}

// NewEvent defines the values for Event
func (e *ImgEvent) NewEvent(id string) model.Event {
	return model.Event{
		Source:           "kyma-showcase",
		SpecVersion:      "1.0",
		EventTypeVersion: "v1",
		Data:             e.Img.ID,
		Id:               id,
		DataContentType:  "application/json",
		EventType:        os.Getenv("EVENT_TYPE"),
	}
}

// EventHandler for event
type EventHandler struct {
	eventFactory EventFactory
	idGenerator  utils.IdGenerator
}

// NewEventHandler returns handler for event
func NewEventHandler(event EventFactory, idGenerator utils.IdGenerator) EventHandler {
	return EventHandler{
		eventFactory: event,
		idGenerator:  idGenerator,
	}
}

//go:generate mockery --name=EventHelper
// EventHelper helps to mock SendEvent method in tests
type EventHelper interface {
	SendEvent() error
}

// SendEvent creates event and sends a request that triggers lambda
func (e EventHandler) SendEvent() error {
	eventID, err := e.idGenerator.NewID()
	if err != nil {
		err = errors.New("SENDEVENT: nanoid error" + err.Error())
		return err
	}

	event := e.eventFactory.NewEvent(eventID)
	postBody, err := json.Marshal(event)
	if err != nil {
		err = errors.New("SENDEVENT: marshal error" + err.Error())
		return err
	}

	responseBody := bytes.NewBuffer(postBody)
	_, err = http.Post(os.Getenv("EVENT_URL"), "application/cloudevents+json", responseBody)
	if err != nil {
		err = errors.New("SENDEVENT: post sending error" + err.Error())
		return err
	}

	return nil
}

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

// Event defines a struct that triggers lambdas
type Event struct {
	Source           string `json:"source"`
	SpecVersion      string `json:"specversion"`
	EventTypeVersion string `json:"eventtypeversion"`
	Data             string `json:"data"`
	DataContentType  string `json:"datacontenttype"`
	Id               string `json:"id"`
	EventType        string `json:"type"`
}

// NewEvent returns new event
func NewEvent(img model.Image) Event {
	return Event{
		Source:           "kyma-showcase",
		SpecVersion:      "1.0",
		EventTypeVersion: "v1",
		Data:             img.ID,
		DataContentType:  "application/json",
		EventType:        os.Getenv("EVENT_TYPE"),
	}
}

// EventHandler for event
type EventHandler struct {
	event       Event
	idGenerator utils.IdGenerator
}

// NewEventHandler returns handler for event
func NewEventHandler(event Event, idGenerator utils.IdGenerator) EventHandler {
	return EventHandler{
		event:       event,
		idGenerator: idGenerator,
	}
}

// SendEvent creates event and sends a request that triggers lambda
func (e EventHandler) SendEvent() error {
	eventID, err := e.idGenerator.NewID()
	if err != nil {
		err = errors.New("SENDEVENT: nanoid error" + err.Error())
		return err
	}

	e.event.Id = eventID
	postBody, err := json.Marshal(e.event)
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

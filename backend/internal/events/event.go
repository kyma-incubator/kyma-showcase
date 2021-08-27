package events

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/kyma-incubator/Kyma-Showcase/internal/model"
	"net/http"
	"time"
)

const eventType = "sap.kyma.custom.showcase.image.uploaded.v1"

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

// newEvent defines the values for Event
func newEvent(id string, img model.Image) Event {
	return Event{
		Source:           "kyma-showcase",
		SpecVersion:      "1.0",
		EventTypeVersion: "v1",
		Data:             img.ID,
		Id:               id,
		DataContentType:  "application/json",
		EventType:        eventType,
	}
}

// EventHandler for event
type EventHandler struct {
	client *http.Client
	url    string
}

// NewEventHandler returns handler for event
func NewEventHandler(url string) EventHandler {
	return EventHandler{
		client: &http.Client{
			Timeout: 30 * time.Second,
		},
		url: url,
	}
}

// SendNewImage creates event and sends a request that triggers lambda
func (e EventHandler) SendNewImage(id string, img model.Image) error {
	event := newEvent(id, img)
	postBody, err := json.Marshal(event)
	if err != nil {
		err = errors.New("SENDEVENT: marshal error" + err.Error())
		return err
	}

	responseBody := bytes.NewBuffer(postBody)

	request, err := http.NewRequest(http.MethodPost, e.url, responseBody)
	if err != nil {
		return err
	}
	request.Header.Add("Content-Type", "application/cloudevents+json")
	resp, err := e.client.Do(request)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusNoContent {
		err = errors.New(fmt.Sprintf("eventing returned not expected status: %s, %v", resp.Status, resp.StatusCode))
	}

	if err != nil {
		err = errors.New("SENDEVENT: post sending error" + err.Error())
		return err
	}

	return nil
}

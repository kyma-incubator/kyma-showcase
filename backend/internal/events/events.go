package events

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/kyma-incubator/Kyma-Showcase/internal/model"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"net/http"
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

//func NewEvent(img model.Image) (Event, error) {
//	_, err := gonanoid.New()
//	if err != nil {
//		log.Error(err)
//		return Event{}, err
//	}
//	return Event{
//		Source:           "kyma-showcase",
//		SpecVersion:      "1.0",
//		EventTypeVersion: "v1",
//		Data:             "",//img,
//		DataContentType:  "application/json",
//		Id:               "759815c3-b142-48f2-bf18-c6502dc0998f",
//		EventType:        "sap.kyma.custom.commerce.order.created.v1",
//	}, nil
//}

// SendEvent creates and processes request that triggers lambda
func SendEvent(img model.Image) error {
	eventID, err := gonanoid.New()
	if err != nil {
		err = errors.New("SENDEVENT: nanoid error" + err.Error())
		return err
	}
	postBody, err := json.Marshal(Event{
		Source:           "kyma-showcase",
		SpecVersion:      "1.0",
		EventTypeVersion: "v1",
		Data:             img.ID,
		DataContentType:  "application/json",
		Id:               eventID,
		EventType:        "sap.kyma.custom.commerce.order.created.v1",
		//EventType: "sap.kyma.custom.kyma-showcase.order.created.v1",
	})
	if err != nil {
		err = errors.New("SENDEVENT: marshal error" + err.Error())
		return err
	}
	responseBody := bytes.NewBuffer(postBody)
	_, err = http.Post("http://eventing-event-publisher-proxy.kyma-system/publish", "application/cloudevents+json", responseBody)
	if err != nil {
		err = errors.New("SENDEVENT: post sending error" + err.Error())
		return err
	}
	return nil
}

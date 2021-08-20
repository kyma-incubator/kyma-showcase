package events

import (
	"testing"
)

//func (e EventHandler) SendEvent() error {
//	eventID, err := e.idGenerator.NewID()
//	if err != nil {
//		err = errors.New("SENDEVENT: nanoid error" + err.Error())
//		return err
//	}
//
//	e.event.Id = eventID
//	postBody, err := json.Marshal(e.event)
//	if err != nil {
//		err = errors.New("SENDEVENT: marshal error" + err.Error())
//		return err
//	}
//
//	responseBody := bytes.NewBuffer(postBody)
//	_, err = http.Post(os.Getenv("EVENT_URL"), "application/cloudevents+json", responseBody)
//	if err != nil {
//		err = errors.New("SENDEVENT: post sending error" + err.Error())
//		return err
//	}
//
//	return nil
//}

func TestSendEvent(t *testing.T) {
	t.Run("", func(t *testing.T) {
		//given

		//when

		//then
	})

}

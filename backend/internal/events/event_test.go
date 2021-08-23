package events

import (
	"errors"
	eventMocks "github.com/kyma-incubator/Kyma-Showcase/internal/events/mocks"
	"github.com/kyma-incubator/Kyma-Showcase/internal/model"
	utilsMocks "github.com/kyma-incubator/Kyma-Showcase/internal/utils/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

const fixedID = "FEA98D88-0669-4FFD-B17A-8F80BB97C381"

func TestSendEvent(t *testing.T) {
	t.Run("should return nano id error when generating id fails", func(t *testing.T) {
		//given
		idMock := utilsMocks.IdGenerator{}
		idMock.On("NewID").Return("", errors.New("SENDEVENT: nanoid error"))
		eventFactoryMock := eventMocks.EventFactory{}
		testSubject := NewEventHandler(&eventFactoryMock, &idMock)

		//when
		err := testSubject.SendEvent()

		//then
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "SENDEVENT: nanoid error")
	})

	t.Run("should return post error when post fails ", func(t *testing.T) {
		//given
		idMock := utilsMocks.IdGenerator{}
		idMock.On("NewID").Return(fixedID, nil)
		eventFactoryMock := eventMocks.EventFactory{}
		eventFactoryMock.On("NewEvent", fixedID).Return(model.Event{})
		testSubject := NewEventHandler(&eventFactoryMock, &idMock)

		//when
		err := testSubject.SendEvent()

		//then
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "SENDEVENT: post sending error")
	})
}

// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package useCase

import (
	"github.com/CoolCodeTeam/2019_2_CoolCodeMicroServices/utils/models"
	"sync"
)

var (
	lockNotificationUseCaseMockOpenConn    sync.RWMutex
	lockNotificationUseCaseMockSendMessage sync.RWMutex
)

// Ensure, that NotificationUseCaseMock does implement NotificationUseCase.
// If this is not the case, regenerate this file with moq.
var _ NotificationUseCase = &NotificationUseCaseMock{}

// NotificationUseCaseMock is a mock implementation of NotificationUseCase.
//
//     func TestSomethingThatUsesNotificationUseCase(t *testing.T) {
//
//         // make and configure a mocked NotificationUseCase
//         mockedNotificationUseCase := &NotificationUseCaseMock{
//             OpenConnFunc: func(ID uint64) (*models.WebSocketHub, error) {
// 	               panic("mock out the OpenConn method")
//             },
//             SendMessageFunc: func(chatID uint64, message []byte) error {
// 	               panic("mock out the SendMessage method")
//             },
//         }
//
//         // use mockedNotificationUseCase in code that requires NotificationUseCase
//         // and then make assertions.
//
//     }
type NotificationUseCaseMock struct {
	// OpenConnFunc mocks the OpenConn method.
	OpenConnFunc func(ID uint64) (*models.WebSocketHub, error)

	// SendMessageFunc mocks the SendMessage method.
	SendMessageFunc func(chatID uint64, message []byte) error

	// calls tracks calls to the methods.
	calls struct {
		// OpenConn holds details about calls to the OpenConn method.
		OpenConn []struct {
			// ID is the ID argument value.
			ID uint64
		}
		// SendMessage holds details about calls to the SendMessage method.
		SendMessage []struct {
			// ChatID is the chatID argument value.
			ChatID uint64
			// Message is the message argument value.
			Message []byte
		}
	}
}

// OpenConn calls OpenConnFunc.
func (mock *NotificationUseCaseMock) OpenConn(ID uint64) (*models.WebSocketHub, error) {
	if mock.OpenConnFunc == nil {
		panic("NotificationUseCaseMock.OpenConnFunc: method is nil but NotificationUseCase.OpenConn was just called")
	}
	callInfo := struct {
		ID uint64
	}{
		ID: ID,
	}
	lockNotificationUseCaseMockOpenConn.Lock()
	mock.calls.OpenConn = append(mock.calls.OpenConn, callInfo)
	lockNotificationUseCaseMockOpenConn.Unlock()
	return mock.OpenConnFunc(ID)
}

// OpenConnCalls gets all the calls that were made to OpenConn.
// Check the length with:
//     len(mockedNotificationUseCase.OpenConnCalls())
func (mock *NotificationUseCaseMock) OpenConnCalls() []struct {
	ID uint64
} {
	var calls []struct {
		ID uint64
	}
	lockNotificationUseCaseMockOpenConn.RLock()
	calls = mock.calls.OpenConn
	lockNotificationUseCaseMockOpenConn.RUnlock()
	return calls
}

// SendMessage calls SendMessageFunc.
func (mock *NotificationUseCaseMock) SendMessage(chatID uint64, message []byte) error {
	if mock.SendMessageFunc == nil {
		panic("NotificationUseCaseMock.SendMessageFunc: method is nil but NotificationUseCase.SendMessage was just called")
	}
	callInfo := struct {
		ChatID  uint64
		Message []byte
	}{
		ChatID:  chatID,
		Message: message,
	}
	lockNotificationUseCaseMockSendMessage.Lock()
	mock.calls.SendMessage = append(mock.calls.SendMessage, callInfo)
	lockNotificationUseCaseMockSendMessage.Unlock()
	return mock.SendMessageFunc(chatID, message)
}

// SendMessageCalls gets all the calls that were made to SendMessage.
// Check the length with:
//     len(mockedNotificationUseCase.SendMessageCalls())
func (mock *NotificationUseCaseMock) SendMessageCalls() []struct {
	ChatID  uint64
	Message []byte
} {
	var calls []struct {
		ChatID  uint64
		Message []byte
	}
	lockNotificationUseCaseMockSendMessage.RLock()
	calls = mock.calls.SendMessage
	lockNotificationUseCaseMockSendMessage.RUnlock()
	return calls
}

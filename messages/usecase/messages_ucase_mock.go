// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package useCase

import (
	"github.com/go-park-mail-ru/2019_2_CoolCodeMicroServices/utils/models"
	"sync"
)

var (
	lockMessagesUseCaseMockDeleteMessage        sync.RWMutex
	lockMessagesUseCaseMockEditMessage          sync.RWMutex
	lockMessagesUseCaseMockFindMessages         sync.RWMutex
	lockMessagesUseCaseMockGetChannelMessages   sync.RWMutex
	lockMessagesUseCaseMockGetChatMessages      sync.RWMutex
	lockMessagesUseCaseMockGetMessageByID       sync.RWMutex
	lockMessagesUseCaseMockHideMessageForAuthor sync.RWMutex
	lockMessagesUseCaseMockSaveChannelMessage   sync.RWMutex
	lockMessagesUseCaseMockSaveChatMessage      sync.RWMutex
)

// Ensure, that MessagesUseCaseMock does implement MessagesUseCase.
// If this is not the case, regenerate this file with moq.
var _ MessagesUseCase = &MessagesUseCaseMock{}

// MessagesUseCaseMock is a mock implementation of MessagesUseCase.
//
//     func TestSomethingThatUsesMessagesUseCase(t *testing.T) {
//
//         // make and configure a mocked MessagesUseCase
//         mockedMessagesUseCase := &MessagesUseCaseMock{
//             DeleteMessageFunc: func(messageID uint64, userID uint64) error {
// 	               panic("mock out the DeleteMessage method")
//             },
//             EditMessageFunc: func(message *models.Message, userID uint64) error {
// 	               panic("mock out the EditMessage method")
//             },
//             FindMessagesFunc: func(findString string, ID uint64) (models.Messages, error) {
// 	               panic("mock out the FindMessages method")
//             },
//             GetChannelMessagesFunc: func(channelID uint64, userID uint64) (models.Messages, error) {
// 	               panic("mock out the GetChannelMessages method")
//             },
//             GetChatMessagesFunc: func(chatID uint64, userID uint64) (models.Messages, error) {
// 	               panic("mock out the GetChatMessages method")
//             },
//             GetMessageByIDFunc: func(messageID uint64) (*models.Message, error) {
// 	               panic("mock out the GetMessageByID method")
//             },
//             HideMessageForAuthorFunc: func(messageID uint64, userID uint64) error {
// 	               panic("mock out the HideMessageForAuthor method")
//             },
//             SaveChannelMessageFunc: func(message *models.Message) (uint64, error) {
// 	               panic("mock out the SaveChannelMessage method")
//             },
//             SaveChatMessageFunc: func(message *models.Message) (uint64, error) {
// 	               panic("mock out the SaveChatMessage method")
//             },
//         }
//
//         // use mockedMessagesUseCase in code that requires MessagesUseCase
//         // and then make assertions.
//
//     }
type MessagesUseCaseMock struct {
	// DeleteMessageFunc mocks the DeleteMessage method.
	DeleteMessageFunc func(messageID uint64, userID uint64) error

	// EditMessageFunc mocks the EditMessage method.
	EditMessageFunc func(message *models.Message, userID uint64) error

	// FindMessagesFunc mocks the FindMessages method.
	FindMessagesFunc func(findString string, ID uint64) (models.Messages, error)

	// GetChannelMessagesFunc mocks the GetChannelMessages method.
	GetChannelMessagesFunc func(channelID uint64, userID uint64) (models.Messages, error)

	// GetChatMessagesFunc mocks the GetChatMessages method.
	GetChatMessagesFunc func(chatID uint64, userID uint64) (models.Messages, error)

	// GetMessageByIDFunc mocks the GetMessageByID method.
	GetMessageByIDFunc func(messageID uint64) (*models.Message, error)

	// HideMessageForAuthorFunc mocks the HideMessageForAuthor method.
	HideMessageForAuthorFunc func(messageID uint64, userID uint64) error

	// SaveChannelMessageFunc mocks the SaveChannelMessage method.
	SaveChannelMessageFunc func(message *models.Message) (uint64, error)

	// SaveChatMessageFunc mocks the SaveChatMessage method.
	SaveChatMessageFunc func(message *models.Message) (uint64, error)

	// calls tracks calls to the methods.
	calls struct {
		// DeleteMessage holds details about calls to the DeleteMessage method.
		DeleteMessage []struct {
			// MessageID is the messageID argument value.
			MessageID uint64
			// UserID is the userID argument value.
			UserID uint64
		}
		// EditMessage holds details about calls to the EditMessage method.
		EditMessage []struct {
			// Message is the message argument value.
			Message *models.Message
			// UserID is the userID argument value.
			UserID uint64
		}
		// FindMessages holds details about calls to the FindMessages method.
		FindMessages []struct {
			// FindString is the findString argument value.
			FindString string
			// ID is the ID argument value.
			ID uint64
		}
		// GetChannelMessages holds details about calls to the GetChannelMessages method.
		GetChannelMessages []struct {
			// ChannelID is the channelID argument value.
			ChannelID uint64
			// UserID is the userID argument value.
			UserID uint64
		}
		// GetChatMessages holds details about calls to the GetChatMessages method.
		GetChatMessages []struct {
			// ChatID is the chatID argument value.
			ChatID uint64
			// UserID is the userID argument value.
			UserID uint64
		}
		// GetMessageByID holds details about calls to the GetMessageByID method.
		GetMessageByID []struct {
			// MessageID is the messageID argument value.
			MessageID uint64
		}
		// HideMessageForAuthor holds details about calls to the HideMessageForAuthor method.
		HideMessageForAuthor []struct {
			// MessageID is the messageID argument value.
			MessageID uint64
			// UserID is the userID argument value.
			UserID uint64
		}
		// SaveChannelMessage holds details about calls to the SaveChannelMessage method.
		SaveChannelMessage []struct {
			// Message is the message argument value.
			Message *models.Message
		}
		// SaveChatMessage holds details about calls to the SaveChatMessage method.
		SaveChatMessage []struct {
			// Message is the message argument value.
			Message *models.Message
		}
	}
}

// DeleteMessage calls DeleteMessageFunc.
func (mock *MessagesUseCaseMock) DeleteMessage(messageID uint64, userID uint64) error {
	if mock.DeleteMessageFunc == nil {
		panic("MessagesUseCaseMock.DeleteMessageFunc: method is nil but MessagesUseCase.DeleteMessage was just called")
	}
	callInfo := struct {
		MessageID uint64
		UserID    uint64
	}{
		MessageID: messageID,
		UserID:    userID,
	}
	lockMessagesUseCaseMockDeleteMessage.Lock()
	mock.calls.DeleteMessage = append(mock.calls.DeleteMessage, callInfo)
	lockMessagesUseCaseMockDeleteMessage.Unlock()
	return mock.DeleteMessageFunc(messageID, userID)
}

// DeleteMessageCalls gets all the calls that were made to DeleteMessage.
// Check the length with:
//     len(mockedMessagesUseCase.DeleteMessageCalls())
func (mock *MessagesUseCaseMock) DeleteMessageCalls() []struct {
	MessageID uint64
	UserID    uint64
} {
	var calls []struct {
		MessageID uint64
		UserID    uint64
	}
	lockMessagesUseCaseMockDeleteMessage.RLock()
	calls = mock.calls.DeleteMessage
	lockMessagesUseCaseMockDeleteMessage.RUnlock()
	return calls
}

// EditMessage calls EditMessageFunc.
func (mock *MessagesUseCaseMock) EditMessage(message *models.Message, userID uint64) error {
	if mock.EditMessageFunc == nil {
		panic("MessagesUseCaseMock.EditMessageFunc: method is nil but MessagesUseCase.EditMessage was just called")
	}
	callInfo := struct {
		Message *models.Message
		UserID  uint64
	}{
		Message: message,
		UserID:  userID,
	}
	lockMessagesUseCaseMockEditMessage.Lock()
	mock.calls.EditMessage = append(mock.calls.EditMessage, callInfo)
	lockMessagesUseCaseMockEditMessage.Unlock()
	return mock.EditMessageFunc(message, userID)
}

// EditMessageCalls gets all the calls that were made to EditMessage.
// Check the length with:
//     len(mockedMessagesUseCase.EditMessageCalls())
func (mock *MessagesUseCaseMock) EditMessageCalls() []struct {
	Message *models.Message
	UserID  uint64
} {
	var calls []struct {
		Message *models.Message
		UserID  uint64
	}
	lockMessagesUseCaseMockEditMessage.RLock()
	calls = mock.calls.EditMessage
	lockMessagesUseCaseMockEditMessage.RUnlock()
	return calls
}

// FindMessages calls FindMessagesFunc.
func (mock *MessagesUseCaseMock) FindMessages(findString string, ID uint64) (models.Messages, error) {
	if mock.FindMessagesFunc == nil {
		panic("MessagesUseCaseMock.FindMessagesFunc: method is nil but MessagesUseCase.FindMessages was just called")
	}
	callInfo := struct {
		FindString string
		ID         uint64
	}{
		FindString: findString,
		ID:         ID,
	}
	lockMessagesUseCaseMockFindMessages.Lock()
	mock.calls.FindMessages = append(mock.calls.FindMessages, callInfo)
	lockMessagesUseCaseMockFindMessages.Unlock()
	return mock.FindMessagesFunc(findString, ID)
}

// FindMessagesCalls gets all the calls that were made to FindMessages.
// Check the length with:
//     len(mockedMessagesUseCase.FindMessagesCalls())
func (mock *MessagesUseCaseMock) FindMessagesCalls() []struct {
	FindString string
	ID         uint64
} {
	var calls []struct {
		FindString string
		ID         uint64
	}
	lockMessagesUseCaseMockFindMessages.RLock()
	calls = mock.calls.FindMessages
	lockMessagesUseCaseMockFindMessages.RUnlock()
	return calls
}

// GetChannelMessages calls GetChannelMessagesFunc.
func (mock *MessagesUseCaseMock) GetChannelMessages(channelID uint64, userID uint64) (models.Messages, error) {
	if mock.GetChannelMessagesFunc == nil {
		panic("MessagesUseCaseMock.GetChannelMessagesFunc: method is nil but MessagesUseCase.GetChannelMessages was just called")
	}
	callInfo := struct {
		ChannelID uint64
		UserID    uint64
	}{
		ChannelID: channelID,
		UserID:    userID,
	}
	lockMessagesUseCaseMockGetChannelMessages.Lock()
	mock.calls.GetChannelMessages = append(mock.calls.GetChannelMessages, callInfo)
	lockMessagesUseCaseMockGetChannelMessages.Unlock()
	return mock.GetChannelMessagesFunc(channelID, userID)
}

// GetChannelMessagesCalls gets all the calls that were made to GetChannelMessages.
// Check the length with:
//     len(mockedMessagesUseCase.GetChannelMessagesCalls())
func (mock *MessagesUseCaseMock) GetChannelMessagesCalls() []struct {
	ChannelID uint64
	UserID    uint64
} {
	var calls []struct {
		ChannelID uint64
		UserID    uint64
	}
	lockMessagesUseCaseMockGetChannelMessages.RLock()
	calls = mock.calls.GetChannelMessages
	lockMessagesUseCaseMockGetChannelMessages.RUnlock()
	return calls
}

// GetChatMessages calls GetChatMessagesFunc.
func (mock *MessagesUseCaseMock) GetChatMessages(chatID uint64, userID uint64) (models.Messages, error) {
	if mock.GetChatMessagesFunc == nil {
		panic("MessagesUseCaseMock.GetChatMessagesFunc: method is nil but MessagesUseCase.GetChatMessages was just called")
	}
	callInfo := struct {
		ChatID uint64
		UserID uint64
	}{
		ChatID: chatID,
		UserID: userID,
	}
	lockMessagesUseCaseMockGetChatMessages.Lock()
	mock.calls.GetChatMessages = append(mock.calls.GetChatMessages, callInfo)
	lockMessagesUseCaseMockGetChatMessages.Unlock()
	return mock.GetChatMessagesFunc(chatID, userID)
}

// GetChatMessagesCalls gets all the calls that were made to GetChatMessages.
// Check the length with:
//     len(mockedMessagesUseCase.GetChatMessagesCalls())
func (mock *MessagesUseCaseMock) GetChatMessagesCalls() []struct {
	ChatID uint64
	UserID uint64
} {
	var calls []struct {
		ChatID uint64
		UserID uint64
	}
	lockMessagesUseCaseMockGetChatMessages.RLock()
	calls = mock.calls.GetChatMessages
	lockMessagesUseCaseMockGetChatMessages.RUnlock()
	return calls
}

// GetMessageByID calls GetMessageByIDFunc.
func (mock *MessagesUseCaseMock) GetMessageByID(messageID uint64) (*models.Message, error) {
	if mock.GetMessageByIDFunc == nil {
		panic("MessagesUseCaseMock.GetMessageByIDFunc: method is nil but MessagesUseCase.GetMessageByID was just called")
	}
	callInfo := struct {
		MessageID uint64
	}{
		MessageID: messageID,
	}
	lockMessagesUseCaseMockGetMessageByID.Lock()
	mock.calls.GetMessageByID = append(mock.calls.GetMessageByID, callInfo)
	lockMessagesUseCaseMockGetMessageByID.Unlock()
	return mock.GetMessageByIDFunc(messageID)
}

// GetMessageByIDCalls gets all the calls that were made to GetMessageByID.
// Check the length with:
//     len(mockedMessagesUseCase.GetMessageByIDCalls())
func (mock *MessagesUseCaseMock) GetMessageByIDCalls() []struct {
	MessageID uint64
} {
	var calls []struct {
		MessageID uint64
	}
	lockMessagesUseCaseMockGetMessageByID.RLock()
	calls = mock.calls.GetMessageByID
	lockMessagesUseCaseMockGetMessageByID.RUnlock()
	return calls
}

// HideMessageForAuthor calls HideMessageForAuthorFunc.
func (mock *MessagesUseCaseMock) HideMessageForAuthor(messageID uint64, userID uint64) error {
	if mock.HideMessageForAuthorFunc == nil {
		panic("MessagesUseCaseMock.HideMessageForAuthorFunc: method is nil but MessagesUseCase.HideMessageForAuthor was just called")
	}
	callInfo := struct {
		MessageID uint64
		UserID    uint64
	}{
		MessageID: messageID,
		UserID:    userID,
	}
	lockMessagesUseCaseMockHideMessageForAuthor.Lock()
	mock.calls.HideMessageForAuthor = append(mock.calls.HideMessageForAuthor, callInfo)
	lockMessagesUseCaseMockHideMessageForAuthor.Unlock()
	return mock.HideMessageForAuthorFunc(messageID, userID)
}

// HideMessageForAuthorCalls gets all the calls that were made to HideMessageForAuthor.
// Check the length with:
//     len(mockedMessagesUseCase.HideMessageForAuthorCalls())
func (mock *MessagesUseCaseMock) HideMessageForAuthorCalls() []struct {
	MessageID uint64
	UserID    uint64
} {
	var calls []struct {
		MessageID uint64
		UserID    uint64
	}
	lockMessagesUseCaseMockHideMessageForAuthor.RLock()
	calls = mock.calls.HideMessageForAuthor
	lockMessagesUseCaseMockHideMessageForAuthor.RUnlock()
	return calls
}

// SaveChannelMessage calls SaveChannelMessageFunc.
func (mock *MessagesUseCaseMock) SaveChannelMessage(message *models.Message) (uint64, error) {
	if mock.SaveChannelMessageFunc == nil {
		panic("MessagesUseCaseMock.SaveChannelMessageFunc: method is nil but MessagesUseCase.SaveChannelMessage was just called")
	}
	callInfo := struct {
		Message *models.Message
	}{
		Message: message,
	}
	lockMessagesUseCaseMockSaveChannelMessage.Lock()
	mock.calls.SaveChannelMessage = append(mock.calls.SaveChannelMessage, callInfo)
	lockMessagesUseCaseMockSaveChannelMessage.Unlock()
	return mock.SaveChannelMessageFunc(message)
}

// SaveChannelMessageCalls gets all the calls that were made to SaveChannelMessage.
// Check the length with:
//     len(mockedMessagesUseCase.SaveChannelMessageCalls())
func (mock *MessagesUseCaseMock) SaveChannelMessageCalls() []struct {
	Message *models.Message
} {
	var calls []struct {
		Message *models.Message
	}
	lockMessagesUseCaseMockSaveChannelMessage.RLock()
	calls = mock.calls.SaveChannelMessage
	lockMessagesUseCaseMockSaveChannelMessage.RUnlock()
	return calls
}

// SaveChatMessage calls SaveChatMessageFunc.
func (mock *MessagesUseCaseMock) SaveChatMessage(message *models.Message) (uint64, error) {
	if mock.SaveChatMessageFunc == nil {
		panic("MessagesUseCaseMock.SaveChatMessageFunc: method is nil but MessagesUseCase.SaveChatMessage was just called")
	}
	callInfo := struct {
		Message *models.Message
	}{
		Message: message,
	}
	lockMessagesUseCaseMockSaveChatMessage.Lock()
	mock.calls.SaveChatMessage = append(mock.calls.SaveChatMessage, callInfo)
	lockMessagesUseCaseMockSaveChatMessage.Unlock()
	return mock.SaveChatMessageFunc(message)
}

// SaveChatMessageCalls gets all the calls that were made to SaveChatMessage.
// Check the length with:
//     len(mockedMessagesUseCase.SaveChatMessageCalls())
func (mock *MessagesUseCaseMock) SaveChatMessageCalls() []struct {
	Message *models.Message
} {
	var calls []struct {
		Message *models.Message
	}
	lockMessagesUseCaseMockSaveChatMessage.RLock()
	calls = mock.calls.SaveChatMessage
	lockMessagesUseCaseMockSaveChatMessage.RUnlock()
	return calls
}
// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package repository

import (
	"github.com/CoolCodeTeam/2019_2_CoolCodeMicroServices/utils/models"
	"sync"
)

var (
	lockMessageRepositoryMockFindMessages         sync.RWMutex
	lockMessageRepositoryMockGetMessageByID       sync.RWMutex
	lockMessageRepositoryMockGetMessagesByChatID  sync.RWMutex
	lockMessageRepositoryMockHideMessageForAuthor sync.RWMutex
	lockMessageRepositoryMockLike                 sync.RWMutex
	lockMessageRepositoryMockPutMessage           sync.RWMutex
	lockMessageRepositoryMockRemoveMessage        sync.RWMutex
	lockMessageRepositoryMockUpdateMessage        sync.RWMutex
)

// Ensure, that MessageRepositoryMock does implement MessageRepository.
// If this is not the case, regenerate this file with moq.
var _ MessageRepository = &MessageRepositoryMock{}

// MessageRepositoryMock is a mock implementation of MessageRepository.
//
//     func TestSomethingThatUsesMessageRepository(t *testing.T) {
//
//         // make and configure a mocked MessageRepository
//         mockedMessageRepository := &MessageRepositoryMock{
//             FindMessagesFunc: func(s string) (models.Messages, error) {
// 	               panic("mock out the FindMessages method")
//             },
//             GetMessageByIDFunc: func(messageID uint64) (*models.Message, error) {
// 	               panic("mock out the GetMessageByID method")
//             },
//             GetMessagesByChatIDFunc: func(chatID uint64) (models.Messages, error) {
// 	               panic("mock out the GetMessagesByChatID method")
//             },
//             HideMessageForAuthorFunc: func(userID uint64) error {
// 	               panic("mock out the HideMessageForAuthor method")
//             },
//             LikeFunc: func(messageID uint64) error {
// 	               panic("mock out the Like method")
//             },
//             PutMessageFunc: func(message *models.Message) (uint64, error) {
// 	               panic("mock out the PutMessage method")
//             },
//             RemoveMessageFunc: func(messageID uint64) error {
// 	               panic("mock out the RemoveMessage method")
//             },
//             UpdateMessageFunc: func(message *models.Message) error {
// 	               panic("mock out the UpdateMessage method")
//             },
//         }
//
//         // use mockedMessageRepository in code that requires MessageRepository
//         // and then make assertions.
//
//     }
type MessageRepositoryMock struct {
	// FindMessagesFunc mocks the FindMessages method.
	FindMessagesFunc func(s string) (models.Messages, error)

	// GetMessageByIDFunc mocks the GetMessageByID method.
	GetMessageByIDFunc func(messageID uint64) (*models.Message, error)

	// GetMessagesByChatIDFunc mocks the GetMessagesByChatID method.
	GetMessagesByChatIDFunc func(chatID uint64) (models.Messages, error)

	// HideMessageForAuthorFunc mocks the HideMessageForAuthor method.
	HideMessageForAuthorFunc func(userID uint64) error

	// LikeFunc mocks the Like method.
	LikeFunc func(messageID uint64) error

	// PutMessageFunc mocks the PutMessage method.
	PutMessageFunc func(message *models.Message) (uint64, error)

	// RemoveMessageFunc mocks the RemoveMessage method.
	RemoveMessageFunc func(messageID uint64) error

	// UpdateMessageFunc mocks the UpdateMessage method.
	UpdateMessageFunc func(message *models.Message) error

	// calls tracks calls to the methods.
	calls struct {
		// FindMessages holds details about calls to the FindMessages method.
		FindMessages []struct {
			// S is the s argument value.
			S string
		}
		// GetMessageByID holds details about calls to the GetMessageByID method.
		GetMessageByID []struct {
			// MessageID is the messageID argument value.
			MessageID uint64
		}
		// GetMessagesByChatID holds details about calls to the GetMessagesByChatID method.
		GetMessagesByChatID []struct {
			// ChatID is the chatID argument value.
			ChatID uint64
		}
		// HideMessageForAuthor holds details about calls to the HideMessageForAuthor method.
		HideMessageForAuthor []struct {
			// UserID is the userID argument value.
			UserID uint64
		}
		// Like holds details about calls to the Like method.
		Like []struct {
			// MessageID is the messageID argument value.
			MessageID uint64
		}
		// PutMessage holds details about calls to the PutMessage method.
		PutMessage []struct {
			// Message is the message argument value.
			Message *models.Message
		}
		// RemoveMessage holds details about calls to the RemoveMessage method.
		RemoveMessage []struct {
			// MessageID is the messageID argument value.
			MessageID uint64
		}
		// UpdateMessage holds details about calls to the UpdateMessage method.
		UpdateMessage []struct {
			// Message is the message argument value.
			Message *models.Message
		}
	}
}

// FindMessages calls FindMessagesFunc.
func (mock *MessageRepositoryMock) FindMessages(s string) (models.Messages, error) {
	if mock.FindMessagesFunc == nil {
		panic("MessageRepositoryMock.FindMessagesFunc: method is nil but MessageRepository.FindMessages was just called")
	}
	callInfo := struct {
		S string
	}{
		S: s,
	}
	lockMessageRepositoryMockFindMessages.Lock()
	mock.calls.FindMessages = append(mock.calls.FindMessages, callInfo)
	lockMessageRepositoryMockFindMessages.Unlock()
	return mock.FindMessagesFunc(s)
}

// FindMessagesCalls gets all the calls that were made to FindMessages.
// Check the length with:
//     len(mockedMessageRepository.FindMessagesCalls())
func (mock *MessageRepositoryMock) FindMessagesCalls() []struct {
	S string
} {
	var calls []struct {
		S string
	}
	lockMessageRepositoryMockFindMessages.RLock()
	calls = mock.calls.FindMessages
	lockMessageRepositoryMockFindMessages.RUnlock()
	return calls
}

// GetMessageByID calls GetMessageByIDFunc.
func (mock *MessageRepositoryMock) GetMessageByID(messageID uint64) (*models.Message, error) {
	if mock.GetMessageByIDFunc == nil {
		panic("MessageRepositoryMock.GetMessageByIDFunc: method is nil but MessageRepository.GetMessageByID was just called")
	}
	callInfo := struct {
		MessageID uint64
	}{
		MessageID: messageID,
	}
	lockMessageRepositoryMockGetMessageByID.Lock()
	mock.calls.GetMessageByID = append(mock.calls.GetMessageByID, callInfo)
	lockMessageRepositoryMockGetMessageByID.Unlock()
	return mock.GetMessageByIDFunc(messageID)
}

// GetMessageByIDCalls gets all the calls that were made to GetMessageByID.
// Check the length with:
//     len(mockedMessageRepository.GetMessageByIDCalls())
func (mock *MessageRepositoryMock) GetMessageByIDCalls() []struct {
	MessageID uint64
} {
	var calls []struct {
		MessageID uint64
	}
	lockMessageRepositoryMockGetMessageByID.RLock()
	calls = mock.calls.GetMessageByID
	lockMessageRepositoryMockGetMessageByID.RUnlock()
	return calls
}

// GetMessagesByChatID calls GetMessagesByChatIDFunc.
func (mock *MessageRepositoryMock) GetMessagesByChatID(chatID uint64) (models.Messages, error) {
	if mock.GetMessagesByChatIDFunc == nil {
		panic("MessageRepositoryMock.GetMessagesByChatIDFunc: method is nil but MessageRepository.GetMessagesByChatID was just called")
	}
	callInfo := struct {
		ChatID uint64
	}{
		ChatID: chatID,
	}
	lockMessageRepositoryMockGetMessagesByChatID.Lock()
	mock.calls.GetMessagesByChatID = append(mock.calls.GetMessagesByChatID, callInfo)
	lockMessageRepositoryMockGetMessagesByChatID.Unlock()
	return mock.GetMessagesByChatIDFunc(chatID)
}

// GetMessagesByChatIDCalls gets all the calls that were made to GetMessagesByChatID.
// Check the length with:
//     len(mockedMessageRepository.GetMessagesByChatIDCalls())
func (mock *MessageRepositoryMock) GetMessagesByChatIDCalls() []struct {
	ChatID uint64
} {
	var calls []struct {
		ChatID uint64
	}
	lockMessageRepositoryMockGetMessagesByChatID.RLock()
	calls = mock.calls.GetMessagesByChatID
	lockMessageRepositoryMockGetMessagesByChatID.RUnlock()
	return calls
}

// HideMessageForAuthor calls HideMessageForAuthorFunc.
func (mock *MessageRepositoryMock) HideMessageForAuthor(userID uint64) error {
	if mock.HideMessageForAuthorFunc == nil {
		panic("MessageRepositoryMock.HideMessageForAuthorFunc: method is nil but MessageRepository.HideMessageForAuthor was just called")
	}
	callInfo := struct {
		UserID uint64
	}{
		UserID: userID,
	}
	lockMessageRepositoryMockHideMessageForAuthor.Lock()
	mock.calls.HideMessageForAuthor = append(mock.calls.HideMessageForAuthor, callInfo)
	lockMessageRepositoryMockHideMessageForAuthor.Unlock()
	return mock.HideMessageForAuthorFunc(userID)
}

// HideMessageForAuthorCalls gets all the calls that were made to HideMessageForAuthor.
// Check the length with:
//     len(mockedMessageRepository.HideMessageForAuthorCalls())
func (mock *MessageRepositoryMock) HideMessageForAuthorCalls() []struct {
	UserID uint64
} {
	var calls []struct {
		UserID uint64
	}
	lockMessageRepositoryMockHideMessageForAuthor.RLock()
	calls = mock.calls.HideMessageForAuthor
	lockMessageRepositoryMockHideMessageForAuthor.RUnlock()
	return calls
}

// Like calls LikeFunc.
func (mock *MessageRepositoryMock) Like(messageID uint64) error {
	if mock.LikeFunc == nil {
		panic("MessageRepositoryMock.LikeFunc: method is nil but MessageRepository.Like was just called")
	}
	callInfo := struct {
		MessageID uint64
	}{
		MessageID: messageID,
	}
	lockMessageRepositoryMockLike.Lock()
	mock.calls.Like = append(mock.calls.Like, callInfo)
	lockMessageRepositoryMockLike.Unlock()
	return mock.LikeFunc(messageID)
}

// LikeCalls gets all the calls that were made to Like.
// Check the length with:
//     len(mockedMessageRepository.LikeCalls())
func (mock *MessageRepositoryMock) LikeCalls() []struct {
	MessageID uint64
} {
	var calls []struct {
		MessageID uint64
	}
	lockMessageRepositoryMockLike.RLock()
	calls = mock.calls.Like
	lockMessageRepositoryMockLike.RUnlock()
	return calls
}

// PutMessage calls PutMessageFunc.
func (mock *MessageRepositoryMock) PutMessage(message *models.Message) (uint64, error) {
	if mock.PutMessageFunc == nil {
		panic("MessageRepositoryMock.PutMessageFunc: method is nil but MessageRepository.PutMessage was just called")
	}
	callInfo := struct {
		Message *models.Message
	}{
		Message: message,
	}
	lockMessageRepositoryMockPutMessage.Lock()
	mock.calls.PutMessage = append(mock.calls.PutMessage, callInfo)
	lockMessageRepositoryMockPutMessage.Unlock()
	return mock.PutMessageFunc(message)
}

// PutMessageCalls gets all the calls that were made to PutMessage.
// Check the length with:
//     len(mockedMessageRepository.PutMessageCalls())
func (mock *MessageRepositoryMock) PutMessageCalls() []struct {
	Message *models.Message
} {
	var calls []struct {
		Message *models.Message
	}
	lockMessageRepositoryMockPutMessage.RLock()
	calls = mock.calls.PutMessage
	lockMessageRepositoryMockPutMessage.RUnlock()
	return calls
}

// RemoveMessage calls RemoveMessageFunc.
func (mock *MessageRepositoryMock) RemoveMessage(messageID uint64) error {
	if mock.RemoveMessageFunc == nil {
		panic("MessageRepositoryMock.RemoveMessageFunc: method is nil but MessageRepository.RemoveMessage was just called")
	}
	callInfo := struct {
		MessageID uint64
	}{
		MessageID: messageID,
	}
	lockMessageRepositoryMockRemoveMessage.Lock()
	mock.calls.RemoveMessage = append(mock.calls.RemoveMessage, callInfo)
	lockMessageRepositoryMockRemoveMessage.Unlock()
	return mock.RemoveMessageFunc(messageID)
}

// RemoveMessageCalls gets all the calls that were made to RemoveMessage.
// Check the length with:
//     len(mockedMessageRepository.RemoveMessageCalls())
func (mock *MessageRepositoryMock) RemoveMessageCalls() []struct {
	MessageID uint64
} {
	var calls []struct {
		MessageID uint64
	}
	lockMessageRepositoryMockRemoveMessage.RLock()
	calls = mock.calls.RemoveMessage
	lockMessageRepositoryMockRemoveMessage.RUnlock()
	return calls
}

// UpdateMessage calls UpdateMessageFunc.
func (mock *MessageRepositoryMock) UpdateMessage(message *models.Message) error {
	if mock.UpdateMessageFunc == nil {
		panic("MessageRepositoryMock.UpdateMessageFunc: method is nil but MessageRepository.UpdateMessage was just called")
	}
	callInfo := struct {
		Message *models.Message
	}{
		Message: message,
	}
	lockMessageRepositoryMockUpdateMessage.Lock()
	mock.calls.UpdateMessage = append(mock.calls.UpdateMessage, callInfo)
	lockMessageRepositoryMockUpdateMessage.Unlock()
	return mock.UpdateMessageFunc(message)
}

// UpdateMessageCalls gets all the calls that were made to UpdateMessage.
// Check the length with:
//     len(mockedMessageRepository.UpdateMessageCalls())
func (mock *MessageRepositoryMock) UpdateMessageCalls() []struct {
	Message *models.Message
} {
	var calls []struct {
		Message *models.Message
	}
	lockMessageRepositoryMockUpdateMessage.RLock()
	calls = mock.calls.UpdateMessage
	lockMessageRepositoryMockUpdateMessage.RUnlock()
	return calls
}

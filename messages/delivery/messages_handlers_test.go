package delivery

import (
	"errors"
	"github.com/go-park-mail-ru/2019_2_CoolCodeMicroServices/utils"
	"github.com/go-park-mail-ru/2019_2_CoolCodeMicroServices/utils/models"

	messages "github.com/go-park-mail-ru/2019_2_CoolCodeMicroServices/messages/usecase"
	notification "github.com/go-park-mail-ru/2019_2_CoolCodeMicroServices/notifications/usecase"
	users "github.com/go-park-mail-ru/2019_2_CoolCodeMicroServices/users/usecase"

	middleware "github.com/go-park-mail-ru/2019_2_CoolCodeMicroServices/utils/middlwares"
	"github.com/sirupsen/logrus"
	"github.com/steinfletcher/apitest"
	"io/ioutil"
	"net/http"
	"testing"
)

var messageApi MessageHandlersImpl
var middlware = middleware.HandlersMiddlwares{}

type MessageTestCase struct {
	name         string
	Body         string
	SessionID    string
	Headers      map[string]string
	Method       string
	URL          string
	Response     string
	StatusCode   int
	Handler      func(w http.ResponseWriter,r *http.Request)
	Messages     messages.MessagesUseCase
	Notification notification.NotificationUseCase
	Users        users.UsersUseCase
	utils        utils.HandlersUtils
}

func runTableMessageAPITests(t *testing.T, cases []*MessageTestCase) {
	for _, c := range cases {
		runMessageAPITest(t, c)
	}
}

func runMessageAPITest(t *testing.T, testCase *MessageTestCase) {
	t.Run(testCase.name, func(t *testing.T) {
		if testCase.Messages != nil {
			messageApi.Messages = testCase.Messages
		}
		if testCase.Notification != nil {
			messageApi.Notifications = testCase.Notification
		}
		if testCase.Users != nil {
			messageApi.Users = testCase.Users
			middlware.Users=testCase.Users
		}
		apitest.New().
			Handler(middlware.AuthMiddleware(testCase.Handler)).
			Method(testCase.Method).
			Headers(testCase.Headers).
			Cookie("session_id", "test").
			URL(testCase.URL).
			Body(testCase.Body).
			Expect(t).
			Status(testCase.StatusCode).End()
	})
}

func init() {
	emptyLogger := logrus.New()
	emptyLogger.Out = ioutil.Discard
	messageApi.utils = utils.NewHandlersUtils(emptyLogger)
	middlware.Logger=emptyLogger
}

func TestMessageHandlersImpl_SendMessage(t *testing.T) {
	tests := []*MessageTestCase{
		{
			name: "InternalErrorTestGetUserByID",
			Users: &users.UsersUseCaseMock{
				GetUserByIDFunc: func(ID uint64) (user models.User, e error) {
					return models.User{}, errors.New("Internal error")
				},
				GetUserBySessionFunc: func(session string) (u uint64, e error) {
					return 0, nil
				},
			},
			StatusCode: 500,
			Handler:    http.HandlerFunc(messageApi.SendMessage),
		},
		{
			name: "InvalidJSONTest",
			Users: &users.UsersUseCaseMock{
				GetUserByIDFunc: func(ID uint64) (user models.User, e error) {
					return models.User{}, nil
				},
				GetUserBySessionFunc: func(session string) (u uint64, e error) {
					return 0, nil
				},
			},
			Body:       "Bad JSON",
			StatusCode: 400,
			Handler:    http.HandlerFunc(messageApi.SendMessage),
		},
		{
			name: "InternalSaveMessageErrorTest",
			Users: &users.UsersUseCaseMock{
				GetUserByIDFunc: func(ID uint64) (user models.User, e error) {
					return models.User{}, nil
				},
				GetUserBySessionFunc: func(session string) (u uint64, e error) {
					return 0, nil
				},
			},
			Messages: &messages.MessagesUseCaseMock{
				SaveChatMessageFunc: func(message *models.Message) (u uint64, e error) {
					return 0, errors.New("Internal error")
				},
			},
			Body:       `{"text":"mem"}`,
			StatusCode: 500,
			Handler:    http.HandlerFunc(messageApi.SendMessage),
		},
		{
			name: "SendMessageErrorTest",
			Users: &users.UsersUseCaseMock{
				GetUserByIDFunc: func(ID uint64) (user models.User, e error) {
					return models.User{}, nil
				},
				GetUserBySessionFunc: func(session string) (u uint64, e error) {
					return 0, nil
				},
			},
			Messages: &messages.MessagesUseCaseMock{
				SaveChatMessageFunc: func(message *models.Message) (u uint64, e error) {
					return 0, nil
				},
			},
			Notification: &notification.NotificationUseCaseMock{
				SendMessageFunc: func(chatID uint64, message []byte) error {
					return errors.New("Internal error")
				},
			},
			Body:       `{"text":"mem"}`,
			StatusCode: 200,
			Handler:    http.HandlerFunc(messageApi.SendMessage),
		},
	}
	runTableMessageAPITests(t, tests)
}

func TestMessageHandlersImpl_EditMessage(t *testing.T) {
	tests := []*MessageTestCase{
		{
			name: "WrongCookieTest",
			Users: &users.UsersUseCaseMock{
				GetUserBySessionFunc: func(session string) (u uint64, e error) {
					return 0, nil
				},
				GetUserByIDFunc: func(id uint64) (user models.User, e error) {
					return models.User{}, errors.New("Internal error")
				},
			},
			StatusCode: 500,
			Handler:    http.HandlerFunc(messageApi.EditMessage),
		},
		{
			name: "InvalidJSONTest",
			Users: &users.UsersUseCaseMock{
				GetUserByIDFunc: func(ID uint64) (user models.User, e error) {
					return models.User{}, nil
				},
				GetUserBySessionFunc: func(session string) (u uint64, e error) {
					return 0, nil
				},
			},
			Body:       "Bad JSON",
			StatusCode: 400,
			Handler:    http.HandlerFunc(messageApi.EditMessage),
		},
		{
			name: "InternalEditMessageErrorTest",
			Users: &users.UsersUseCaseMock{
				GetUserByIDFunc: func(ID uint64) (user models.User, e error) {
					return models.User{}, nil
				},
				GetUserBySessionFunc: func(session string) (u uint64, e error) {
					return 0, nil
				},
			},
			Messages: &messages.MessagesUseCaseMock{
				EditMessageFunc: func(message *models.Message, userID uint64) error {
					return errors.New("Internal error")
				},
				GetMessageByIDFunc: func(messageID uint64) (message *models.Message, e error) {
					return &models.Message{}, nil
				},
			},
			Body:       `{"text":"mem"}`,
			StatusCode: 500,
			Handler:    http.HandlerFunc(messageApi.EditMessage),
		},
		{
			name: "Successtest",
			Users: &users.UsersUseCaseMock{
				GetUserByIDFunc: func(ID uint64) (user models.User, e error) {
					return models.User{}, nil
				},
				GetUserBySessionFunc: func(session string) (u uint64, e error) {
					return 0, nil
				},
			},
			Messages: &messages.MessagesUseCaseMock{
				EditMessageFunc: func(message *models.Message, userID uint64) error {
					return nil
				},
				GetMessageByIDFunc: func(messageID uint64) (message *models.Message, e error) {
					return &models.Message{}, nil
				},
			},
			Notification:&notification.NotificationUseCaseMock{
				SendMessageFunc: func(chatID uint64, message []byte) error {
					return nil
				},
			},
			Body:       `{"text":"mem"}`,
			StatusCode: 200,
			Handler:    http.HandlerFunc(messageApi.EditMessage),
		},
		
		
	}

	runTableMessageAPITests(t, tests)

}

func TestMessageHandlersImpl_GetMessagesByChatID(t *testing.T) {
	tests := []*MessageTestCase{
		{
			name: "WrongCookieTest",
			Users: &users.UsersUseCaseMock{
				GetUserBySessionFunc: func(session string) (u uint64, e error) {
					return 0, nil
				},
				GetUserByIDFunc: func(id uint64) (user models.User, e error) {
					return models.User{}, errors.New("Internal error")
				},
			},
			StatusCode: 500,
			Handler:    http.HandlerFunc(messageApi.GetMessagesByChatID),
		},
		{
			name: "InternalGetMessageErrorTest",

			Users: &users.UsersUseCaseMock{
				GetUserByIDFunc: func(ID uint64) (user models.User, e error) {
					return models.User{}, nil
				},
				GetUserBySessionFunc: func(session string) (u uint64, e error) {
					return 0, nil
				},
			},
			Messages: &messages.MessagesUseCaseMock{
				GetChatMessagesFunc: func(chatID uint64, userID uint64) (messages models.Messages, e error) {
					return models.Messages{}, errors.New("Internal error")
				},
			},
			Body:       `{"text":"mem"}`,
			StatusCode: 500,
			Handler:    http.HandlerFunc(messageApi.GetMessagesByChatID),
		},
		{
			name: "SuccessTest",
			Users: &users.UsersUseCaseMock{
				GetUserByIDFunc: func(ID uint64) (user models.User, e error) {
					return models.User{}, nil
				},
				GetUserBySessionFunc: func(session string) (u uint64, e error) {
					return 0, nil
				},
			},
			Messages: &messages.MessagesUseCaseMock{
				GetChatMessagesFunc: func(chatID uint64, userID uint64) (messages models.Messages, e error) {
					return models.Messages{}, nil
				},
			},
			Body:       `{"text":"mem"}`,
			StatusCode: 200,
			Handler:    http.HandlerFunc(messageApi.GetMessagesByChatID),
		},
	}

	runTableMessageAPITests(t, tests)
}

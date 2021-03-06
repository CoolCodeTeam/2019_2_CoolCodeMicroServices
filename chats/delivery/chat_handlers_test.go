package delivery

import (
	"errors"
	useCase "github.com/CoolCodeTeam/2019_2_CoolCodeMicroServices/chats/usecase"
	"github.com/CoolCodeTeam/2019_2_CoolCodeMicroServices/users/repository"
	users "github.com/CoolCodeTeam/2019_2_CoolCodeMicroServices/users/usecase"
	"github.com/CoolCodeTeam/2019_2_CoolCodeMicroServices/utils"
	middleware "github.com/CoolCodeTeam/2019_2_CoolCodeMicroServices/utils/middlwares"
	"github.com/CoolCodeTeam/2019_2_CoolCodeMicroServices/utils/models"

	"github.com/sirupsen/logrus"
	"github.com/steinfletcher/apitest"
	"io/ioutil"
	"net/http"
	"testing"
)

var chatApi ChatHandlers
var middlware = middleware.HandlersMiddlwares{}

type ChatTestCase struct {
	name       string
	Body       string
	SessionID  string
	Headers    map[string]string
	Method     string
	URL        string
	Response   string
	StatusCode int
	Handler    func(w http.ResponseWriter, r *http.Request)
	Chats      useCase.ChatsUseCase
	Sessions   repository.SessionRepository
	Users      users.UsersUseCase
}

func runTableChatAPITests(t *testing.T, cases []*ChatTestCase) {
	for _, c := range cases {
		runChatAPITest(t, c)
	}
}

func runChatAPITest(t *testing.T, testCase *ChatTestCase) {
	t.Run(testCase.name, func(t *testing.T) {
		if testCase.Chats != nil {
			chatApi.Chats = testCase.Chats
		}
		if testCase.Users != nil {
			chatApi.Users = testCase.Users
			middlware.Users = testCase.Users
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
	chatApi.utils = utils.NewHandlersUtils(emptyLogger)
	middlware.Logger = emptyLogger
}

func TestChatHandlers_PostChat(t *testing.T) {
	tests := []*ChatTestCase{

		{
			name: "InternalError",
			Users: &users.UsersUseCaseMock{
				GetUserByIDFunc: func(id uint64) (user models.User, e error) {
					return models.User{}, errors.New("Internal error")
				},
				GetUserBySessionFunc: func(session string) (u uint64, e error) {
					return 0, nil
				},
			},
			StatusCode: 500,
			Handler:    http.HandlerFunc(chatApi.PostChat),
		},
		{
			name: "InvalidJsonTest",
			Users: &users.UsersUseCaseMock{
				GetUserByIDFunc: func(id uint64) (user models.User, e error) {
					return models.User{}, nil
				},
				GetUserBySessionFunc: func(session string) (u uint64, e error) {
					return 0, nil
				},
			},
			StatusCode: 400,
			Handler:    http.HandlerFunc(chatApi.PostChat),
			Body:       "BadJson",
		},
		{
			name: "BadUserToTest",
			Users: &users.UsersUseCaseMock{
				GetUserByIDFunc: func(id uint64) (user models.User, e error) {
					if id == 1 {
						return models.User{}, errors.New("Internal error")
					}
					return models.User{}, nil
				},
				GetUserBySessionFunc: func(session string) (u uint64, e error) {
					return 0, nil
				},
			},
			StatusCode: 500,
			Handler:    http.HandlerFunc(chatApi.PostChat),
			Body:       `{"user_id":1}`,
		},
		{
			name: "PutChatErrorTest",

			Users: &users.UsersUseCaseMock{
				GetUserByIDFunc: func(id uint64) (user models.User, e error) {
					return models.User{}, nil
				},
				GetUserBySessionFunc: func(session string) (u uint64, e error) {
					return 0, nil
				},
			},
			Chats: &useCase.ChatsUseCaseMock{
				PutChatFunc: func(Chat *models.Chat) (u uint64, e error) {
					return 0, errors.New("Internal error")
				},
			},
			StatusCode: 500,
			Handler:    http.HandlerFunc(chatApi.PostChat),
			Body:       `{"user_id":1}`,
		},
		{
			name: "TestSuccess",
			Users: &users.UsersUseCaseMock{
				GetUserByIDFunc: func(id uint64) (user models.User, e error) {
					return models.User{}, nil
				},
				GetUserBySessionFunc: func(session string) (u uint64, e error) {
					return 0, nil
				},
			},
			Chats: &useCase.ChatsUseCaseMock{
				PutChatFunc: func(Chat *models.Chat) (u uint64, e error) {
					return 0, nil
				},
			},
			StatusCode: 200,
			Handler:    http.HandlerFunc(chatApi.PostChat),
			Body:       `{"user_id":1}`,
		},
	}
	runTableChatAPITests(t, tests)
}

func TestChatHandlers_GetChatsByUser(t *testing.T) {
	tests := []*ChatTestCase{
		{
			name: "WrongCookieTest",
			Users: &users.UsersUseCaseMock{
				GetUserBySessionFunc: func(session string) (u uint64, e error) {
					return 0, errors.New("Internal error")
				},
			},
			StatusCode: 401,
			Handler:    http.HandlerFunc(chatApi.GetChatsByUser),
		},
		{
			name: "WrongRequestedIDTest",
			Users: &users.UsersUseCaseMock{
				GetUserBySessionFunc: func(session string) (u uint64, e error) {
					return 1, nil
				},
			},
			StatusCode: 401,
			Handler:    http.HandlerFunc(chatApi.GetChatsByUser),
		},
		{
			name: "GetChatsInternalError",
			Users: &users.UsersUseCaseMock{
				GetUserBySessionFunc: func(session string) (u uint64, e error) {
					return 0, nil
				},
			},
			Chats: &useCase.ChatsUseCaseMock{
				GetChatsByUserIDFunc: func(ID uint64) (chats []models.Chat, e error) {
					return []models.Chat{}, errors.New("Internal error")
				},
			},
			StatusCode: 500,
			Handler:    http.HandlerFunc(chatApi.GetChatsByUser),
		},
		{
			name: "GetWorkspacesInternalError",
			Users: &users.UsersUseCaseMock{
				GetUserBySessionFunc: func(session string) (u uint64, e error) {
					return 0, nil
				},
			},
			Chats: &useCase.ChatsUseCaseMock{
				GetChatsByUserIDFunc: func(ID uint64) (chats []models.Chat, e error) {
					return []models.Chat{}, nil
				},
				GetWorkspacesByUserIDFunc: func(ID uint64) (workspaces []models.Workspace, e error) {
					return []models.Workspace{}, errors.New("Internal error")
				},
			},
			StatusCode: 500,
			Handler:    http.HandlerFunc(chatApi.GetChatsByUser),
		},
		{
			name: "SuccessTest",
			Sessions: &repository.SessionRepositoryMock{
				GetIDFunc: func(session string) (u uint64, e error) {
					return 0, nil
				},
			},
			Chats: &useCase.ChatsUseCaseMock{
				GetChatsByUserIDFunc: func(ID uint64) (chats []models.Chat, e error) {
					return []models.Chat{}, nil
				},
				GetWorkspacesByUserIDFunc: func(ID uint64) (workspaces []models.Workspace, e error) {
					return []models.Workspace{}, nil
				},
			},
			StatusCode: 200,
			Handler:    http.HandlerFunc(chatApi.GetChatsByUser),
		},
	}
	runTableChatAPITests(t, tests)
}

func TestChatHandlers_GetChatById(t *testing.T) {
	tests := []*ChatTestCase{
		{
			name: "WrongCookieTest",
			Users: &users.UsersUseCaseMock{
				GetUserBySessionFunc: func(session string) (u uint64, e error) {
					return 0, errors.New("Internal error")
				},
			},
			StatusCode: 401,
			Handler:    http.HandlerFunc(chatApi.GetChatById),
		},
		{
			name: "InternalErrorTest",
			Users: &users.UsersUseCaseMock{
				GetUserByIDFunc: func(id uint64) (user models.User, e error) {
					return models.User{}, errors.New("Internal error")
				},
				GetUserBySessionFunc: func(session string) (u uint64, e error) {
					return 0, nil
				},
			},
			Chats: &useCase.ChatsUseCaseMock{
				GetChatByIDFunc: func(userID uint64, ID uint64) (chat models.Chat, e error) {
					return models.Chat{}, errors.New("Internal error")
				},
			},
			StatusCode: 500,
			Handler:    http.HandlerFunc(chatApi.GetChatById),
		},
		{
			name: "InternalErrorTest1",
			Users: &users.UsersUseCaseMock{
				GetUserByIDFunc: func(id uint64) (user models.User, e error) {
					return models.User{}, nil
				},
				GetUserBySessionFunc: func(session string) (u uint64, e error) {
					return 0, nil
				},
			},
			Chats: &useCase.ChatsUseCaseMock{
				GetChatByIDFunc: func(userID uint64, ID uint64) (chat models.Chat, e error) {
					return models.Chat{}, errors.New("Internal error")
				},
			},
			StatusCode: 500,
			Handler:    http.HandlerFunc(chatApi.GetChatById),
		},
		{
			name: "SuccessTest",
			Users: &users.UsersUseCaseMock{
				GetUserByIDFunc: func(id uint64) (user models.User, e error) {
					return models.User{}, nil
				},
				GetUserBySessionFunc: func(session string) (u uint64, e error) {
					return 0, nil
				},
			},
			Chats: &useCase.ChatsUseCaseMock{
				GetChatByIDFunc: func(userID uint64, ID uint64) (chat models.Chat, e error) {
					return models.Chat{}, nil
				},
			},
			StatusCode: 200,
			Handler:    http.HandlerFunc(chatApi.GetChatById),
		},
	}
	runTableChatAPITests(t, tests)
}

func TestChatHandlers_RemoveChat(t *testing.T) {
	tests := []*ChatTestCase{
		{
			name: "WrongCookieTest",
			Users: &users.UsersUseCaseMock{
				GetUserBySessionFunc: func(session string) (u uint64, e error) {
					return 0, errors.New("Internal error")
				},
			},
			StatusCode: 401,
			Handler:    http.HandlerFunc(chatApi.RemoveChat),
		},
		{
			name: "WrongCookieTest1",
			Users: &users.UsersUseCaseMock{
				GetUserBySessionFunc: func(session string) (u uint64, e error) {
					return 0, nil
				},
				GetUserByIDFunc: func(id uint64) (user models.User, e error) {
					return models.User{}, errors.New("internal error")
				},
			},
			StatusCode: 401,
			Handler:    http.HandlerFunc(chatApi.RemoveChat),
		},

		{
			name: "InternalErrorTest",
			Sessions: &repository.SessionRepositoryMock{
				GetIDFunc: func(session string) (u uint64, e error) {
					return 0, nil
				},
			},
			Users: &users.UsersUseCaseMock{
				GetUserByIDFunc: func(id uint64) (user models.User, e error) {
					return models.User{}, nil
				},
				GetUserBySessionFunc: func(session string) (u uint64, e error) {
					return 0, nil
				},
			},
			Chats: &useCase.ChatsUseCaseMock{
				DeleteChatFunc: func(userID uint64, chatId uint64) error {
					return errors.New("Internal error")
				},
			},
			StatusCode: 500,
			Handler:    http.HandlerFunc(chatApi.RemoveChat),
		},
		{
			name:     "SuccessTest",
			Sessions: &repository.SessionRepositoryMock{},
			Users: &users.UsersUseCaseMock{
				GetUserByIDFunc: func(id uint64) (user models.User, e error) {
					return models.User{}, nil
				},
				GetUserBySessionFunc: func(session string) (u uint64, e error) {
					return 0, nil
				},
			},
			Chats: &useCase.ChatsUseCaseMock{
				DeleteChatFunc: func(userID uint64, chatId uint64) error {
					return nil
				},
			},
			StatusCode: 200,
			Handler:    http.HandlerFunc(chatApi.RemoveChat),
		},
	}
	runTableChatAPITests(t, tests)
}

func TestChatHandlers_PostChannel(t *testing.T) {
	tests := []*ChatTestCase{
		{
			name: "WrongCookieTest",
			Users: &users.UsersUseCaseMock{
				GetUserBySessionFunc: func(session string) (u uint64, e error) {
					return 0, errors.New("Internal error")
				},
			},
			StatusCode: 401,
			Handler:    http.HandlerFunc(chatApi.PostChannel),
		},
		{
			name: "WrongCookieTest1",
			Users: &users.UsersUseCaseMock{
				GetUserBySessionFunc: func(session string) (u uint64, e error) {
					return 0, nil
				},
				GetUserByIDFunc: func(id uint64) (user models.User, e error) {
					return models.User{}, errors.New("internal error")
				},
			},
			StatusCode: 401,
			Handler:    http.HandlerFunc(chatApi.PostChannel),
		},
		{
			name: "InvalidJsonTest",
			Users: &users.UsersUseCaseMock{
				GetUserByIDFunc: func(id uint64) (user models.User, e error) {
					return models.User{}, nil
				},
				GetUserBySessionFunc: func(session string) (u uint64, e error) {
					return 0, nil
				},
			},
			StatusCode: 400,
			Handler:    http.HandlerFunc(chatApi.PostChat),
			Body:       "BadJson",
		},
		{
			name: "PutChannelErrorTest",
			Users: &users.UsersUseCaseMock{
				GetUserByIDFunc: func(id uint64) (user models.User, e error) {
					return models.User{}, nil
				},
				GetUserBySessionFunc: func(session string) (u uint64, e error) {
					return 0, nil
				},
			},
			Chats: &useCase.ChatsUseCaseMock{
				CreateChannelFunc: func(channel *models.Channel) (u uint64, e error) {
					return 0, errors.New("Internal error")
				},
			},
			StatusCode: 500,
			Handler:    http.HandlerFunc(chatApi.PostChannel),
			Body:       `{"user_id":1}`,
		},
		{
			name: "TestSuccess",
			Users: &users.UsersUseCaseMock{
				GetUserByIDFunc: func(id uint64) (user models.User, e error) {
					return models.User{}, nil
				},
				GetUserBySessionFunc: func(session string) (u uint64, e error) {
					return 0, nil
				},
			},
			Chats: &useCase.ChatsUseCaseMock{
				CreateChannelFunc: func(channel *models.Channel) (u uint64, e error) {
					return 0, nil
				},
			},
			StatusCode: 200,
			Handler:    http.HandlerFunc(chatApi.PostChannel),
			Body:       `{"user_id":1}`,
		},
	}
	runTableChatAPITests(t, tests)
}

func TestChatHandlers_GetChannelById(t *testing.T) {
	tests := []*ChatTestCase{
		{
			name: "WrongCookieTest",
			Users: &users.UsersUseCaseMock{
				GetUserBySessionFunc: func(session string) (u uint64, e error) {
					return 0, errors.New("Internal error")
				},
			},
			StatusCode: 401,
			Handler:    http.HandlerFunc(chatApi.GetChannelById),
		},
		{
			name: "WrongCookieTest1",
			Users: &users.UsersUseCaseMock{
				GetUserBySessionFunc: func(session string) (u uint64, e error) {
					return 0, nil
				},
				GetUserByIDFunc: func(id uint64) (user models.User, e error) {
					return models.User{}, errors.New("internal error")
				},
			},
			StatusCode: 500,
			Handler:    http.HandlerFunc(chatApi.GetChannelById),
		},
		{
			name: "InternalErrorTest",
			Users: &users.UsersUseCaseMock{
				GetUserByIDFunc: func(id uint64) (user models.User, e error) {
					return models.User{}, nil
				},
				GetUserBySessionFunc: func(session string) (u uint64, e error) {
					return 0, nil
				},
			},
			Chats: &useCase.ChatsUseCaseMock{
				GetChannelByIDFunc: func(userID uint64, ID uint64) (channel models.Channel, e error) {
					return models.Channel{}, errors.New("Internal error")
				},
			},
			StatusCode: 500,
			Handler:    http.HandlerFunc(chatApi.GetChannelById),
		},
		{
			name: "SuccessTest",
			Users: &users.UsersUseCaseMock{
				GetUserByIDFunc: func(id uint64) (user models.User, e error) {
					return models.User{}, nil
				},
				GetUserBySessionFunc: func(session string) (u uint64, e error) {
					return 0, nil
				},
			},
			Chats: &useCase.ChatsUseCaseMock{
				GetChannelByIDFunc: func(userID uint64, ID uint64) (channel models.Channel, e error) {
					return models.Channel{}, nil
				},
			},
			StatusCode: 200,
			Handler:    http.HandlerFunc(chatApi.GetChannelById),
		},
	}
	runTableChatAPITests(t, tests)
}

func TestChatHandlers_EditChannel(t *testing.T) {

	tests := []*ChatTestCase{
		{
			name: "WrongCookieTest1",
			Users: &users.UsersUseCaseMock{
				GetUserBySessionFunc: func(session string) (u uint64, e error) {
					return 0, nil
				},
				GetUserByIDFunc: func(id uint64) (user models.User, e error) {
					return models.User{}, errors.New("internal error")
				},
			},
			StatusCode: 500,
			Handler:    http.HandlerFunc(chatApi.EditChannel),
		},
		{
			name: "InvalidJsonTest",
			Users: &users.UsersUseCaseMock{
				GetUserByIDFunc: func(id uint64) (user models.User, e error) {
					return models.User{}, nil
				},
				GetUserBySessionFunc: func(session string) (u uint64, e error) {
					return 0, nil
				},
			},
			StatusCode: 400,
			Handler:    http.HandlerFunc(chatApi.EditChannel),
			Body:       "BadJson",
		},
		{
			name: "EditChannelErrorTest",
			Users: &users.UsersUseCaseMock{
				GetUserByIDFunc: func(id uint64) (user models.User, e error) {
					return models.User{}, nil
				},
				GetUserBySessionFunc: func(session string) (u uint64, e error) {
					return 0, nil
				},
			},
			Chats: &useCase.ChatsUseCaseMock{
				EditChannelFunc: func(userID uint64, channel *models.Channel) error {
					return errors.New("Internal error")
				},
			},
			StatusCode: 500,
			Handler:    http.HandlerFunc(chatApi.EditChannel),
			Body:       `{"user_id":1}`,
		},
		{
			name: "TestSuccess",
			Users: &users.UsersUseCaseMock{
				GetUserByIDFunc: func(id uint64) (user models.User, e error) {
					return models.User{}, nil
				},
				GetUserBySessionFunc: func(session string) (u uint64, e error) {
					return 0, nil
				},
			},
			Chats: &useCase.ChatsUseCaseMock{
				EditChannelFunc: func(userID uint64, channel *models.Channel) error {
					return nil
				},
			},
			StatusCode: 200,
			Handler:    http.HandlerFunc(chatApi.EditChannel),
			Body:       `{"user_id":1}`,
		},
	}
	runTableChatAPITests(t, tests)
}

func TestChatHandlers_RemoveChannel(t *testing.T) {
	tests := []*ChatTestCase{
		{
			name: "WrongCookieTest1",
			Users: &users.UsersUseCaseMock{
				GetUserBySessionFunc: func(session string) (u uint64, e error) {
					return 0, nil
				},
				GetUserByIDFunc: func(id uint64) (user models.User, e error) {
					return models.User{}, errors.New("internal error")
				},
			},
			StatusCode: 500,
			Handler:    http.HandlerFunc(chatApi.RemoveChannel),
		},
		{
			name: "InternalErrorTest",
			Users: &users.UsersUseCaseMock{
				GetUserByIDFunc: func(id uint64) (user models.User, e error) {
					return models.User{}, nil
				},
				GetUserBySessionFunc: func(session string) (u uint64, e error) {
					return 0, nil
				},
			},
			Chats: &useCase.ChatsUseCaseMock{
				DeleteChannelFunc: func(userID uint64, chatId uint64) error {
					return errors.New("Internal error")
				},
			},
			StatusCode: 500,
			Handler:    http.HandlerFunc(chatApi.RemoveChannel),
		},
		{
			name: "SuccessTest",

			Users: &users.UsersUseCaseMock{
				GetUserByIDFunc: func(id uint64) (user models.User, e error) {
					return models.User{}, nil
				},
				GetUserBySessionFunc: func(session string) (u uint64, e error) {
					return 0, nil
				},
			},
			Chats: &useCase.ChatsUseCaseMock{
				DeleteChannelFunc: func(userID uint64, chatId uint64) error {
					return nil
				},
			},
			StatusCode: 200,
			Handler:    http.HandlerFunc(chatApi.RemoveChannel),
		},
	}
	runTableChatAPITests(t, tests)
}

func TestChatHandlers_PostWorkspace(t *testing.T) {
	tests := []*ChatTestCase{
		{
			name: "WrongCookieTest1",
			Users: &users.UsersUseCaseMock{
				GetUserBySessionFunc: func(session string) (u uint64, e error) {
					return 0, nil
				},
				GetUserByIDFunc: func(id uint64) (user models.User, e error) {
					return models.User{}, errors.New("internal error")
				},
			},
			StatusCode: 500,
			Handler:    http.HandlerFunc(chatApi.PostWorkspace),
		},
		{
			name: "InvalidJsonTest",
			Users: &users.UsersUseCaseMock{
				GetUserByIDFunc: func(id uint64) (user models.User, e error) {
					return models.User{}, nil
				},
				GetUserBySessionFunc: func(session string) (u uint64, e error) {
					return 0, nil
				},
			},
			StatusCode: 400,
			Handler:    http.HandlerFunc(chatApi.PostWorkspace),
			Body:       "BadJson",
		},
		{
			name: "PutChannelErrorTest",
			Users: &users.UsersUseCaseMock{
				GetUserByIDFunc: func(id uint64) (user models.User, e error) {
					return models.User{}, nil
				},
				GetUserBySessionFunc: func(session string) (u uint64, e error) {
					return 0, nil
				},
			},
			Chats: &useCase.ChatsUseCaseMock{
				CreateWorkspaceFunc: func(workspace *models.Workspace) (u uint64, e error) {
					return 0, errors.New("Internal error")
				},
			},
			StatusCode: 500,
			Handler:    http.HandlerFunc(chatApi.PostWorkspace),
			Body:       `{"user_id":1}`,
		},
		{
			name: "TestSuccess",

			Users: &users.UsersUseCaseMock{
				GetUserByIDFunc: func(id uint64) (user models.User, e error) {
					return models.User{}, nil
				},
				GetUserBySessionFunc: func(session string) (u uint64, e error) {
					return 0, nil
				},
			},
			Chats: &useCase.ChatsUseCaseMock{
				CreateWorkspaceFunc: func(workspace *models.Workspace) (u uint64, e error) {
					return 0, nil
				},
			},
			StatusCode: 200,
			Handler:    http.HandlerFunc(chatApi.PostWorkspace),
			Body:       `{"user_id":1}`,
		},
	}
	runTableChatAPITests(t, tests)
}

func TestChatHandlers_GetWorkspaceById(t *testing.T) {
	tests := []*ChatTestCase{
		{
			name: "WrongCookieTest1",
			Users: &users.UsersUseCaseMock{
				GetUserBySessionFunc: func(session string) (u uint64, e error) {
					return 0, nil
				},
				GetUserByIDFunc: func(id uint64) (user models.User, e error) {
					return models.User{}, errors.New("internal error")
				},
			},
			StatusCode: 500,
			Handler:    http.HandlerFunc(chatApi.GetWorkspaceById),
		},
		{
			name: "InternalErrorTest",

			Users: &users.UsersUseCaseMock{
				GetUserByIDFunc: func(id uint64) (user models.User, e error) {
					return models.User{}, nil
				},
				GetUserBySessionFunc: func(session string) (u uint64, e error) {
					return 0, nil
				},
			},
			Chats: &useCase.ChatsUseCaseMock{
				GetWorkspaceByIDFunc: func(userID uint64, ID uint64) (workspace models.Workspace, e error) {
					return models.Workspace{}, errors.New("Internal error")
				},
			},
			StatusCode: 500,
			Handler:    http.HandlerFunc(chatApi.GetWorkspaceById),
		},
		{
			name: "SuccessTest",
			Users: &users.UsersUseCaseMock{
				GetUserByIDFunc: func(id uint64) (user models.User, e error) {
					return models.User{}, nil
				},
				GetUserBySessionFunc: func(session string) (u uint64, e error) {
					return 0, nil
				},
			},
			Chats: &useCase.ChatsUseCaseMock{
				GetWorkspaceByIDFunc: func(userID uint64, ID uint64) (workspace models.Workspace, e error) {
					return models.Workspace{}, nil
				},
			},
			StatusCode: 200,
			Handler:    http.HandlerFunc(chatApi.GetWorkspaceById),
		},
	}
	runTableChatAPITests(t, tests)
}

func TestChatHandlers_EditWorkspace(t *testing.T) {
	tests := []*ChatTestCase{
		{
			name: "WrongCookieTest1",
			Users: &users.UsersUseCaseMock{
				GetUserBySessionFunc: func(session string) (u uint64, e error) {
					return 0, nil
				},
				GetUserByIDFunc: func(id uint64) (user models.User, e error) {
					return models.User{}, errors.New("internal error")
				},
			},
			StatusCode: 401,
			Handler:    http.HandlerFunc(chatApi.EditWorkspace),
		},
		{
			name: "InvalidJsonTest",
			Users: &users.UsersUseCaseMock{
				GetUserByIDFunc: func(id uint64) (user models.User, e error) {
					return models.User{}, nil
				},
				GetUserBySessionFunc: func(session string) (u uint64, e error) {
					return 0, nil
				},
			},
			StatusCode: 400,
			Handler:    http.HandlerFunc(chatApi.EditWorkspace),
			Body:       "BadJson",
		},
		{
			name: "EditWorkspaceErrorTest",
			Users: &users.UsersUseCaseMock{
				GetUserByIDFunc: func(id uint64) (user models.User, e error) {
					return models.User{}, nil
				},
				GetUserBySessionFunc: func(session string) (u uint64, e error) {
					return 0, nil
				},
			},
			Chats: &useCase.ChatsUseCaseMock{
				EditWorkspaceFunc: func(userID uint64, workspace *models.Workspace) error {
					return errors.New("Internal error")
				},
			},
			StatusCode: 500,
			Handler:    http.HandlerFunc(chatApi.EditWorkspace),
			Body:       `{"user_id":1}`,
		},
		{
			name: "TestSuccess",
			Users: &users.UsersUseCaseMock{
				GetUserByIDFunc: func(id uint64) (user models.User, e error) {
					return models.User{}, nil
				},
				GetUserBySessionFunc: func(session string) (u uint64, e error) {
					return 0, nil
				},
			},
			Chats: &useCase.ChatsUseCaseMock{
				EditWorkspaceFunc: func(userID uint64, workspace *models.Workspace) error {
					return nil
				},
			},
			StatusCode: 200,
			Handler:    http.HandlerFunc(chatApi.EditWorkspace),
			Body:       `{"user_id":1}`,
		},
	}
	runTableChatAPITests(t, tests)
}

func TestChatHandlers_RemoveWorkspace(t *testing.T) {
	tests := []*ChatTestCase{
		{
			name: "WrongCookieTest1",
			Users: &users.UsersUseCaseMock{
				GetUserBySessionFunc: func(session string) (u uint64, e error) {
					return 0, nil
				},
				GetUserByIDFunc: func(id uint64) (user models.User, e error) {
					return models.User{}, errors.New("internal error")
				},
			},
			StatusCode: 401,
			Handler:    http.HandlerFunc(chatApi.RemoveWorkspace),
		},
		{
			name: "InternalErrorTest",
			Users: &users.UsersUseCaseMock{
				GetUserByIDFunc: func(id uint64) (user models.User, e error) {
					return models.User{}, nil
				},
				GetUserBySessionFunc: func(session string) (u uint64, e error) {
					return 0, nil
				},
			},
			Chats: &useCase.ChatsUseCaseMock{
				DeleteWorkspaceFunc: func(userID uint64, chatId uint64) error {
					return errors.New("Internal error")
				},
			},
			StatusCode: 500,
			Handler:    http.HandlerFunc(chatApi.RemoveWorkspace),
		},
		{
			name: "SuccessTest",
			Users: &users.UsersUseCaseMock{
				GetUserByIDFunc: func(id uint64) (user models.User, e error) {
					return models.User{}, nil
				},
				GetUserBySessionFunc: func(session string) (u uint64, e error) {
					return 0, nil
				},
			},
			Chats: &useCase.ChatsUseCaseMock{
				DeleteWorkspaceFunc: func(userID uint64, chatId uint64) error {
					return nil
				},
			},
			StatusCode: 200,
			Handler:    http.HandlerFunc(chatApi.RemoveWorkspace),
		},
	}
	runTableChatAPITests(t, tests)
}

package delivery

import (
	"net/http"
	"strings"

	useCase "github.com/CoolCodeTeam/2019_2_CoolCodeMicroServices/notifications/usecase"

	chats "github.com/CoolCodeTeam/2019_2_CoolCodeMicroServices/chats/usecase"
	users "github.com/CoolCodeTeam/2019_2_CoolCodeMicroServices/users/usecase"
	"github.com/CoolCodeTeam/2019_2_CoolCodeMicroServices/utils"
	"github.com/CoolCodeTeam/2019_2_CoolCodeMicroServices/utils/models"
	"github.com/gorilla/websocket"
)

type NotificationHandlers struct {
	notificationUseCase useCase.NotificationUseCase
	chatsUseCase        chats.ChatsUseCase
	Users               users.UsersUseCase

	utils utils.HandlersUtils
}

func NewNotificationHandlers(users users.UsersUseCase,
	chats chats.ChatsUseCase, notifications useCase.NotificationUseCase, utils utils.HandlersUtils) NotificationHandlers {
	return NotificationHandlers{
		notificationUseCase: notifications,
		chatsUseCase:        chats,
		Users:               users,

		utils: utils,
	}
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (h *NotificationHandlers) HandleNewWSConnection(w http.ResponseWriter, r *http.Request) {

}

func (h NotificationHandlers) parseCookie(r *http.Request) (models.User, error) {
	id := r.Context().Value("user_id").(uint64)
	print(id)
	user, err := h.Users.GetUserByID(id)
	if err == nil {
		return user, nil
	} else {
		return user, err
	}
}

func isChannel(r *http.Request) bool {
	return strings.Contains(r.URL.String(), "channels")
}

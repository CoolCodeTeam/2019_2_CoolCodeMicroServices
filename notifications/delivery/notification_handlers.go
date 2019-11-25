package delivery

import (
	chats "github.com/go-park-mail-ru/2019_2_CoolCodeMicroServices/chats/usecase"
	"github.com/go-park-mail-ru/2019_2_CoolCodeMicroServices/notifications/usecase"
	users "github.com/go-park-mail-ru/2019_2_CoolCodeMicroServices/users/usecase"
	"github.com/go-park-mail-ru/2019_2_CoolCodeMicroServices/utils"
	"github.com/go-park-mail-ru/2019_2_CoolCodeMicroServices/utils/models"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"net/http"
	"strconv"
	"strings"
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
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		h.utils.HandleError(models.NewServerError(err, http.StatusBadRequest, "Can not upgrade connection"), w, r)
		return
	}

	requestedID, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		h.utils.LogError(err, r)
	}

	userID := r.Context().Value("user_id").(uint64)

	//Проверяем доступ к чату
	ok := true
	if isChannel(r) {
		ok, err = h.chatsUseCase.CheckChannelPermission(userID, uint64(requestedID))
	} else {
		ok, err = h.chatsUseCase.CheckChatPermission(userID, uint64(requestedID))
	}
	if err != nil {
		h.utils.HandleError(err, w, r)
		return
	}
	if !ok {
		h.utils.HandleError(models.NewClientError(nil, http.StatusForbidden, "Not permission to chat:("),
			w, r)
		return
	}
	//Достаем Handler с помощью Messages
	hub, err := h.notificationUseCase.OpenConn(uint64(requestedID))
	go hub.Run()
	//Запускаем event loop
	hub.AddClientChan <- ws

	for {
		var m []byte

		_, m, err := ws.ReadMessage()

		if err != nil {
			hub.RemoveClient(ws)
			return
		}
		hub.BroadcastChan <- m
	}

}

func (h NotificationHandlers) parseCookie(r *http.Request) (models.User, error) {
	id := r.Context().Value("user_id").(uint64)
	//if err != nil {
	//	return models.User{}, models.NewClientError(err, http.StatusUnauthorized, "Bad request : not valid cookie:(")
	//}
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

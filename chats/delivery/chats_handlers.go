package delivery

import (
	"encoding/json"
	"fmt"
	"github.com/CoolCodeTeam/2019_2_CoolCodeMicroServices/chats/usecase"
	users "github.com/CoolCodeTeam/2019_2_CoolCodeMicroServices/users/usecase"
	"github.com/CoolCodeTeam/2019_2_CoolCodeMicroServices/utils"
	"github.com/CoolCodeTeam/2019_2_CoolCodeMicroServices/utils/models"
	"github.com/mailru/easyjson"

	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type ChatHandlers struct {
	Chats useCase.ChatsUseCase
	Users users.UsersUseCase
	utils utils.HandlersUtils
}

func NewChatHandlers(users users.UsersUseCase,
	chats useCase.ChatsUseCase, utils utils.HandlersUtils) ChatHandlers {
	return ChatHandlers{
		Chats: chats,
		Users: users,
		utils: utils,
	}
}

func (c *ChatHandlers) PostChat(w http.ResponseWriter, r *http.Request) {

	user, err := c.parseCookie(r)
	if err != nil {
		c.utils.HandleError(err, w, r)
		return
	}
	var newChatModel models.CreateChatModel
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&newChatModel) //TODO ?
	if err != nil {
		c.utils.HandleError(models.NewClientError(err, http.StatusBadRequest, "Bad request: malformed data:("), w, r)
		return
	}
	userTo, err := c.Users.GetUserByID(newChatModel.UserID)
	if err != nil {
		c.utils.HandleError(err, w, r)
		return
	}

	model := models.NewChatModel(userTo.Username, user.ID, userTo.ID)
	id, err := c.Chats.PutChat(model)
	if err != nil {
		c.utils.HandleError(err, w, r)
		return
	}
	jsonResponse, _ := json.Marshal(map[string]uint64{
		"id": id,
	}) //TODO: ?
	_, err = w.Write(jsonResponse)
	if err != nil {
		c.utils.LogError(err, r)
	}
	w.WriteHeader(http.StatusOK)
}

func (c *ChatHandlers) GetChatsByUser(w http.ResponseWriter, r *http.Request) {
	requestedID, _ := strconv.Atoi(mux.Vars(r)["id"])
	cookieID := r.Context().Value("user_id").(uint64)
	if cookieID != uint64(requestedID) {
		c.utils.HandleError(
			models.NewClientError(nil, http.StatusUnauthorized, fmt.Sprintf("Actual id: %d, Requested id: %d", cookieID, requestedID)),
			w, r)
		return
	}
	chats, err := c.Chats.GetChatsByUserID(uint64(requestedID))
	if err != nil {
		c.utils.HandleError(err, w, r)
		return
	}
	workspaces, err := c.Chats.GetWorkspacesByUserID(uint64(requestedID))
	if err != nil {
		c.utils.HandleError(err, w, r)
		return
	}
	responseChats := models.ResponseChatsArray{Chats: chats, Workspaces: workspaces}
	jsonChat, err := easyjson.Marshal(responseChats)
	if err != nil {
		c.utils.LogError(err, r)
	}
	_, err = w.Write(jsonChat)
	if err != nil {
		c.utils.LogError(err, r)
	}
}

func (c *ChatHandlers) GetChatById(w http.ResponseWriter, r *http.Request) {
	requestedID, _ := strconv.Atoi(mux.Vars(r)["id"])
	user, err := c.parseCookie(r)
	if err != nil {
		c.utils.HandleError(err, w, r)
		return
	}
	chat, err := c.Chats.GetChatByID(user.ID, uint64(requestedID))
	if err != nil {
		c.utils.HandleError(err, w, r)
		return
	}
	jsonChat, err := easyjson.Marshal(chat)
	if err != nil {
		c.utils.LogError(err, r)
	}
	_, err = w.Write(jsonChat)
	if err != nil {
		c.utils.LogError(err, r)
	}
}

func (c *ChatHandlers) RemoveChat(w http.ResponseWriter, r *http.Request) {

	requestedID, _ := strconv.Atoi(mux.Vars(r)["id"])
	//TODO:Check error
	user, err := c.parseCookie(r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	err = c.Chats.DeleteChat(user.ID, uint64(requestedID))
	if err != nil {
		c.utils.HandleError(err, w, r)
		return
	}
}

func (c *ChatHandlers) PostChannel(w http.ResponseWriter, r *http.Request) {
	requestedID, _ := strconv.Atoi(mux.Vars(r)["id"])

	user, err := c.parseCookie(r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	var newChannelModel models.Channel
	newChannelModel.Members = append(newChannelModel.Members, user.ID)
	newChannelModel.Admins = append(newChannelModel.Admins, user.ID)
	newChannelModel.CreatorID = user.ID
	newChannelModel.WorkspaceID = uint64(requestedID)
	err = easyjson.UnmarshalFromReader(r.Body, &newChannelModel)
	if err != nil {
		c.utils.HandleError(models.NewClientError(err, http.StatusBadRequest,
			"Bad request: malformed data:("), w, r)
		return
	}

	id, err := c.Chats.CreateChannel(&newChannelModel)
	if err != nil {
		c.utils.HandleError(err, w, r)
		return
	}
	jsonResponse, err := json.Marshal(map[string]uint64{
		"id": id,
	}) //TODO: ?
	if err != nil {
		c.utils.LogError(err, r)
	}
	_, err = w.Write(jsonResponse)
	if err != nil {
		c.utils.LogError(err, r)
	}
	w.WriteHeader(http.StatusOK)
}

func (c *ChatHandlers) GetChannelById(w http.ResponseWriter, r *http.Request) {
	requestedID, _ := strconv.Atoi(mux.Vars(r)["id"])
	user, err := c.parseCookie(r)
	if err != nil {
		c.utils.HandleError(err, w, r)
		return
	}
	channel, err := c.Chats.GetChannelByID(user.ID, uint64(requestedID))
	if err != nil {
		c.utils.HandleError(err, w, r)
		return
	}
	jsonChannel, err := easyjson.Marshal(channel)
	if err != nil {
		c.utils.LogError(err, r)
	}
	_, err = w.Write(jsonChannel)
	if err != nil {
		c.utils.LogError(err, r)
	}
}

func (c *ChatHandlers) EditChannel(w http.ResponseWriter, r *http.Request) {

	user, err := c.parseCookie(r)
	if err != nil {
		c.utils.HandleError(err, w, r)
		return
	}
	var newChannel models.Channel

	err = easyjson.UnmarshalFromReader(r.Body, &newChannel)
	if err != nil {
		c.utils.HandleError(models.NewClientError(err, http.StatusBadRequest,
			"Bad request: malformed data:("), w, r)
		return
	}

	err = c.Chats.EditChannel(user.ID, &newChannel)
	if err != nil {
		c.utils.HandleError(err, w, r)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (c *ChatHandlers) RemoveChannel(w http.ResponseWriter, r *http.Request) {
	requestedID, _ := strconv.Atoi(mux.Vars(r)["id"])
	user, err := c.parseCookie(r)
	if err != nil {
		c.utils.HandleError(err, w, r)
		return
	}
	err = c.Chats.DeleteChannel(user.ID, uint64(requestedID))
	if err != nil {
		c.utils.HandleError(err, w, r)
		return
	}
}

func (c *ChatHandlers) PostWorkspace(w http.ResponseWriter, r *http.Request) {
	user, err := c.parseCookie(r)
	if err != nil {
		c.utils.HandleError(err, w, r)
		return
	}
	var newWorkspace models.Workspace
	newWorkspace.Members = append(newWorkspace.Members, user.ID)
	newWorkspace.Admins = append(newWorkspace.Admins, user.ID)
	newWorkspace.CreatorID = user.ID
	err = easyjson.UnmarshalFromReader(r.Body, &newWorkspace)
	if err != nil {
		c.utils.HandleError(models.NewClientError(err, http.StatusBadRequest,
			"Bad request: malformed data:("), w, r)
		return
	}

	id, err := c.Chats.CreateWorkspace(&newWorkspace)
	if err != nil {
		c.utils.HandleError(err, w, r)
		return
	}
	newWorkspace.ID = id
	jsonResponse, err := easyjson.Marshal(newWorkspace)
	if err != nil {
		c.utils.HandleError(err, w, r)
		return
	}
	_, err = w.Write(jsonResponse)
	if err != nil {
		c.utils.LogError(err, r)
	}
	w.WriteHeader(http.StatusOK)
}

func (c *ChatHandlers) GetWorkspaceById(w http.ResponseWriter, r *http.Request) {
	requestedID, _ := strconv.Atoi(mux.Vars(r)["id"])
	user, err := c.parseCookie(r)
	if err != nil {
		c.utils.HandleError(err, w, r)
		return
	}
	workspace, err := c.Chats.GetWorkspaceByID(user.ID, uint64(requestedID))
	if err != nil {
		c.utils.HandleError(err, w, r)
		return
	}
	jsonWorkspace, err := easyjson.Marshal(workspace)
	if err != nil {
		c.utils.LogError(err, r)
	}
	_, err = w.Write(jsonWorkspace)
	if err != nil {
		c.utils.LogError(err, r)
	}
}

func (c *ChatHandlers) EditWorkspace(w http.ResponseWriter, r *http.Request) {
	user, err := c.parseCookie(r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	var newWorkspace models.Workspace
	err = easyjson.UnmarshalFromReader(r.Body, &newWorkspace)
	if err != nil {
		c.utils.HandleError(models.NewClientError(err, http.StatusBadRequest,
			"Bad request: malformed data:("), w, r)
		return
	}

	err = c.Chats.EditWorkspace(user.ID, &newWorkspace)
	if err != nil {
		c.utils.HandleError(err, w, r)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (c *ChatHandlers) RemoveWorkspace(w http.ResponseWriter, r *http.Request) {
	requestedID, _ := strconv.Atoi(mux.Vars(r)["id"])
	user, err := c.parseCookie(r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	err = c.Chats.DeleteWorkspace(user.ID, uint64(requestedID))
	if err != nil {
		c.utils.HandleError(err, w, r)
	}
}

func (c ChatHandlers) parseCookie(r *http.Request) (models.User, error) {
	//cookie, _ := r.Cookie("session_id")
	id := r.Context().Value("user_id").(uint64)
	//if err != nil {
	//	return models.User{}, models.NewClientError(err, http.StatusUnauthorized, "Bad request : not valid cookie:(")
	//}
	user, err := c.Users.GetUserByID(id)
	if err == nil {
		return user, nil
	} else {
		return user, err
	}
}

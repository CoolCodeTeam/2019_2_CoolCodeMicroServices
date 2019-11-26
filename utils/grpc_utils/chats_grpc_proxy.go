package grpc_utils

import (
	"context"
	useCase "github.com/CoolCodeTeam/2019_2_CoolCodeMicroServices/chats/usecase"
	"github.com/CoolCodeTeam/2019_2_CoolCodeMicroServices/utils/models"
)

type ChatsGRPCProxy struct {
	client ChatsServiceClient
}

func (c *ChatsGRPCProxy) CheckChannelPermission(authorID uint64, channelID uint64) (bool, error) {
	ok, err := c.client.CheckChannelPermission(context.Background(), &RequestMessage{
		UserID: authorID,
		ChatID: channelID,
	})
	if err != nil {
		return false, err
	}
	return ok.Ok, err
}

func (c *ChatsGRPCProxy) CheckChatPermission(userID uint64, chatID uint64) (bool, error) {
	ok, err := c.client.CheckChatPermission(context.Background(), &RequestMessage{
		UserID: userID,
		ChatID: chatID,
	})
	if err != nil {
		return false, err
	}
	return ok.Ok, err
}

func (c *ChatsGRPCProxy) GetChatByID(userID uint64, ID uint64) (models.Chat, error) {
	response, err := c.client.GetChatByID(context.Background(), &RequestMessage{
		UserID: userID,
		ChatID: ID,
	})
	if err != nil {
		return models.Chat{}, err
	}
	return *parseChat(response.Chat), err
}

func (c *ChatsGRPCProxy) GetChatsByUserID(ID uint64) ([]models.Chat, error) {
	response, err := c.client.GetChatsByUserID(context.Background(), &RequestMessage{
		UserID: ID,
	})
	if err != nil {
		return []models.Chat{}, err
	}
	chats := make([]models.Chat, len(response.Chats))
	for _, chat := range response.Chats {
		chats = append(chats, *parseChat(chat))
	}
	return chats, err
}

func (c *ChatsGRPCProxy) PutChat(chat *models.Chat) (uint64, error) {
	chatID, err := c.client.PutChat(context.Background(), &Chat{
		ID:            chat.ID,
		Name:          chat.Name,
		TotalMSGCount: chat.TotalMSGCount,
		Members:       chat.Members,
		LastMessage:   chat.LastMessage,
	})
	if err != nil {
		return 0, err
	}
	return chatID.Number, err
}

func (c *ChatsGRPCProxy) Contains(chat models.Chat) error {
	_, err := c.client.Contains(context.Background(), &Chat{
		ID:            chat.ID,
		Name:          chat.Name,
		TotalMSGCount: chat.TotalMSGCount,
		Members:       chat.Members,
		LastMessage:   chat.LastMessage,
	})
	return err
}

func (c *ChatsGRPCProxy) GetWorkspaceByID(userID uint64, ID uint64) (models.Workspace, error) {
	response, err := c.client.GetWorkspaceByID(context.Background(), &RequestMessage{
		UserID: userID,
		ChatID: ID,
	})
	if err != nil {
		return models.Workspace{}, err
	}
	return *parseWorkspace(response.Workspace), err
}

func (c *ChatsGRPCProxy) GetWorkspacesByUserID(ID uint64) ([]models.Workspace, error) {
	response, err := c.client.GetWorkspacesByUserID(context.Background(), &RequestMessage{
		UserID: ID,
	})
	if err != nil {
		return []models.Workspace{}, err
	}
	workspaces := make([]models.Workspace, len(response.Workspaces))
	for _, workspace := range response.Workspaces {
		workspaces = append(workspaces, *parseWorkspace(workspace))
	}
	return workspaces, err
}

func (c *ChatsGRPCProxy) CreateWorkspace(room *models.Workspace) (uint64, error) {
	id, err := c.client.CreateWorkspace(context.Background(), &Workspace{
		ID:        room.ID,
		Name:      room.Name,
		Channels:  nil,
		Members:   room.Members,
		Admins:    room.Admins,
		CreatorID: room.CreatorID,
	})
	if err != nil {
		return 0, err
	}
	return id.Number, err
}

func (c *ChatsGRPCProxy) CreateChannel(channel *models.Channel) (uint64, error) {
	id, err := c.client.CreateChannel(context.Background(), &Channel{
		ID:        channel.ID,
		Name:      channel.Name,
		Members:   channel.Members,
		Admins:    channel.Admins,
		CreatorID: channel.CreatorID,
	})
	if err != nil {
		return 0, err
	}
	return id.Number, err
}

func (c *ChatsGRPCProxy) GetChannelByID(userID uint64, ID uint64) (models.Channel, error) {
	response, err := c.client.GetChannelByID(context.Background(), &RequestMessage{
		UserID: userID,
		ChatID: ID,
	})
	if err != nil {
		return models.Channel{}, err
	}
	return *parseChannel(response.Channel), err
}

func (c *ChatsGRPCProxy) EditWorkspace(userID uint64, room *models.Workspace) error {
	_, err := c.client.EditWorkspace(context.Background(), &RequestMessage{
		UserID: userID,
		Workspace: &Workspace{
			ID:        room.ID,
			Name:      room.Name,
			Channels:  nil,
			Members:   room.Members,
			Admins:    room.Admins,
			CreatorID: room.CreatorID,
		},
	})

	return err
}

func (c *ChatsGRPCProxy) EditChannel(userID uint64, channel *models.Channel) error {
	_, err := c.client.EditChannel(context.Background(), &RequestMessage{
		UserID: userID,
		Channel: &Channel{
			ID:        channel.ID,
			Name:      channel.Name,
			Members:   channel.Members,
			Admins:    channel.Admins,
			CreatorID: channel.CreatorID,
		},
	})
	return err
}

func (c *ChatsGRPCProxy) LogoutFromWorkspace(userID uint64, workspaceID uint64) error {
	_, err := c.client.LogoutFromWorkspace(context.Background(), &RequestMessage{
		UserID: userID,
		ChatID: workspaceID,
	})
	return err
}

func (c *ChatsGRPCProxy) LogoutFromChannel(userID uint64, channelID uint64) error {
	_, err := c.client.LogoutFromWorkspace(context.Background(), &RequestMessage{
		UserID: userID,
		ChatID: channelID,
	})
	return err
}

func (c *ChatsGRPCProxy) DeleteWorkspace(userID uint64, workspaceID uint64) error {
	_, err := c.client.DeleteWorkspace(context.Background(), &RequestMessage{
		UserID: userID,
		ChatID: workspaceID,
	})
	return err
}

func (c *ChatsGRPCProxy) DeleteChannel(userID uint64, channelID uint64) error {
	_, err := c.client.DeleteChannel(context.Background(), &RequestMessage{
		UserID: userID,
		ChatID: channelID,
	})
	return err
}

func (c *ChatsGRPCProxy) DeleteChat(userID uint64, chatId uint64) error {
	_, err := c.client.DeleteChat(context.Background(), &RequestMessage{
		UserID: userID,
		ChatID: chatId,
	})
	return err
}

func NewChatsGRPCProxy(client ChatsServiceClient) useCase.ChatsUseCase {
	return &ChatsGRPCProxy{client: client}
}

func parseChannel(channel *Channel) *models.Channel {
	return &models.Channel{
		ID:            channel.ID,
		Name:          channel.Name,
		TotalMSGCount: channel.TotalMSGCount,
		Members:       channel.Members,
		Admins:        channel.Admins,
		WorkspaceID:   channel.WorkspaceID,
		CreatorID:     channel.CreatorID,
	}
}

func parseChat(chat *Chat) *models.Chat {
	return &models.Chat{
		ID:            chat.ID,
		Name:          chat.Name,
		TotalMSGCount: chat.TotalMSGCount,
		Members:       chat.Members,
		LastMessage:   chat.LastMessage,
	}
}

func parseWorkspace(workspace *Workspace) *models.Workspace {
	channels := make([]*models.Channel, len(workspace.Channels))
	for _, chat := range workspace.Channels {
		channels = append(channels, &models.Channel{
			ID:            chat.ID,
			Name:          chat.Name,
			TotalMSGCount: chat.TotalMSGCount,
			Members:       chat.Members,
			Admins:        chat.Admins,
			WorkspaceID:   chat.WorkspaceID,
			CreatorID:     chat.CreatorID,
		})
	}

	return &models.Workspace{
		ID:        workspace.ID,
		Name:      workspace.Name,
		Channels:  channels,
		Members:   workspace.Members,
		Admins:    workspace.Admins,
		CreatorID: 0,
	}
}

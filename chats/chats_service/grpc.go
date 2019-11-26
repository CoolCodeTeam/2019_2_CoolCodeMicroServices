package chats_service

import (
	"context"
	useCase "github.com/CoolCodeTeam/2019_2_CoolCodeMicroServices/chats/usecase"
	"github.com/CoolCodeTeam/2019_2_CoolCodeMicroServices/utils/grpc_utils"
	"github.com/CoolCodeTeam/2019_2_CoolCodeMicroServices/utils/models"
	"github.com/sirupsen/logrus"
)

type ChatsServiceImpl struct {
	UseCase useCase.ChatsUseCase
}

func (c ChatsServiceImpl) CheckChatPermission(ctx context.Context, r *grpc_utils.RequestMessage) (*grpc_utils.ResponseMessage, error) {
	logger := logrus.WithFields(logrus.Fields{
		"method":  "grpc_GetChatPermission",
		"user_id": r.UserID,
		"chat_id": r.ChatID,
	})

	ok, err := c.UseCase.CheckChatPermission(r.UserID, r.ChatID)
	if err != nil {
		logger.Errorf("can not get chat permission: %s", err)
		return nil, err
	}

	logger.Info("successful")
	return &grpc_utils.ResponseMessage{
		Ok: ok,
	}, nil
}

func (c ChatsServiceImpl) CheckChannelPermission(ctx context.Context, r *grpc_utils.RequestMessage) (*grpc_utils.ResponseMessage, error) {
	logger := logrus.WithFields(logrus.Fields{
		"method":  "grpc_GetChannelPermission",
		"user_id": r.UserID,
		"chat_id": r.ChatID,
	})

	ok, err := c.UseCase.CheckChannelPermission(r.UserID, r.ChatID)
	if err != nil {
		logger.Errorf("can not get chat permission: %s", err)
		return nil, err
	}

	logger.Info("successful")
	return &grpc_utils.ResponseMessage{
		Ok: ok,
	}, nil
}

func (c ChatsServiceImpl) GetChatByID(ctx context.Context, r *grpc_utils.RequestMessage) (*grpc_utils.ResponseMessage, error) {
	logger := logrus.WithFields(logrus.Fields{
		"method":  "grpc_GetChatByID",
		"chat_id": r.ChatID,
	})

	chat, err := c.UseCase.GetChatByID(r.UserID, r.ChatID)
	if err != nil {
		logger.Errorf("can not get chat: %s", err)
		return nil, err
	}

	logger.Info("successful")
	return &grpc_utils.ResponseMessage{
		Chat: encodeChat(chat),
	}, nil
}

func (c ChatsServiceImpl) GetChatsByUserID(ctx context.Context, r *grpc_utils.RequestMessage) (*grpc_utils.ResponseMessage, error) {
	logger := logrus.WithFields(logrus.Fields{
		"method":  "grpc_GetChatsByUserID",
		"chat_id": r.ChatID,
	})

	chats, err := c.UseCase.GetChatsByUserID(r.UserID)
	if err != nil {
		logger.Errorf("can not get chats: %s", err)
		return nil, err
	}
	chatsResponse := make([]*grpc_utils.Chat, len(chats))

	for _, chat := range chats {
		chatsResponse = append(chatsResponse, encodeChat(chat))
	}
	return &grpc_utils.ResponseMessage{
		Chats: chatsResponse,
	}, nil
}

func (c ChatsServiceImpl) PutChat(ctx context.Context, chat *grpc_utils.Chat) (*grpc_utils.ResponseMessage, error) {
	logger := logrus.WithFields(logrus.Fields{
		"method": "grpc_PutChat",
	})

	id, err := c.UseCase.PutChat(parseChat(chat))
	if err != nil {
		logger.Errorf("can not put chat: %s", err)
		return nil, err
	}

	return &grpc_utils.ResponseMessage{
		Number: id,
	}, nil
}

func (c ChatsServiceImpl) Contains(ctx context.Context, chat *grpc_utils.Chat) (*grpc_utils.EmptyChats, error) {
	panic("implement me")
}

func (c ChatsServiceImpl) GetWorkspaceByID(ctx context.Context, r *grpc_utils.RequestMessage) (*grpc_utils.ResponseMessage, error) {
	logger := logrus.WithFields(logrus.Fields{
		"method":  "grpc_GetWorkspaceByID",
		"chat_id": r.ChatID,
	})

	workspace, err := c.UseCase.GetWorkspaceByID(r.UserID, r.ChatID)
	if err != nil {
		logger.Errorf("can not get workspace: %s", err)
		return nil, err
	}

	logger.Info("successful")
	return &grpc_utils.ResponseMessage{
		Workspace: encodeWorkspace(workspace),
	}, nil
}

func (c ChatsServiceImpl) GetWorkspacesByUserID(ctx context.Context, r *grpc_utils.RequestMessage) (*grpc_utils.ResponseMessage, error) {
	logger := logrus.WithFields(logrus.Fields{
		"method":  "grpc_GetWorkspacesByUserID",
		"chat_id": r.ChatID,
	})

	workspaces, err := c.UseCase.GetWorkspacesByUserID(r.UserID)
	if err != nil {
		logger.Errorf("can not get workspaces: %s", err)
		return nil, err
	}

	workspaceResponse := make([]*grpc_utils.Workspace, len(workspaces))
	for _, workspace := range workspaces {
		workspaceResponse = append(workspaceResponse, encodeWorkspace(workspace))
	}
	return &grpc_utils.ResponseMessage{
		Workspaces: workspaceResponse,
	}, nil
}

func (c ChatsServiceImpl) CreateWorkspace(ctx context.Context, r *grpc_utils.Workspace) (*grpc_utils.ResponseMessage, error) {
	logger := logrus.WithFields(logrus.Fields{
		"method": "grpc_CreateWorkspace",
	})

	id, err := c.UseCase.CreateWorkspace(parseWorkspace(r))
	if err != nil {
		logger.Errorf("can not create workspaces: %s", err)
		return nil, err
	}

	return &grpc_utils.ResponseMessage{
		Number: id,
	}, nil
}

func (c ChatsServiceImpl) CreateChannel(ctx context.Context, r *grpc_utils.Channel) (*grpc_utils.ResponseMessage, error) {
	logger := logrus.WithFields(logrus.Fields{
		"method": "grpc_CreateChannel",
	})

	id, err := c.UseCase.CreateChannel(parseChannel(r))
	if err != nil {
		logger.Errorf("can not create channel: %s", err)
		return nil, err
	}

	return &grpc_utils.ResponseMessage{
		Number: id,
	}, nil
}

func (c ChatsServiceImpl) GetChannelByID(ctx context.Context, r *grpc_utils.RequestMessage) (*grpc_utils.ResponseMessage, error) {
	logger := logrus.WithFields(logrus.Fields{
		"method":  "grpc_GetChannelByID",
		"chat_id": r.ChatID,
	})

	channel, err := c.UseCase.GetChannelByID(r.UserID, r.ChatID)
	if err != nil {
		logger.Errorf("can not get channel: %s", err)
		return nil, err
	}

	logger.Info("successful")
	return &grpc_utils.ResponseMessage{
		Channel: encodeChannel(channel),
	}, nil
}

func (c ChatsServiceImpl) EditWorkspace(ctx context.Context, r *grpc_utils.RequestMessage) (*grpc_utils.EmptyChats, error) {
	logger := logrus.WithFields(logrus.Fields{
		"method": "grpc_EditWorkspace",
	})

	err := c.UseCase.EditWorkspace(r.UserID, parseWorkspace(r.Workspace))
	if err != nil {
		logger.Errorf("can not edit workspaces: %s", err)
		return nil, err
	}

	return &grpc_utils.EmptyChats{}, nil
}

func (c ChatsServiceImpl) EditChannel(ctx context.Context, r *grpc_utils.RequestMessage) (*grpc_utils.EmptyChats, error) {
	logger := logrus.WithFields(logrus.Fields{
		"method": "grpc_EditChannel",
	})

	err := c.UseCase.EditChannel(r.UserID, parseChannel(r.Channel))
	if err != nil {
		logger.Errorf("can not edit chanel: %s", err)
		return nil, err
	}

	return &grpc_utils.EmptyChats{}, nil
}

func (c ChatsServiceImpl) LogoutFromWorkspace(ctx context.Context, r *grpc_utils.RequestMessage) (*grpc_utils.EmptyChats, error) {
	logger := logrus.WithFields(logrus.Fields{
		"method": "grpc_LogoutWorkspace",
	})

	err := c.UseCase.LogoutFromWorkspace(r.UserID, r.ChatID)
	if err != nil {
		logger.Errorf("can not logout from workspace: %s", err)
		return nil, err
	}

	return &grpc_utils.EmptyChats{}, nil
}

func (c ChatsServiceImpl) LogoutFromChannel(ctx context.Context, r *grpc_utils.RequestMessage) (*grpc_utils.EmptyChats, error) {
	logger := logrus.WithFields(logrus.Fields{
		"method": "grpc_LogoutChannel",
	})

	err := c.UseCase.LogoutFromChannel(r.UserID, r.ChatID)
	if err != nil {
		logger.Errorf("can not logout from channel: %s", err)
		return nil, err
	}

	return &grpc_utils.EmptyChats{}, nil
}

func (c ChatsServiceImpl) DeleteWorkspace(ctx context.Context, r *grpc_utils.RequestMessage) (*grpc_utils.EmptyChats, error) {
	logger := logrus.WithFields(logrus.Fields{
		"method": "grpc_DeleteWorkspace",
	})

	err := c.UseCase.DeleteWorkspace(r.UserID, r.ChatID)
	if err != nil {
		logger.Errorf("can not delete workspace: %s", err)
		return nil, err
	}

	return &grpc_utils.EmptyChats{}, nil
}

func (c ChatsServiceImpl) DeleteChannel(ctx context.Context, r *grpc_utils.RequestMessage) (*grpc_utils.EmptyChats, error) {
	logger := logrus.WithFields(logrus.Fields{
		"method": "grpc_DeleteChannel",
	})

	err := c.UseCase.DeleteChannel(r.UserID, r.ChatID)
	if err != nil {
		logger.Errorf("can not delete channel: %s", err)
		return nil, err
	}

	return &grpc_utils.EmptyChats{}, nil
}

func (c ChatsServiceImpl) DeleteChat(ctx context.Context, r *grpc_utils.RequestMessage) (*grpc_utils.EmptyChats, error) {
	logger := logrus.WithFields(logrus.Fields{
		"method": "grpc_DeleteChat",
	})

	err := c.UseCase.DeleteChat(r.UserID, r.ChatID)
	if err != nil {
		logger.Errorf("can not delete chat: %s", err)
		return nil, err
	}

	return &grpc_utils.EmptyChats{}, nil
}

func NewGRPCChatsService(useCase useCase.ChatsUseCase) grpc_utils.ChatsServiceServer {
	return &ChatsServiceImpl{UseCase: useCase}
}

func encodeChat(chat models.Chat) *grpc_utils.Chat {
	return &grpc_utils.Chat{
		ID:            chat.ID,
		Name:          chat.Name,
		TotalMSGCount: chat.TotalMSGCount,
		Members:       chat.Members,
		LastMessage:   chat.LastMessage,
	}
}

func encodeChannel(channel models.Channel) *grpc_utils.Channel {
	return &grpc_utils.Channel{
		ID:            channel.ID,
		Name:          channel.Name,
		TotalMSGCount: channel.TotalMSGCount,
		Members:       channel.Members,
		Admins:        channel.Admins,
		WorkspaceID:   channel.WorkspaceID,
		CreatorID:     channel.CreatorID,
	}
}

func encodeWorkspace(workspace models.Workspace) *grpc_utils.Workspace {
	channels := make([]*grpc_utils.Channel, len(workspace.Channels))
	for _, channel := range workspace.Channels {
		channels = append(channels, encodeChannel(*channel))
	}
	return &grpc_utils.Workspace{
		ID:        workspace.ID,
		Name:      workspace.Name,
		Channels:  channels,
		Members:   workspace.Members,
		Admins:    workspace.Admins,
		CreatorID: workspace.CreatorID,
	}
}

func parseChat(chat *grpc_utils.Chat) *models.Chat {
	return &models.Chat{
		ID:            chat.ID,
		Name:          chat.Name,
		TotalMSGCount: chat.TotalMSGCount,
		Members:       chat.Members,
		LastMessage:   chat.LastMessage,
	}
}

func parseWorkspace(workspace *grpc_utils.Workspace) *models.Workspace {
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

func parseChannel(channel *grpc_utils.Channel) *models.Channel {
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

package grpc_utils

import (
	"errors"
	useCase "github.com/CoolCodeTeam/2019_2_CoolCodeMicroServices/notifications/usecase"
	"github.com/CoolCodeTeam/2019_2_CoolCodeMicroServices/utils/models"
)

type NotificationsGRPCProxy struct {
	client NotificationsServiceClient
}

func (n NotificationsGRPCProxy) OpenConn(ID uint64) (*models.WebSocketHub, error) {
	return &models.WebSocketHub{}, errors.New("Not implemented")
}

func (n NotificationsGRPCProxy) SendMessage(chatID uint64, message []byte) error {
	return nil
}

func NewNotificationsGRPCProxy(client NotificationsServiceClient) useCase.NotificationUseCase {
	return &NotificationsGRPCProxy{client: client}
}

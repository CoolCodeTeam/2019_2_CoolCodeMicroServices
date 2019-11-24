package grpc_utils

import (
	"context"
	"errors"
	useCase "github.com/go-park-mail-ru/2019_2_CoolCodeMicroServices/notifications/usecase"
	"github.com/go-park-mail-ru/2019_2_CoolCodeMicroServices/utils/models"
)

type NotificationsGRPCProxy struct {
	client NotificationsServiceClient
}

func (n NotificationsGRPCProxy) OpenConn(ID uint64) (*models.WebSocketHub, error) {
	return &models.WebSocketHub{}, errors.New("Not implemented")
}

func (n NotificationsGRPCProxy) SendMessage(chatID uint64, message []byte) error {
	_, err := n.client.SendMessage(context.Background(), &RequestMessageNotification{
		ChatID:  chatID,
		Message: message,
	})
	return err
}

func NewNotificationsGRPCProxy(client NotificationsServiceClient) useCase.NotificationUseCase {
	return &NotificationsGRPCProxy{client: client}
}

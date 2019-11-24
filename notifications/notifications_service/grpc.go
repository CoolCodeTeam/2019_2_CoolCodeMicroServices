package notifications_service

import (
	"context"
	useCase "github.com/go-park-mail-ru/2019_2_CoolCodeMicroServices/notifications/usecase"
	"github.com/go-park-mail-ru/2019_2_CoolCodeMicroServices/utils/grpc_utils"
)

type NotificationsServiceImpl struct {
	UseCase useCase.NotificationUseCase
}

func (n *NotificationsServiceImpl) SendMessage(ctx context.Context, r *grpc_utils.RequestMessageNotification) (*grpc_utils.EmptyNotification, error) {
	chatID := r.ChatID
	body := r.Message
	err := n.UseCase.SendMessage(chatID, body)
	return &grpc_utils.EmptyNotification{}, err
}

func NewNotificationsGRPCService(usecase useCase.NotificationUseCase) grpc_utils.NotificationsServiceServer {
	return &NotificationsServiceImpl{UseCase: usecase}
}

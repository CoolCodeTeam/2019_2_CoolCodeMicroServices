package useCase

import (
	"github.com/CoolCodeTeam/2019_2_CoolCodeMicroServices/notifications/repository"
	"github.com/CoolCodeTeam/2019_2_CoolCodeMicroServices/utils/models"
)

//go:generate moq -out notifications_ucase_mock.go . NotificationUseCase
type NotificationUseCase interface {
	OpenConn(ID uint64) (*models.WebSocketHub, error)
	SendMessage(chatID uint64, message []byte) error
}

type NotificationUseCaseImpl struct {
	notificationRepository repository.NotificationRepository
}

func NewNotificationUseCase() NotificationUseCase {
	return &NotificationUseCaseImpl{notificationRepository: repository.NewArrayRepo()}
}

func (u *NotificationUseCaseImpl) OpenConn(ID uint64) (*models.WebSocketHub, error) {
	return u.notificationRepository.GetNotificationHub(ID), nil
}

func (u *NotificationUseCaseImpl) SendMessage(chatID uint64, message []byte) error {
	hub := u.notificationRepository.GetNotificationHub(chatID)
	if len(hub.Clients) > 0 {
		hub.BroadcastChan <- message
	}
	return nil
}

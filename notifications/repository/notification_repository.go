package repository

import "github.com/CoolCodeTeam/2019_2_CoolCodeMicroServices/utils/models"

//go:generate moq -out notifications_repo_mock.go . NotificationRepository
type NotificationRepository interface {
	GetNotificationHub(chatID uint64) *models.WebSocketHub
}

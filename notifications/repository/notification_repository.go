package repository

import "github.com/CoolCodeTeam/2019_2_CoolCodeMicroServices/utils/models"

type NotificationRepository interface {
	GetNotificationHub(chatID uint64) *models.WebSocketHub
}

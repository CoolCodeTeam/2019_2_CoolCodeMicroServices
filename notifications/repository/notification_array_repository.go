package repository

import (
	"github.com/CoolCodeTeam/2019_2_CoolCodeMicroServices/utils/models"
	"sync"
)

type NotificationArrayRepository struct {
	Hubs  map[uint64]*models.WebSocketHub
	mutex sync.Mutex
}

func (n *NotificationArrayRepository) GetNotificationHub(chatID uint64) *models.WebSocketHub {
	n.mutex.Lock()
	defer n.mutex.Unlock()
	if hub, ok := n.Hubs[chatID]; ok {
		return hub
	}
	n.Hubs[chatID] = models.NewHub()
	return n.Hubs[chatID]
}

func NewArrayRepo() NotificationRepository {
	return &NotificationArrayRepository{Hubs: make(map[uint64]*models.WebSocketHub), mutex: sync.Mutex{}}
}

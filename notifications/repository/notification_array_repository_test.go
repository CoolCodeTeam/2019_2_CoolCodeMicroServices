package repository

import (
	"github.com/CoolCodeTeam/2019_2_CoolCodeMicroServices/utils/models"
	"github.com/gorilla/websocket"
	"reflect"
	"sync"
	"testing"
)

func TestNewArrayRepo(t *testing.T) {
	tests := []struct {
		name string
		want NotificationRepository
	}{
		{
			name: "DefaultTest",
			want: &NotificationArrayRepository{Hubs: make(map[uint64]*models.WebSocketHub), mutex: sync.Mutex{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewArrayRepo(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewArrayRepo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNotificationArrayRepository_GetNotificationHub(t *testing.T) {
	type fields struct {
		Hubs  map[uint64]*models.WebSocketHub
		mutex sync.Mutex
	}
	type args struct {
		chatID uint64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *models.WebSocketHub
	}{
		{
			name: "Default test",
			fields: fields{
				Hubs: make(map[uint64]*models.WebSocketHub), mutex: sync.Mutex{},
			},
			args: args{chatID: 0},
			want: &models.WebSocketHub{
				Clients:          make(map[string]*models.WebSocketClient),
				AddClientChan:    make(chan *websocket.Conn),
				RemoveClientChan: make(chan *websocket.Conn),
				BroadcastChan:    make(chan []byte),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &NotificationArrayRepository{
				Hubs:  tt.fields.Hubs,
				mutex: tt.fields.mutex,
			}
			if got := n.GetNotificationHub(tt.args.chatID); !reflect.DeepEqual(got, got) {
				t.Errorf("GetNotificationHub() = %v, want %v", got, tt.want)
			}
		})
	}
}

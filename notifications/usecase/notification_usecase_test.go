package useCase

import (
	"github.com/CoolCodeTeam/2019_2_CoolCodeMicroServices/notifications/repository"
	"github.com/CoolCodeTeam/2019_2_CoolCodeMicroServices/utils/models"
	"reflect"
	"testing"
)

func TestNewNotificationUseCase(t *testing.T) {
	tests := []struct {
		name string
		want NotificationUseCase
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewNotificationUseCase(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewNotificationUseCase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNotificationUseCaseImpl_OpenConn(t *testing.T) {
	type fields struct {
		notificationRepository repository.NotificationRepository
	}
	type args struct {
		ID uint64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *models.WebSocketHub
		wantErr bool
	}{
		{
			name: "DefaultTest",
			fields: fields{notificationRepository: &repository.NotificationRepositoryMock{
				GetNotificationHubFunc: func(chatID uint64) *models.WebSocketHub {
					return &models.WebSocketHub{}
				},
			}},
			args:    args{ID: 0},
			want:    &models.WebSocketHub{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &NotificationUseCaseImpl{
				notificationRepository: tt.fields.notificationRepository,
			}
			got, err := u.OpenConn(tt.args.ID)
			if (err != nil) != tt.wantErr {
				t.Errorf("OpenConn() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OpenConn() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNotificationUseCaseImpl_SendMessage(t *testing.T) {
	type fields struct {
		notificationRepository repository.NotificationRepository
	}
	type args struct {
		chatID  uint64
		message []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "EmptyUsers",
			fields: fields{notificationRepository: &repository.NotificationRepositoryMock{
				GetNotificationHubFunc: func(chatID uint64) *models.WebSocketHub {
					return &models.WebSocketHub{}
				},
			}},
			args: args{
				chatID:  0,
				message: nil,
			},
			wantErr: false,
		},
		{
			name: "SuccessTest",
			fields: fields{notificationRepository: &repository.NotificationRepositoryMock{
				GetNotificationHubFunc: func(chatID uint64) *models.WebSocketHub {
					return &models.WebSocketHub{Clients: map[string]*models.WebSocketClient{
						"test": &models.WebSocketClient{},
					},
						BroadcastChan: make(chan []byte, 1)}
				},
			}},
			args: args{
				chatID:  0,
				message: nil,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &NotificationUseCaseImpl{
				notificationRepository: tt.fields.notificationRepository,
			}
			if err := u.SendMessage(tt.args.chatID, tt.args.message); (err != nil) != tt.wantErr {
				t.Errorf("SendMessage() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNotificationUseCaseMock_OpenConn(t *testing.T) {
	type fields struct {
		OpenConnFunc    func(ID uint64) (*models.WebSocketHub, error)
		SendMessageFunc func(chatID uint64, message []byte) error
		calls           struct {
			// OpenConn holds details about calls to the OpenConn method.
			OpenConn []struct {
				// ID is the ID argument value.
				ID uint64
			}
			// SendMessage holds details about calls to the SendMessage method.
			SendMessage []struct {
				// ChatID is the chatID argument value.
				ChatID uint64
				// Message is the message argument value.
				Message []byte
			}
		}
	}
	type args struct {
		ID uint64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *models.WebSocketHub
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := &NotificationUseCaseMock{
				OpenConnFunc:    tt.fields.OpenConnFunc,
				SendMessageFunc: tt.fields.SendMessageFunc,
				calls:           tt.fields.calls,
			}
			got, err := mock.OpenConn(tt.args.ID)
			if (err != nil) != tt.wantErr {
				t.Errorf("OpenConn() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OpenConn() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNotificationUseCaseMock_OpenConnCalls(t *testing.T) {
	type fields struct {
		OpenConnFunc    func(ID uint64) (*models.WebSocketHub, error)
		SendMessageFunc func(chatID uint64, message []byte) error
		calls           struct {
			// OpenConn holds details about calls to the OpenConn method.
			OpenConn []struct {
				// ID is the ID argument value.
				ID uint64
			}
			// SendMessage holds details about calls to the SendMessage method.
			SendMessage []struct {
				// ChatID is the chatID argument value.
				ChatID uint64
				// Message is the message argument value.
				Message []byte
			}
		}
	}
	tests := []struct {
		name   string
		fields fields
		want   []struct {
			ID uint64
		}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := &NotificationUseCaseMock{
				OpenConnFunc:    tt.fields.OpenConnFunc,
				SendMessageFunc: tt.fields.SendMessageFunc,
				calls:           tt.fields.calls,
			}
			if got := mock.OpenConnCalls(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OpenConnCalls() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNotificationUseCaseMock_SendMessage(t *testing.T) {
	type fields struct {
		OpenConnFunc    func(ID uint64) (*models.WebSocketHub, error)
		SendMessageFunc func(chatID uint64, message []byte) error
		calls           struct {
			// OpenConn holds details about calls to the OpenConn method.
			OpenConn []struct {
				// ID is the ID argument value.
				ID uint64
			}
			// SendMessage holds details about calls to the SendMessage method.
			SendMessage []struct {
				// ChatID is the chatID argument value.
				ChatID uint64
				// Message is the message argument value.
				Message []byte
			}
		}
	}
	type args struct {
		chatID  uint64
		message []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := &NotificationUseCaseMock{
				OpenConnFunc:    tt.fields.OpenConnFunc,
				SendMessageFunc: tt.fields.SendMessageFunc,
				calls:           tt.fields.calls,
			}
			if err := mock.SendMessage(tt.args.chatID, tt.args.message); (err != nil) != tt.wantErr {
				t.Errorf("SendMessage() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNotificationUseCaseMock_SendMessageCalls(t *testing.T) {
	type fields struct {
		OpenConnFunc    func(ID uint64) (*models.WebSocketHub, error)
		SendMessageFunc func(chatID uint64, message []byte) error
		calls           struct {
			// OpenConn holds details about calls to the OpenConn method.
			OpenConn []struct {
				// ID is the ID argument value.
				ID uint64
			}
			// SendMessage holds details about calls to the SendMessage method.
			SendMessage []struct {
				// ChatID is the chatID argument value.
				ChatID uint64
				// Message is the message argument value.
				Message []byte
			}
		}
	}
	tests := []struct {
		name   string
		fields fields
		want   []struct {
			ChatID  uint64
			Message []byte
		}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := &NotificationUseCaseMock{
				OpenConnFunc:    tt.fields.OpenConnFunc,
				SendMessageFunc: tt.fields.SendMessageFunc,
				calls:           tt.fields.calls,
			}
			if got := mock.SendMessageCalls(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SendMessageCalls() = %v, want %v", got, tt.want)
			}
		})
	}
}

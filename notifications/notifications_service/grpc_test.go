package notifications_service

import (
	"context"
	"github.com/CoolCodeTeam/2019_2_CoolCodeMicroServices/notifications/usecase"
	"github.com/CoolCodeTeam/2019_2_CoolCodeMicroServices/utils/grpc_utils"
	"reflect"
	"testing"
)

func TestNewNotificationsGRPCService(t *testing.T) {
	type args struct {
		usecase useCase.NotificationUseCase
	}
	tests := []struct {
		name string
		args args
		want grpc_utils.NotificationsServiceServer
	}{
		{
			name: "DefaultTest",
			args: args{usecase: &useCase.NotificationUseCaseMock{}},
			want: NewNotificationsGRPCService(&useCase.NotificationUseCaseMock{}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewNotificationsGRPCService(tt.args.usecase); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewNotificationsGRPCService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNotificationsServiceImpl_SendMessage(t *testing.T) {
	type fields struct {
		UseCase useCase.NotificationUseCase
	}
	type args struct {
		ctx context.Context
		r   *grpc_utils.RequestMessageNotification
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *grpc_utils.EmptyNotification
		wantErr bool
	}{
		{
			name: "DefaultTest",
			fields: fields{UseCase: &useCase.NotificationUseCaseMock{
				SendMessageFunc: func(chatID uint64, message []byte) error {
					return nil
				},
			}},
			args: args{
				ctx: context.Background(),
				r:   &grpc_utils.RequestMessageNotification{},
			},
			want:    &grpc_utils.EmptyNotification{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &NotificationsServiceImpl{
				UseCase: tt.fields.UseCase,
			}
			got, err := n.SendMessage(tt.args.ctx, tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("SendMessage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SendMessage() got = %v, want %v", got, tt.want)
			}
		})
	}
}

package grpc_utils

import (
	"github.com/CoolCodeTeam/2019_2_CoolCodeMicroServices/notifications/usecase"
	"github.com/CoolCodeTeam/2019_2_CoolCodeMicroServices/utils/models"
	"google.golang.org/grpc"
	"reflect"
	"testing"
)

func TestNewNotificationsGRPCProxy(t *testing.T) {
	type args struct {
		client NotificationsServiceClient
	}
	tests := []struct {
		name string
		args args
		want useCase.NotificationUseCase
	}{
		{
			name: "DefaultTest",
			args: args{
				client: NewNotificationsServiceClient(&grpc.ClientConn{}),
			},
			want: &NotificationsGRPCProxy{client: NewNotificationsServiceClient(&grpc.ClientConn{})},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewNotificationsGRPCProxy(tt.args.client); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewNotificationsGRPCProxy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNotificationsGRPCProxy_OpenConn(t *testing.T) {
	type fields struct {
		client NotificationsServiceClient
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
			name:    "DefaultTest",
			wantErr: true,
			want:    &models.WebSocketHub{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := NotificationsGRPCProxy{
				client: tt.fields.client,
			}
			got, err := n.OpenConn(tt.args.ID)
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

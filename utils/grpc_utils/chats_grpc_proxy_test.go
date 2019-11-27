package grpc_utils

import (
	"context"
	"errors"
	useCase "github.com/CoolCodeTeam/2019_2_CoolCodeMicroServices/chats/usecase"
	"github.com/CoolCodeTeam/2019_2_CoolCodeMicroServices/utils/models"
	"google.golang.org/grpc"
	"reflect"
	"testing"
)

var internalError = errors.New("Internal error")

func TestChatsGRPCProxy_CheckChannelPermission(t *testing.T) {
	type fields struct {
		client ChatsServiceClient
	}
	type args struct {
		authorID  uint64
		channelID uint64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
	}{
		{
			name: "ErrorTest",
			fields: fields{client: &ChatsServiceClientMock{
				CheckChannelPermissionFunc: func(ctx context.Context, in *RequestMessage, opts ...grpc.CallOption) (message *ResponseMessage, e error) {
					return &ResponseMessage{}, internalError
				},
			}},
			args:    args{},
			want:    false,
			wantErr: true,
		},
		{
			name: "SuccessTest",
			fields: fields{client: &ChatsServiceClientMock{
				CheckChannelPermissionFunc: func(ctx context.Context, in *RequestMessage, opts ...grpc.CallOption) (message *ResponseMessage, e error) {
					return &ResponseMessage{
						Ok: true,
					}, nil
				},
			}},
			args:    args{},
			want:    true,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &ChatsGRPCProxy{
				client: tt.fields.client,
			}
			got, err := c.CheckChannelPermission(tt.args.authorID, tt.args.channelID)
			if (err != nil) != tt.wantErr {
				t.Errorf("CheckChannelPermission() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CheckChannelPermission() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestChatsGRPCProxy_CheckChatPermission(t *testing.T) {
	type fields struct {
		client ChatsServiceClient
	}
	type args struct {
		userID uint64
		chatID uint64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
	}{
		{
			name: "ErrorTest",
			fields: fields{client: &ChatsServiceClientMock{
				CheckChatPermissionFunc: func(ctx context.Context, in *RequestMessage, opts ...grpc.CallOption) (message *ResponseMessage, e error) {
					return &ResponseMessage{}, internalError
				},
			}},
			args:    args{},
			want:    false,
			wantErr: true,
		},
		{
			name: "SuccessTest",
			fields: fields{client: &ChatsServiceClientMock{
				CheckChatPermissionFunc: func(ctx context.Context, in *RequestMessage, opts ...grpc.CallOption) (message *ResponseMessage, e error) {
					return &ResponseMessage{
						Ok: true,
					}, nil
				},
			}},
			args:    args{},
			want:    true,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &ChatsGRPCProxy{
				client: tt.fields.client,
			}
			got, err := c.CheckChatPermission(tt.args.userID, tt.args.chatID)
			if (err != nil) != tt.wantErr {
				t.Errorf("CheckChatPermission() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CheckChatPermission() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestChatsGRPCProxy_Contains(t *testing.T) {
	type fields struct {
		client ChatsServiceClient
	}
	type args struct {
		chat models.Chat
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "DefaultTest",
			fields: fields{
				client: &ChatsServiceClientMock{
					ContainsFunc: func(ctx context.Context, in *Chat, opts ...grpc.CallOption) (chats *EmptyChats, e error) {
						return &EmptyChats{}, nil
					},
				},
			},
			args:    args{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &ChatsGRPCProxy{
				client: tt.fields.client,
			}
			if err := c.Contains(tt.args.chat); (err != nil) != tt.wantErr {
				t.Errorf("Contains() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestChatsGRPCProxy_CreateChannel(t *testing.T) {
	type fields struct {
		client ChatsServiceClient
	}
	type args struct {
		channel *models.Channel
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    uint64
		wantErr bool
	}{
		{
			name: "ErrorTest",
			fields: fields{client: &ChatsServiceClientMock{
				CreateChannelFunc: func(ctx context.Context, in *Channel, opts ...grpc.CallOption) (message *ResponseMessage, e error) {
					return &ResponseMessage{}, internalError
				},
			}},
			args: args{
				channel: &models.Channel{},
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "SuccessTest",
			fields: fields{client: &ChatsServiceClientMock{
				CreateChannelFunc: func(ctx context.Context, in *Channel, opts ...grpc.CallOption) (message *ResponseMessage, e error) {
					return &ResponseMessage{}, nil
				},
			}},
			args: args{
				channel: &models.Channel{},
			},
			want:    0,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &ChatsGRPCProxy{
				client: tt.fields.client,
			}
			got, err := c.CreateChannel(tt.args.channel)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateChannel() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CreateChannel() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestChatsGRPCProxy_CreateWorkspace(t *testing.T) {
	type fields struct {
		client ChatsServiceClient
	}
	type args struct {
		room *models.Workspace
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    uint64
		wantErr bool
	}{
		{
			name: "ErrorTest",
			fields: fields{client: &ChatsServiceClientMock{
				CreateWorkspaceFunc: func(ctx context.Context, in *Workspace, opts ...grpc.CallOption) (message *ResponseMessage, e error) {
					return &ResponseMessage{}, internalError
				},
			}},
			args: args{
				room: &models.Workspace{},
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "SuccessTest",
			fields: fields{client: &ChatsServiceClientMock{
				CreateWorkspaceFunc: func(ctx context.Context, in *Workspace, opts ...grpc.CallOption) (message *ResponseMessage, e error) {
					return &ResponseMessage{}, nil
				},
			}},
			args: args{
				room: &models.Workspace{},
			},
			want:    0,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &ChatsGRPCProxy{
				client: tt.fields.client,
			}
			got, err := c.CreateWorkspace(tt.args.room)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateWorkspace() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CreateWorkspace() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestChatsGRPCProxy_DeleteChannel(t *testing.T) {
	type fields struct {
		client ChatsServiceClient
	}
	type args struct {
		userID    uint64
		channelID uint64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "ErrorTest",
			fields: fields{client: &ChatsServiceClientMock{
				DeleteChannelFunc: func(ctx context.Context, in *RequestMessage, opts ...grpc.CallOption) (chats *EmptyChats, e error) {
					return &EmptyChats{}, internalError
				},
			}},
			args:    args{},
			wantErr: true,
		},
		{
			name: "SuccessTest",
			fields: fields{client: &ChatsServiceClientMock{
				DeleteChannelFunc: func(ctx context.Context, in *RequestMessage, opts ...grpc.CallOption) (chats *EmptyChats, e error) {
					return &EmptyChats{}, nil
				},
			},
			},
			args:    args{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &ChatsGRPCProxy{
				client: tt.fields.client,
			}
			if err := c.DeleteChannel(tt.args.userID, tt.args.channelID); (err != nil) != tt.wantErr {
				t.Errorf("DeleteChannel() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestChatsGRPCProxy_DeleteChat(t *testing.T) {
	type fields struct {
		client ChatsServiceClient
	}
	type args struct {
		userID uint64
		chatId uint64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "ErrorTest",
			fields: fields{client: &ChatsServiceClientMock{
				DeleteChatFunc: func(ctx context.Context, in *RequestMessage, opts ...grpc.CallOption) (chats *EmptyChats, e error) {
					return &EmptyChats{}, internalError
				},
			}},
			args:    args{},
			wantErr: true,
		},
		{
			name: "SuccessTest",
			fields: fields{client: &ChatsServiceClientMock{
				DeleteChatFunc: func(ctx context.Context, in *RequestMessage, opts ...grpc.CallOption) (chats *EmptyChats, e error) {
					return &EmptyChats{}, nil
				},
			},
			},
			args:    args{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &ChatsGRPCProxy{
				client: tt.fields.client,
			}
			if err := c.DeleteChat(tt.args.userID, tt.args.chatId); (err != nil) != tt.wantErr {
				t.Errorf("DeleteChat() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestChatsGRPCProxy_DeleteWorkspace(t *testing.T) {
	type fields struct {
		client ChatsServiceClient
	}
	type args struct {
		userID      uint64
		workspaceID uint64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "ErrorTest",
			fields: fields{client: &ChatsServiceClientMock{
				DeleteWorkspaceFunc: func(ctx context.Context, in *RequestMessage, opts ...grpc.CallOption) (chats *EmptyChats, e error) {
					return &EmptyChats{}, internalError
				},
			}},
			args:    args{},
			wantErr: true,
		},
		{
			name: "SuccessTest",
			fields: fields{client: &ChatsServiceClientMock{
				DeleteWorkspaceFunc: func(ctx context.Context, in *RequestMessage, opts ...grpc.CallOption) (chats *EmptyChats, e error) {
					return &EmptyChats{}, nil
				},
			},
			},
			args:    args{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &ChatsGRPCProxy{
				client: tt.fields.client,
			}
			if err := c.DeleteWorkspace(tt.args.userID, tt.args.workspaceID); (err != nil) != tt.wantErr {
				t.Errorf("DeleteWorkspace() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestChatsGRPCProxy_EditChannel(t *testing.T) {
	type fields struct {
		client ChatsServiceClient
	}
	type args struct {
		userID  uint64
		channel *models.Channel
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "ErrorTest",
			fields: fields{client: &ChatsServiceClientMock{
				EditChannelFunc: func(ctx context.Context, in *RequestMessage, opts ...grpc.CallOption) (chats *EmptyChats, e error) {
					return &EmptyChats{}, internalError
				},
			}},
			args: args{
				channel: &models.Channel{},
			},
			wantErr: true,
		},
		{
			name: "SuccessTest",
			fields: fields{client: &ChatsServiceClientMock{
				EditChannelFunc: func(ctx context.Context, in *RequestMessage, opts ...grpc.CallOption) (chats *EmptyChats, e error) {
					return &EmptyChats{}, nil
				},
			},
			},
			args: args{
				channel: &models.Channel{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &ChatsGRPCProxy{
				client: tt.fields.client,
			}
			if err := c.EditChannel(tt.args.userID, tt.args.channel); (err != nil) != tt.wantErr {
				t.Errorf("EditChannel() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestChatsGRPCProxy_EditWorkspace(t *testing.T) {
	type fields struct {
		client ChatsServiceClient
	}
	type args struct {
		userID uint64
		room   *models.Workspace
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "ErrorTest",
			fields: fields{client: &ChatsServiceClientMock{
				EditWorkspaceFunc: func(ctx context.Context, in *RequestMessage, opts ...grpc.CallOption) (chats *EmptyChats, e error) {
					return &EmptyChats{}, internalError
				},
			}},
			args: args{
				room: &models.Workspace{},
			},
			wantErr: true,
		},
		{
			name: "SuccessTest",
			fields: fields{client: &ChatsServiceClientMock{
				EditWorkspaceFunc: func(ctx context.Context, in *RequestMessage, opts ...grpc.CallOption) (chats *EmptyChats, e error) {
					return &EmptyChats{}, nil
				},
			},
			},
			args: args{
				room: &models.Workspace{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &ChatsGRPCProxy{
				client: tt.fields.client,
			}
			if err := c.EditWorkspace(tt.args.userID, tt.args.room); (err != nil) != tt.wantErr {
				t.Errorf("EditWorkspace() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestChatsGRPCProxy_GetChannelByID(t *testing.T) {
	type fields struct {
		client ChatsServiceClient
	}
	type args struct {
		userID uint64
		ID     uint64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    models.Channel
		wantErr bool
	}{
		{
			name: "ErrorTest",
			fields: fields{client: &ChatsServiceClientMock{
				GetChannelByIDFunc: func(ctx context.Context, in *RequestMessage, opts ...grpc.CallOption) (message *ResponseMessage, e error) {
					return &ResponseMessage{}, internalError
				},
			}},
			args:    args{},
			want:    models.Channel{},
			wantErr: true,
		},
		{
			name: "SuccessTest",
			fields: fields{client: &ChatsServiceClientMock{
				GetChannelByIDFunc: func(ctx context.Context, in *RequestMessage, opts ...grpc.CallOption) (message *ResponseMessage, e error) {
					return &ResponseMessage{
						Channel: &Channel{},
					}, nil
				},
			}},
			args:    args{},
			want:    models.Channel{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &ChatsGRPCProxy{
				client: tt.fields.client,
			}
			got, err := c.GetChannelByID(tt.args.userID, tt.args.ID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetChannelByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetChannelByID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestChatsGRPCProxy_GetChatByID(t *testing.T) {
	type fields struct {
		client ChatsServiceClient
	}
	type args struct {
		userID uint64
		ID     uint64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    models.Chat
		wantErr bool
	}{
		{
			name: "ErrorTest",
			fields: fields{client: &ChatsServiceClientMock{
				GetChatByIDFunc: func(ctx context.Context, in *RequestMessage, opts ...grpc.CallOption) (message *ResponseMessage, e error) {
					return &ResponseMessage{}, internalError
				},
			}},
			args:    args{},
			want:    models.Chat{},
			wantErr: true,
		},
		{
			name: "SuccessTest",
			fields: fields{client: &ChatsServiceClientMock{
				GetChatByIDFunc: func(ctx context.Context, in *RequestMessage, opts ...grpc.CallOption) (message *ResponseMessage, e error) {
					return &ResponseMessage{
						Chat: &Chat{},
					}, nil
				},
			}},
			args:    args{},
			want:    models.Chat{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &ChatsGRPCProxy{
				client: tt.fields.client,
			}
			got, err := c.GetChatByID(tt.args.userID, tt.args.ID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetChatByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetChatByID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestChatsGRPCProxy_GetChatsByUserID(t *testing.T) {
	type fields struct {
		client ChatsServiceClient
	}
	type args struct {
		ID uint64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []models.Chat
		wantErr bool
	}{
		{
			name: "ErrorTest",
			fields: fields{client: &ChatsServiceClientMock{
				GetChatsByUserIDFunc: func(ctx context.Context, in *RequestMessage, opts ...grpc.CallOption) (message *ResponseMessage, e error) {
					return &ResponseMessage{}, internalError
				},
			}},
			args:    args{},
			want:    []models.Chat{},
			wantErr: true,
		},
		{
			name: "SuccessTest",
			fields: fields{client: &ChatsServiceClientMock{
				GetChatsByUserIDFunc: func(ctx context.Context, in *RequestMessage, opts ...grpc.CallOption) (message *ResponseMessage, e error) {
					return &ResponseMessage{
						Chats: []*Chat{},
					}, nil
				},
			}},
			args:    args{},
			want:    []models.Chat{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &ChatsGRPCProxy{
				client: tt.fields.client,
			}
			got, err := c.GetChatsByUserID(tt.args.ID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetChatsByUserID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetChatsByUserID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestChatsGRPCProxy_GetWorkspaceByID(t *testing.T) {
	type fields struct {
		client ChatsServiceClient
	}
	type args struct {
		userID uint64
		ID     uint64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    models.Workspace
		wantErr bool
	}{
		{
			name: "ErrorTest",
			fields: fields{client: &ChatsServiceClientMock{
				GetWorkspaceByIDFunc: func(ctx context.Context, in *RequestMessage, opts ...grpc.CallOption) (message *ResponseMessage, e error) {
					return &ResponseMessage{}, internalError
				},
			}},
			args:    args{},
			want:    models.Workspace{},
			wantErr: true,
		},
		//{
		//	name: "SuccessTest",
		//	fields: fields{client: &ChatsServiceClientMock{
		//		GetWorkspaceByIDFunc: func(ctx context.Context, in *RequestMessage, opts ...grpc.CallOption) (message *ResponseMessage, e error) {
		//			return &ResponseMessage{
		//				Workspace: &Workspace{},
		//			}, nil
		//		},
		//	}},
		//	args:    args{},
		//	want:    models.Workspace{},
		//	wantErr: false,
		//},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &ChatsGRPCProxy{
				client: tt.fields.client,
			}
			got, err := c.GetWorkspaceByID(tt.args.userID, tt.args.ID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetWorkspaceByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetWorkspaceByID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestChatsGRPCProxy_GetWorkspacesByUserID(t *testing.T) {
	type fields struct {
		client ChatsServiceClient
	}
	type args struct {
		ID uint64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []models.Workspace
		wantErr bool
	}{
		{
			name: "ErrorTest",
			fields: fields{client: &ChatsServiceClientMock{
				GetWorkspacesByUserIDFunc: func(ctx context.Context, in *RequestMessage, opts ...grpc.CallOption) (message *ResponseMessage, e error) {
					return &ResponseMessage{}, internalError
				},
			}},
			args:    args{},
			want:    []models.Workspace{},
			wantErr: true,
		},
		{
			name: "SuccessTest",
			fields: fields{client: &ChatsServiceClientMock{
				GetWorkspacesByUserIDFunc: func(ctx context.Context, in *RequestMessage, opts ...grpc.CallOption) (message *ResponseMessage, e error) {
					return &ResponseMessage{
						Workspaces: []*Workspace{},
					}, nil
				},
			}},
			args:    args{},
			want:    []models.Workspace{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &ChatsGRPCProxy{
				client: tt.fields.client,
			}
			got, err := c.GetWorkspacesByUserID(tt.args.ID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetWorkspacesByUserID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetWorkspacesByUserID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestChatsGRPCProxy_LogoutFromChannel(t *testing.T) {
	type fields struct {
		client ChatsServiceClient
	}
	type args struct {
		userID    uint64
		channelID uint64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "ErrorTest",
			fields: fields{client: &ChatsServiceClientMock{
				LogoutFromChannelFunc: func(ctx context.Context, in *RequestMessage, opts ...grpc.CallOption) (chats *EmptyChats, e error) {
					return &EmptyChats{}, internalError
				},
			}},
			args:    args{},
			wantErr: true,
		},
		{
			name: "SuccessTest",
			fields: fields{client: &ChatsServiceClientMock{
				LogoutFromChannelFunc: func(ctx context.Context, in *RequestMessage, opts ...grpc.CallOption) (chats *EmptyChats, e error) {
					return &EmptyChats{}, nil
				},
			}},
			args:    args{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &ChatsGRPCProxy{
				client: tt.fields.client,
			}
			if err := c.LogoutFromChannel(tt.args.userID, tt.args.channelID); (err != nil) != tt.wantErr {
				t.Errorf("LogoutFromChannel() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestChatsGRPCProxy_LogoutFromWorkspace(t *testing.T) {
	type fields struct {
		client ChatsServiceClient
	}
	type args struct {
		userID      uint64
		workspaceID uint64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "ErrorTest",
			fields: fields{client: &ChatsServiceClientMock{
				LogoutFromWorkspaceFunc: func(ctx context.Context, in *RequestMessage, opts ...grpc.CallOption) (chats *EmptyChats, e error) {
					return &EmptyChats{}, internalError
				},
			}},
			args:    args{},
			wantErr: true,
		},
		{
			name: "SuccessTest",
			fields: fields{client: &ChatsServiceClientMock{
				LogoutFromWorkspaceFunc: func(ctx context.Context, in *RequestMessage, opts ...grpc.CallOption) (chats *EmptyChats, e error) {
					return &EmptyChats{}, nil
				},
			}},
			args:    args{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &ChatsGRPCProxy{
				client: tt.fields.client,
			}
			if err := c.LogoutFromWorkspace(tt.args.userID, tt.args.workspaceID); (err != nil) != tt.wantErr {
				t.Errorf("LogoutFromWorkspace() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestChatsGRPCProxy_PutChat(t *testing.T) {
	type fields struct {
		client ChatsServiceClient
	}
	type args struct {
		chat *models.Chat
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    uint64
		wantErr bool
	}{
		{
			name: "ErrorTest",
			fields: fields{client: &ChatsServiceClientMock{
				PutChatFunc: func(ctx context.Context, in *Chat, opts ...grpc.CallOption) (message *ResponseMessage, e error) {
					return &ResponseMessage{}, internalError
				},
			}},
			args: args{
				chat: &models.Chat{},
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "SuccessTest",
			fields: fields{client: &ChatsServiceClientMock{
				PutChatFunc: func(ctx context.Context, in *Chat, opts ...grpc.CallOption) (message *ResponseMessage, e error) {
					return &ResponseMessage{}, nil
				},
			}},
			args: args{
				chat: &models.Chat{},
			},
			want:    0,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &ChatsGRPCProxy{
				client: tt.fields.client,
			}
			got, err := c.PutChat(tt.args.chat)
			if (err != nil) != tt.wantErr {
				t.Errorf("PutChat() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("PutChat() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewChatsGRPCProxy(t *testing.T) {
	type args struct {
		client ChatsServiceClient
	}
	tests := []struct {
		name string
		args args
		want useCase.ChatsUseCase
	}{
		{
			name: "DefaulTest",
			args: args{client: &ChatsServiceClientMock{}},
			want: &ChatsGRPCProxy{client: &ChatsServiceClientMock{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewChatsGRPCProxy(tt.args.client); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewChatsGRPCProxy() = %v, want %v", got, tt.want)
			}
		})
	}
}

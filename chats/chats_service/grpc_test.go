package chats_service

import (
	"context"
	"errors"
	"github.com/go-park-mail-ru/2019_2_CoolCodeMicroServices/chats/usecase"
	"github.com/go-park-mail-ru/2019_2_CoolCodeMicroServices/utils/grpc_utils"
	"github.com/go-park-mail-ru/2019_2_CoolCodeMicroServices/utils/models"
	"reflect"
	"testing"
)

var internalError error = errors.New("Internal error")
var contextArg context.Context = context.Background()

func TestChatsServiceImpl_CheckChannelPermission(t *testing.T) {
	type fields struct {
		UseCase useCase.ChatsUseCase
	}
	type args struct {
		ctx context.Context
		r   *grpc_utils.RequestMessage
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *grpc_utils.ResponseMessage
		wantErr bool
	}{
		{
			name: "InternalErrorTest",
			fields: fields{
				UseCase: &useCase.ChatsUseCaseMock{
					CheckChannelPermissionFunc: func(authorID uint64, channelID uint64) (b bool, e error) {
						return false, internalError
					},
				}},
			args: args{
				ctx: contextArg,
				r: &grpc_utils.RequestMessage{
					UserID: 1,
					ChatID: 1,
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "SuccessTest",
			fields: fields{
				UseCase: &useCase.ChatsUseCaseMock{
					CheckChannelPermissionFunc: func(authorID uint64, channelID uint64) (b bool, e error) {
						return false, nil
					},
				}},
			args: args{
				ctx: contextArg,
				r: &grpc_utils.RequestMessage{
					UserID: 1,
					ChatID: 1,
				},
			},
			want: &grpc_utils.ResponseMessage{
				Ok: false,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := ChatsServiceImpl{
				UseCase: tt.fields.UseCase,
			}
			got, err := c.CheckChannelPermission(tt.args.ctx, tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("CheckChannelPermission() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CheckChannelPermission() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestChatsServiceImpl_CheckChatPermission(t *testing.T) {
	type fields struct {
		UseCase useCase.ChatsUseCase
	}
	type args struct {
		ctx context.Context
		r   *grpc_utils.RequestMessage
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *grpc_utils.ResponseMessage
		wantErr bool
	}{
		{
			name: "InternalErrorTest",
			fields: fields{
				UseCase: &useCase.ChatsUseCaseMock{
					CheckChatPermissionFunc: func(authorID uint64, channelID uint64) (b bool, e error) {
						return false, internalError
					},
				}},
			args: args{
				ctx: contextArg,
				r: &grpc_utils.RequestMessage{
					UserID: 1,
					ChatID: 1,
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "SuccessTest",
			fields: fields{
				UseCase: &useCase.ChatsUseCaseMock{
					CheckChatPermissionFunc: func(authorID uint64, channelID uint64) (b bool, e error) {
						return false, nil
					},
				}},
			args: args{
				ctx: contextArg,
				r: &grpc_utils.RequestMessage{
					UserID: 1,
					ChatID: 1,
				},
			},
			want: &grpc_utils.ResponseMessage{
				Ok: false,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := ChatsServiceImpl{
				UseCase: tt.fields.UseCase,
			}
			got, err := c.CheckChatPermission(tt.args.ctx, tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("CheckChatPermission() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CheckChatPermission() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestChatsServiceImpl_CreateChannel(t *testing.T) {
	type fields struct {
		UseCase useCase.ChatsUseCase
	}
	type args struct {
		ctx context.Context
		r   *grpc_utils.Channel
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *grpc_utils.ResponseMessage
		wantErr bool
	}{
		{
			name: "InternalErrorTest",
			fields: fields{
				UseCase: &useCase.ChatsUseCaseMock{
					CreateChannelFunc: func(channel *models.Channel) (u uint64, e error) {
						return 0, internalError
					},
				}},
			args: args{
				ctx: contextArg,
				r:   &grpc_utils.Channel{},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "SuccessTest",
			fields: fields{
				UseCase: &useCase.ChatsUseCaseMock{
					CreateChannelFunc: func(channel *models.Channel) (u uint64, e error) {
						return 0, nil
					},
				}},
			args: args{
				ctx: contextArg,
				r:   &grpc_utils.Channel{},
			},
			want: &grpc_utils.ResponseMessage{
				Number: 0,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := ChatsServiceImpl{
				UseCase: tt.fields.UseCase,
			}
			got, err := c.CreateChannel(tt.args.ctx, tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateChannel() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateChannel() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestChatsServiceImpl_CreateWorkspace(t *testing.T) {
	type fields struct {
		UseCase useCase.ChatsUseCase
	}
	type args struct {
		ctx context.Context
		r   *grpc_utils.Workspace
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *grpc_utils.ResponseMessage
		wantErr bool
	}{
		{
			name: "InternalErrorTest",
			fields: fields{
				UseCase: &useCase.ChatsUseCaseMock{
					CreateWorkspaceFunc: func(channel *models.Workspace) (u uint64, e error) {
						return 0, internalError
					},
				}},
			args: args{
				ctx: contextArg,
				r:   &grpc_utils.Workspace{},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "SuccessTest",
			fields: fields{
				UseCase: &useCase.ChatsUseCaseMock{
					CreateWorkspaceFunc: func(channel *models.Workspace) (u uint64, e error) {
						return 0, nil
					},
				}},
			args: args{
				ctx: contextArg,
				r:   &grpc_utils.Workspace{},
			},
			want: &grpc_utils.ResponseMessage{
				Number: 0,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := ChatsServiceImpl{
				UseCase: tt.fields.UseCase,
			}
			got, err := c.CreateWorkspace(tt.args.ctx, tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateWorkspace() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateWorkspace() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestChatsServiceImpl_DeleteChannel(t *testing.T) {
	type fields struct {
		UseCase useCase.ChatsUseCase
	}
	type args struct {
		ctx context.Context
		r   *grpc_utils.RequestMessage
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *grpc_utils.EmptyChats
		wantErr bool
	}{
		{
			name: "InternalErrorTest",
			fields: fields{
				UseCase: &useCase.ChatsUseCaseMock{
					DeleteChannelFunc: func(userID uint64, channelID uint64) error {
						return internalError
					},
				}},
			args: args{
				ctx: contextArg,
				r:   &grpc_utils.RequestMessage{},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "SuccessTest",
			fields: fields{
				UseCase: &useCase.ChatsUseCaseMock{
					DeleteChannelFunc: func(userID uint64, channelID uint64) error {
						return nil
					},
				}},
			args: args{
				ctx: contextArg,
				r:   &grpc_utils.RequestMessage{},
			},
			want:    &grpc_utils.EmptyChats{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := ChatsServiceImpl{
				UseCase: tt.fields.UseCase,
			}
			got, err := c.DeleteChannel(tt.args.ctx, tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeleteChannel() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteChannel() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestChatsServiceImpl_DeleteChat(t *testing.T) {
	type fields struct {
		UseCase useCase.ChatsUseCase
	}
	type args struct {
		ctx context.Context
		r   *grpc_utils.RequestMessage
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *grpc_utils.EmptyChats
		wantErr bool
	}{
		{
			name: "InternalErrorTest",
			fields: fields{
				UseCase: &useCase.ChatsUseCaseMock{
					DeleteChatFunc: func(userID uint64, channelID uint64) error {
						return internalError
					},
				}},
			args: args{
				ctx: contextArg,
				r:   &grpc_utils.RequestMessage{},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "SuccessTest",
			fields: fields{
				UseCase: &useCase.ChatsUseCaseMock{
					DeleteChatFunc: func(userID uint64, channelID uint64) error {
						return nil
					},
				}},
			args: args{
				ctx: contextArg,
				r:   &grpc_utils.RequestMessage{},
			},
			want:    &grpc_utils.EmptyChats{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := ChatsServiceImpl{
				UseCase: tt.fields.UseCase,
			}
			got, err := c.DeleteChat(tt.args.ctx, tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeleteChat() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteChat() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestChatsServiceImpl_DeleteWorkspace(t *testing.T) {
	type fields struct {
		UseCase useCase.ChatsUseCase
	}
	type args struct {
		ctx context.Context
		r   *grpc_utils.RequestMessage
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *grpc_utils.EmptyChats
		wantErr bool
	}{
		{
			name: "InternalErrorTest",
			fields: fields{
				UseCase: &useCase.ChatsUseCaseMock{
					DeleteWorkspaceFunc: func(userID uint64, channelID uint64) error {
						return internalError
					},
				}},
			args: args{
				ctx: contextArg,
				r:   &grpc_utils.RequestMessage{},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "SuccessTest",
			fields: fields{
				UseCase: &useCase.ChatsUseCaseMock{
					DeleteWorkspaceFunc: func(userID uint64, channelID uint64) error {
						return nil
					},
				}},
			args: args{
				ctx: contextArg,
				r:   &grpc_utils.RequestMessage{},
			},
			want:    &grpc_utils.EmptyChats{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := ChatsServiceImpl{
				UseCase: tt.fields.UseCase,
			}
			got, err := c.DeleteWorkspace(tt.args.ctx, tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeleteWorkspace() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteWorkspace() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestChatsServiceImpl_EditChannel(t *testing.T) {
	type fields struct {
		UseCase useCase.ChatsUseCase
	}
	type args struct {
		ctx context.Context
		r   *grpc_utils.RequestMessage
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *grpc_utils.EmptyChats
		wantErr bool
	}{
		{
			name: "InternalErrorTest",
			fields: fields{
				UseCase: &useCase.ChatsUseCaseMock{
					EditChannelFunc: func(userID uint64, channel *models.Channel) error {
						return internalError
					},
				}},
			args: args{
				ctx: contextArg,
				r: &grpc_utils.RequestMessage{
					Channel: encodeChannel(models.Channel{}),
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "SuccessTest",
			fields: fields{
				UseCase: &useCase.ChatsUseCaseMock{
					EditChannelFunc: func(userID uint64, channel *models.Channel) error {
						return nil
					},
				}},
			args: args{
				ctx: contextArg,
				r:   &grpc_utils.RequestMessage{Channel: encodeChannel(models.Channel{})},
			},
			want:    &grpc_utils.EmptyChats{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := ChatsServiceImpl{
				UseCase: tt.fields.UseCase,
			}
			got, err := c.EditChannel(tt.args.ctx, tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("EditChannel() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EditChannel() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestChatsServiceImpl_EditWorkspace(t *testing.T) {
	type fields struct {
		UseCase useCase.ChatsUseCase
	}
	type args struct {
		ctx context.Context
		r   *grpc_utils.RequestMessage
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *grpc_utils.EmptyChats
		wantErr bool
	}{
		{
			name: "InternalErrorTest",
			fields: fields{
				UseCase: &useCase.ChatsUseCaseMock{
					EditWorkspaceFunc: func(userID uint64, channel *models.Workspace) error {
						return internalError
					},
				}},
			args: args{
				ctx: contextArg,
				r: &grpc_utils.RequestMessage{
					Workspace: encodeWorkspace(models.Workspace{}),
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "SuccessTest",
			fields: fields{
				UseCase: &useCase.ChatsUseCaseMock{
					EditWorkspaceFunc: func(userID uint64, channel *models.Workspace) error {
						return nil
					},
				}},
			args: args{
				ctx: contextArg,
				r: &grpc_utils.RequestMessage{
					Workspace: encodeWorkspace(models.Workspace{}),
				},
			},
			want:    &grpc_utils.EmptyChats{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := ChatsServiceImpl{
				UseCase: tt.fields.UseCase,
			}
			got, err := c.EditWorkspace(tt.args.ctx, tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("EditWorkspace() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EditWorkspace() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestChatsServiceImpl_GetChannelByID(t *testing.T) {
	type fields struct {
		UseCase useCase.ChatsUseCase
	}
	type args struct {
		ctx context.Context
		r   *grpc_utils.RequestMessage
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *grpc_utils.ResponseMessage
		wantErr bool
	}{
		{
			name: "InternalErrorTest",
			fields: fields{
				UseCase: &useCase.ChatsUseCaseMock{
					GetChannelByIDFunc: func(userID uint64, ID uint64) (channel models.Channel, e error) {
						return models.Channel{}, internalError
					},
				}},
			args: args{
				ctx: contextArg,
				r:   &grpc_utils.RequestMessage{},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "SuccessTest",
			fields: fields{
				UseCase: &useCase.ChatsUseCaseMock{
					GetChannelByIDFunc: func(userID uint64, ID uint64) (channel models.Channel, e error) {
						return models.Channel{}, nil
					},
				}},
			args: args{
				ctx: contextArg,
				r:   &grpc_utils.RequestMessage{},
			},
			want: &grpc_utils.ResponseMessage{
				Channel: encodeChannel(models.Channel{}),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := ChatsServiceImpl{
				UseCase: tt.fields.UseCase,
			}
			got, err := c.GetChannelByID(tt.args.ctx, tt.args.r)
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

func TestChatsServiceImpl_GetChatByID(t *testing.T) {
	type fields struct {
		UseCase useCase.ChatsUseCase
	}
	type args struct {
		ctx context.Context
		r   *grpc_utils.RequestMessage
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *grpc_utils.ResponseMessage
		wantErr bool
	}{
		{
			name: "InternalErrorTest",
			fields: fields{
				UseCase: &useCase.ChatsUseCaseMock{
					GetChatByIDFunc: func(userID uint64, ID uint64) (chat models.Chat, e error) {
						return models.Chat{}, internalError
					},
				}},
			args: args{
				ctx: contextArg,
				r:   &grpc_utils.RequestMessage{},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "SuccessTest",
			fields: fields{
				UseCase: &useCase.ChatsUseCaseMock{
					GetChatByIDFunc: func(userID uint64, ID uint64) (chat models.Chat, e error) {
						return models.Chat{}, nil
					},
				}},
			args: args{
				ctx: contextArg,
				r:   &grpc_utils.RequestMessage{},
			},
			want: &grpc_utils.ResponseMessage{
				Chat: encodeChat(models.Chat{}),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := ChatsServiceImpl{
				UseCase: tt.fields.UseCase,
			}
			got, err := c.GetChatByID(tt.args.ctx, tt.args.r)
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

func TestChatsServiceImpl_GetChatsByUserID(t *testing.T) {
	type fields struct {
		UseCase useCase.ChatsUseCase
	}
	type args struct {
		ctx context.Context
		r   *grpc_utils.RequestMessage
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *grpc_utils.ResponseMessage
		wantErr bool
	}{
		{
			name: "InternalErrorTest",
			fields: fields{
				UseCase: &useCase.ChatsUseCaseMock{
					GetChatsByUserIDFunc: func(ID uint64) (chats []models.Chat, e error) {
						return nil, internalError
					},
				}},
			args: args{
				ctx: contextArg,
				r:   &grpc_utils.RequestMessage{},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "SuccessTest",
			fields: fields{
				UseCase: &useCase.ChatsUseCaseMock{
					GetChatsByUserIDFunc: func(ID uint64) (chats []models.Chat, e error) {
						return nil, nil
					},
				}},
			args: args{
				ctx: contextArg,
				r:   &grpc_utils.RequestMessage{},
			},
			want: &grpc_utils.ResponseMessage{
				Chats: make([]*grpc_utils.Chat, 0),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := ChatsServiceImpl{
				UseCase: tt.fields.UseCase,
			}
			got, err := c.GetChatsByUserID(tt.args.ctx, tt.args.r)
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

func TestChatsServiceImpl_GetWorkspaceByID(t *testing.T) {
	type fields struct {
		UseCase useCase.ChatsUseCase
	}
	type args struct {
		ctx context.Context
		r   *grpc_utils.RequestMessage
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *grpc_utils.ResponseMessage
		wantErr bool
	}{
		{
			name: "InternalErrorTest",
			fields: fields{
				UseCase: &useCase.ChatsUseCaseMock{
					GetWorkspaceByIDFunc: func(userID uint64, ID uint64) (workspace models.Workspace, e error) {
						return models.Workspace{}, internalError
					},
				}},
			args: args{
				ctx: contextArg,
				r:   &grpc_utils.RequestMessage{},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "SuccessTest",
			fields: fields{
				UseCase: &useCase.ChatsUseCaseMock{
					GetWorkspaceByIDFunc: func(userID uint64, ID uint64) (workspace models.Workspace, e error) {
						return models.Workspace{}, nil
					},
				}},
			args: args{
				ctx: contextArg,
				r:   &grpc_utils.RequestMessage{},
			},
			want: &grpc_utils.ResponseMessage{
				Workspace: encodeWorkspace(models.Workspace{}),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := ChatsServiceImpl{
				UseCase: tt.fields.UseCase,
			}
			got, err := c.GetWorkspaceByID(tt.args.ctx, tt.args.r)
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

func TestChatsServiceImpl_GetWorkspacesByUserID(t *testing.T) {
	type fields struct {
		UseCase useCase.ChatsUseCase
	}
	type args struct {
		ctx context.Context
		r   *grpc_utils.RequestMessage
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *grpc_utils.ResponseMessage
		wantErr bool
	}{
		{
			name: "InternalErrorTest",
			fields: fields{
				UseCase: &useCase.ChatsUseCaseMock{
					GetWorkspacesByUserIDFunc: func(ID uint64) (workspaces []models.Workspace, e error) {
						return nil, internalError
					},
				}},
			args: args{
				ctx: contextArg,
				r:   &grpc_utils.RequestMessage{},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "SuccessTest",
			fields: fields{
				UseCase: &useCase.ChatsUseCaseMock{
					GetWorkspacesByUserIDFunc: func(ID uint64) (workspaces []models.Workspace, e error) {
						return nil, nil
					},
				}},
			args: args{
				ctx: contextArg,
				r:   &grpc_utils.RequestMessage{},
			},
			want: &grpc_utils.ResponseMessage{
				Workspaces: make([]*grpc_utils.Workspace, 0),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := ChatsServiceImpl{
				UseCase: tt.fields.UseCase,
			}
			got, err := c.GetWorkspacesByUserID(tt.args.ctx, tt.args.r)
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

func TestChatsServiceImpl_LogoutFromChannel(t *testing.T) {
	type fields struct {
		UseCase useCase.ChatsUseCase
	}
	type args struct {
		ctx context.Context
		r   *grpc_utils.RequestMessage
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *grpc_utils.EmptyChats
		wantErr bool
	}{
		{
			name: "InternalErrorTest",
			fields: fields{
				UseCase: &useCase.ChatsUseCaseMock{
					LogoutFromChannelFunc: func(userID uint64, channelID uint64) error {
						return internalError
					},
				}},
			args: args{
				ctx: contextArg,
				r:   &grpc_utils.RequestMessage{},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "SuccessTest",
			fields: fields{
				UseCase: &useCase.ChatsUseCaseMock{
					LogoutFromChannelFunc: func(userID uint64, channelID uint64) error {
						return nil
					},
				}},
			args: args{
				ctx: contextArg,
				r:   &grpc_utils.RequestMessage{},
			},
			want:    &grpc_utils.EmptyChats{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := ChatsServiceImpl{
				UseCase: tt.fields.UseCase,
			}
			got, err := c.LogoutFromChannel(tt.args.ctx, tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("LogoutFromChannel() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LogoutFromChannel() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestChatsServiceImpl_LogoutFromWorkspace(t *testing.T) {
	type fields struct {
		UseCase useCase.ChatsUseCase
	}
	type args struct {
		ctx context.Context
		r   *grpc_utils.RequestMessage
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *grpc_utils.EmptyChats
		wantErr bool
	}{
		{
			name: "InternalErrorTest",
			fields: fields{
				UseCase: &useCase.ChatsUseCaseMock{
					LogoutFromWorkspaceFunc: func(userID uint64, channelID uint64) error {
						return internalError
					},
				}},
			args: args{
				ctx: contextArg,
				r:   &grpc_utils.RequestMessage{},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "SuccessTest",
			fields: fields{
				UseCase: &useCase.ChatsUseCaseMock{
					LogoutFromWorkspaceFunc: func(userID uint64, channelID uint64) error {
						return nil
					},
				}},
			args: args{
				ctx: contextArg,
				r:   &grpc_utils.RequestMessage{},
			},
			want:    &grpc_utils.EmptyChats{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := ChatsServiceImpl{
				UseCase: tt.fields.UseCase,
			}
			got, err := c.LogoutFromWorkspace(tt.args.ctx, tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("LogoutFromWorkspace() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LogoutFromWorkspace() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestChatsServiceImpl_PutChat(t *testing.T) {
	type fields struct {
		UseCase useCase.ChatsUseCase
	}
	type args struct {
		ctx  context.Context
		chat *grpc_utils.Chat
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *grpc_utils.ResponseMessage
		wantErr bool
	}{
		{
			name: "InternalErrorTest",
			fields: fields{
				UseCase: &useCase.ChatsUseCaseMock{
					PutChatFunc: func(chat *models.Chat) (u uint64, e error) {
						return 0, internalError
					},
				}},
			args: args{
				ctx:  contextArg,
				chat: &grpc_utils.Chat{},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "SuccessTest",
			fields: fields{
				UseCase: &useCase.ChatsUseCaseMock{
					PutChatFunc: func(chat *models.Chat) (u uint64, e error) {
						return 0, nil
					},
				}},
			args: args{
				ctx:  contextArg,
				chat: &grpc_utils.Chat{},
			},
			want: &grpc_utils.ResponseMessage{
				Number: 0,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := ChatsServiceImpl{
				UseCase: tt.fields.UseCase,
			}
			got, err := c.PutChat(tt.args.ctx, tt.args.chat)
			if (err != nil) != tt.wantErr {
				t.Errorf("PutChat() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PutChat() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewGRPCChatsService(t *testing.T) {
	type args struct {
		useCase useCase.ChatsUseCase
	}
	tests := []struct {
		name string
		args args
		want grpc_utils.ChatsServiceServer
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewGRPCChatsService(tt.args.useCase); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewGRPCChatsService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_encodeChannel(t *testing.T) {
	type args struct {
		channel models.Channel
	}
	tests := []struct {
		name string
		args args
		want *grpc_utils.Channel
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := encodeChannel(tt.args.channel); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("encodeChannel() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_encodeChat(t *testing.T) {
	type args struct {
		chat models.Chat
	}
	tests := []struct {
		name string
		args args
		want *grpc_utils.Chat
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := encodeChat(tt.args.chat); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("encodeChat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_encodeWorkspace(t *testing.T) {
	type args struct {
		workspace models.Workspace
	}
	tests := []struct {
		name string
		args args
		want *grpc_utils.Workspace
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := encodeWorkspace(tt.args.workspace); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("encodeWorkspace() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseChannel(t *testing.T) {
	type args struct {
		channel *grpc_utils.Channel
	}
	tests := []struct {
		name string
		args args
		want *models.Channel
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseChannel(tt.args.channel); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseChannel() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseChat(t *testing.T) {
	type args struct {
		chat *grpc_utils.Chat
	}
	tests := []struct {
		name string
		args args
		want *models.Chat
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseChat(tt.args.chat); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseChat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseWorkspace(t *testing.T) {
	type args struct {
		workspace *grpc_utils.Workspace
	}
	tests := []struct {
		name string
		args args
		want *models.Workspace
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseWorkspace(tt.args.workspace); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseWorkspace() = %v, want %v", got, tt.want)
			}
		})
	}
}

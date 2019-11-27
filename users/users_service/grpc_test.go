package users_service

import (
	"context"
	"errors"
	"github.com/CoolCodeTeam/2019_2_CoolCodeMicroServices/users/usecase"
	"github.com/CoolCodeTeam/2019_2_CoolCodeMicroServices/utils/grpc_utils"
	"github.com/CoolCodeTeam/2019_2_CoolCodeMicroServices/utils/models"
	"reflect"
	"testing"
)

var internalError error = errors.New("Internal error")
var contextArg context.Context = context.Background()

func TestNewGRPCUsersService(t *testing.T) {
	type args struct {
		useCase useCase.UsersUseCase
	}
	tests := []struct {
		name string
		args args
		want grpc_utils.UsersServiceServer
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewGRPCUsersService(tt.args.useCase); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewGRPCUsersService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUsersServiceImpl_ChangeUser(t *testing.T) {
	type fields struct {
		UseCase useCase.UsersUseCase
	}
	type args struct {
		ctx  context.Context
		user *grpc_utils.User
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *grpc_utils.Empty
		wantErr bool
	}{
		{
			name: "InternalError",
			fields: fields{
				UseCase: &useCase.UsersUseCaseMock{
					ChangeUserFunc: func(user *models.User) error {
						return internalError
					},
				},
			},
			args: args{
				ctx:  nil,
				user: &grpc_utils.User{},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Success",
			fields: fields{
				UseCase: &useCase.UsersUseCaseMock{
					ChangeUserFunc: func(user *models.User) error {
						return nil
					},
				},
			},
			args: args{
				ctx:  nil,
				user: &grpc_utils.User{},
			},
			want:    &grpc_utils.Empty{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UsersServiceImpl{
				UseCase: tt.fields.UseCase,
			}
			got, err := u.ChangeUser(tt.args.ctx, tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("ChangeUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ChangeUser() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUsersServiceImpl_FindUsers(t *testing.T) {
	type fields struct {
		UseCase useCase.UsersUseCase
	}
	type args struct {
		ctx  context.Context
		name *grpc_utils.UserName
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *grpc_utils.Users
		wantErr bool
	}{
		{
			name: "InternalError",
			fields: fields{
				UseCase: &useCase.UsersUseCaseMock{
					FindUsersFunc: func(name string) (users models.Users, e error) {
						return models.Users{}, internalError
					},
				},
			},
			args: args{
				ctx:  nil,
				name: &grpc_utils.UserName{},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Success",
			fields: fields{
				UseCase: &useCase.UsersUseCaseMock{
					FindUsersFunc: func(name string) (users models.Users, e error) {
						return models.Users{}, nil
					},
				},
			},
			args: args{
				ctx:  nil,
				name: &grpc_utils.UserName{},
			},
			want: &grpc_utils.Users{
				Users: make([]*grpc_utils.User, 0),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UsersServiceImpl{
				UseCase: tt.fields.UseCase,
			}
			got, err := u.FindUsers(tt.args.ctx, tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindUsers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindUsers() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUsersServiceImpl_GetUserByEmail(t *testing.T) {
	type fields struct {
		UseCase useCase.UsersUseCase
	}
	type args struct {
		ctx   context.Context
		email *grpc_utils.UserEmail
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *grpc_utils.User
		wantErr bool
	}{
		{
			name: "InternalError",
			fields: fields{
				UseCase: &useCase.UsersUseCaseMock{
					GetUserByEmailFunc: func(email string) (user models.User, e error) {
						return models.User{}, internalError
					},
				},
			},
			args: args{
				ctx:   nil,
				email: &grpc_utils.UserEmail{},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Success",
			fields: fields{
				UseCase: &useCase.UsersUseCaseMock{
					GetUserByEmailFunc: func(email string) (user models.User, e error) {
						return models.User{}, nil
					},
				},
			},
			args: args{
				ctx:   nil,
				email: &grpc_utils.UserEmail{},
			},
			want:    &grpc_utils.User{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UsersServiceImpl{
				UseCase: tt.fields.UseCase,
			}
			got, err := u.GetUserByEmail(tt.args.ctx, tt.args.email)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUserByEmail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUserByEmail() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUsersServiceImpl_GetUserByID(t *testing.T) {
	type fields struct {
		UseCase useCase.UsersUseCase
	}
	type args struct {
		ctx    context.Context
		userID *grpc_utils.UserID
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *grpc_utils.User
		wantErr bool
	}{
		{
			name: "InternalError",
			fields: fields{
				UseCase: &useCase.UsersUseCaseMock{
					GetUserByIDFunc: func(id uint64) (user models.User, e error) {
						return models.User{}, internalError
					},
				},
			},
			args: args{
				ctx:    nil,
				userID: &grpc_utils.UserID{},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Success",
			fields: fields{
				UseCase: &useCase.UsersUseCaseMock{
					GetUserByIDFunc: func(id uint64) (user models.User, e error) {
						return models.User{}, nil
					},
				},
			},
			args: args{
				ctx:    nil,
				userID: &grpc_utils.UserID{},
			},
			want:    &grpc_utils.User{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UsersServiceImpl{
				UseCase: tt.fields.UseCase,
			}
			got, err := u.GetUserByID(tt.args.ctx, tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUserByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUserByID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUsersServiceImpl_GetUserBySession(t *testing.T) {
	type fields struct {
		UseCase useCase.UsersUseCase
	}
	type args struct {
		ctx context.Context
		r   *grpc_utils.Session
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *grpc_utils.UserID
		wantErr bool
	}{
		{
			name: "InternalError",
			fields: fields{
				UseCase: &useCase.UsersUseCaseMock{
					GetUserBySessionFunc: func(session string) (u uint64, e error) {
						return 0, internalError
					},
				},
			},
			args: args{
				ctx: nil,
				r:   &grpc_utils.Session{},
			},
			want:    &grpc_utils.UserID{},
			wantErr: true,
		},
		{
			name: "Success",
			fields: fields{
				UseCase: &useCase.UsersUseCaseMock{
					GetUserBySessionFunc: func(session string) (u uint64, e error) {
						return 0, nil
					},
				},
			},
			args: args{
				ctx: nil,
				r:   &grpc_utils.Session{},
			},
			want:    &grpc_utils.UserID{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UsersServiceImpl{
				UseCase: tt.fields.UseCase,
			}
			got, err := u.GetUserBySession(tt.args.ctx, tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUserBySession() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUserBySession() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUsersServiceImpl_Login(t *testing.T) {
	type fields struct {
		UseCase useCase.UsersUseCase
	}
	type args struct {
		ctx  context.Context
		user *grpc_utils.User
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *grpc_utils.User
		wantErr bool
	}{
		{
			name: "InternalError",
			fields: fields{
				UseCase: &useCase.UsersUseCaseMock{
					LoginFunc: func(user models.User) (user2 models.User, e error) {
						return models.User{}, internalError
					},
				},
			},
			args: args{
				ctx:  nil,
				user: &grpc_utils.User{},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Success",
			fields: fields{
				UseCase: &useCase.UsersUseCaseMock{
					LoginFunc: func(user models.User) (user2 models.User, e error) {
						return models.User{}, nil
					},
				},
			},
			args: args{
				ctx:  nil,
				user: &grpc_utils.User{},
			},
			want:    &grpc_utils.User{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UsersServiceImpl{
				UseCase: tt.fields.UseCase,
			}
			got, err := u.Login(tt.args.ctx, tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("Login() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Login() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUsersServiceImpl_SignUp(t *testing.T) {
	type fields struct {
		UseCase useCase.UsersUseCase
	}
	type args struct {
		ctx  context.Context
		user *grpc_utils.User
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *grpc_utils.Empty
		wantErr bool
	}{
		{
			name: "InternalError",
			fields: fields{
				UseCase: &useCase.UsersUseCaseMock{
					SignUpFunc: func(user *models.User) error {
						return internalError
					},
				},
			},
			args: args{
				ctx:  nil,
				user: &grpc_utils.User{},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Success",
			fields: fields{
				UseCase: &useCase.UsersUseCaseMock{
					SignUpFunc: func(user *models.User) error {
						return nil
					},
				},
			},
			args: args{
				ctx:  nil,
				user: &grpc_utils.User{},
			},
			want:    &grpc_utils.Empty{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UsersServiceImpl{
				UseCase: tt.fields.UseCase,
			}
			got, err := u.SignUp(tt.args.ctx, tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("SignUp() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SignUp() got = %v, want %v", got, tt.want)
			}
		})
	}
}

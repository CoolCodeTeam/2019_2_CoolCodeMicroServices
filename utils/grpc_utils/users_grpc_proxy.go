package grpc_utils

import (
	"context"
	"errors"

	useCase "github.com/CoolCodeTeam/2019_2_CoolCodeMicroServices/users/usecase"
	"github.com/CoolCodeTeam/2019_2_CoolCodeMicroServices/utils/models"
)

type UsersGRPCProxy struct {
	client UsersServiceClient
}

func (u *UsersGRPCProxy) PutStickerpack(userID uint64, stickerpackID uint64) error {
	return errors.New("Not implemented")
}

func (u *UsersGRPCProxy) GetUserBySession(session string) (uint64, error) {
	id, err := u.client.GetUserBySession(context.Background(), &Session{
		Value: session,
	})
	if err != nil {
		return 0, err
	}
	return id.ID, err
}

func (u *UsersGRPCProxy) GetUserByID(id uint64) (models.User, error) {
	user, err := u.client.GetUserByID(context.Background(), &UserID{
		ID: id,
	})
	return models.User{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Name:     user.Name,
		Password: user.Password,
		Status:   user.Status,
		Phone:    user.Phone,
	}, err
}

func (u *UsersGRPCProxy) GetUserByEmail(email string) (models.User, error) {
	user, err := u.client.GetUserByEmail(context.Background(), &UserEmail{
		Email: email,
	})
	return models.User{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Name:     user.Name,
		Password: user.Password,
		Status:   user.Status,
		Phone:    user.Phone,
	}, err
}

func (u *UsersGRPCProxy) SignUp(user *models.User) error {
	_, err := u.client.SignUp(context.Background(), &User{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Name:     user.Name,
		Password: user.Password,
		Status:   user.Status,
		Phone:    user.Phone,
	})
	return err
}

func (u *UsersGRPCProxy) Login(user models.User) (models.User, error) {
	returnUser, err := u.client.Login(context.Background(), &User{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Name:     user.Name,
		Password: user.Password,
		Status:   user.Status,
		Phone:    user.Phone,
	})
	return models.User{
		ID:       returnUser.ID,
		Username: returnUser.Username,
		Email:    returnUser.Email,
		Name:     returnUser.Name,
		Password: returnUser.Password,
		Status:   returnUser.Status,
		Phone:    returnUser.Phone,
	}, err
}

func (u *UsersGRPCProxy) ChangeUser(user *models.User) error {
	_, err := u.client.Login(context.Background(), &User{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Name:     user.Name,
		Password: user.Password,
		Status:   user.Status,
		Phone:    user.Phone,
	})
	return err
}

func (u *UsersGRPCProxy) FindUsers(name string) (models.Users, error) {
	users, err := u.client.FindUsers(context.Background(), &UserName{
		Name: name,
	})
	result := make([]*models.User, len(users.GetUsers()))
	for _, user := range users.GetUsers() {
		result = append(result, &models.User{
			ID:       user.ID,
			Username: user.Username,
			Email:    user.Email,
			Name:     user.Name,
			Password: user.Password,
			Status:   user.Status,
			Phone:    user.Phone,
		})
	}
	return models.Users{Users: result}, err
}

func NewUsersGRPCProxy(client UsersServiceClient) useCase.UsersUseCase {
	return &UsersGRPCProxy{client: client}
}

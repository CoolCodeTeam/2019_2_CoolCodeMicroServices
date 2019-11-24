package users_service

import (
	"context"
	"github.com/go-park-mail-ru/2019_2_CoolCodeMicroServices/users/usecase"
	"github.com/go-park-mail-ru/2019_2_CoolCodeMicroServices/utils/grpc_utils"
	"github.com/go-park-mail-ru/2019_2_CoolCodeMicroServices/utils/models"
	"github.com/sirupsen/logrus"
)

type UsersService interface {
	GetUserByID(id uint64) (models.User, error)
	GetUserByEmail(email string) (models.User, error)
	SignUp(user *models.User) error
	Login(user models.User) (models.User, error)
	ChangeUser(user *models.User) error
	FindUsers(name string) (models.Users, error)
}

type UsersServiceImpl struct {
	UseCase useCase.UsersUseCase
}

func (u *UsersServiceImpl) GetUserBySession(ctx context.Context,r *grpc_utils.Session) (*grpc_utils.UserID, error) {
	id,err:=u.UseCase.GetUserBySession(r.Value)
	return &grpc_utils.UserID{
		ID: id,
	},err
}

func (u *UsersServiceImpl) GetUserByID(ctx context.Context, userID *grpc_utils.UserID) (*grpc_utils.User, error) {
	logger := logrus.WithFields(logrus.Fields{
		"method":  "grpc_GetUserByID",
		"user_id": userID.ID,
	})
	user, err := u.UseCase.GetUserByID(userID.ID)
	if err != nil {
		logger.Errorf("can not get user by id: %s", err)
		return nil, err
	}

	logger.Info("successful")
	return &grpc_utils.User{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Name:     user.Name,
		Password: user.Password,
		Status:   user.Status,
		Phone:    user.Phone,
	}, nil

}

func (u *UsersServiceImpl) GetUserByEmail(ctx context.Context, email *grpc_utils.UserEmail) (*grpc_utils.User, error) {
	logger := logrus.WithFields(logrus.Fields{
		"method":     "grpc_GetUserByEmail",
		"user_email": email.Email,
	})
	user, err := u.UseCase.GetUserByEmail(email.Email)
	if err != nil {
		logger.Errorf("can not get user by email: %s", err)
		return nil, err
	}

	logger.Info("successful")
	return &grpc_utils.User{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Name:     user.Name,
		Password: user.Password,
		Status:   user.Status,
		Phone:    user.Phone,
	}, nil
}

func (u *UsersServiceImpl) SignUp(ctx context.Context, user *grpc_utils.User) (*grpc_utils.Empty, error) {
	newUser := &models.User{
		Username: user.Username,
		Email:    user.Email,
	}
	logger := logrus.WithFields(logrus.Fields{
		"method":        "grpc_SignUp",
		"user_email":    user.Email,
		"user_username": user.Username,
	})
	err := u.UseCase.SignUp(newUser)
	if err != nil {
		logger.Errorf("can not sign up: %s", err)
		return nil, err
	}

	logger.Info("successful")
	return &grpc_utils.Empty{}, nil
}

func (u *UsersServiceImpl) Login(ctx context.Context, user *grpc_utils.User) (*grpc_utils.User, error) {
	newUser := models.User{
		Username: user.Username,
		Email:    user.Email,
	}
	logger := logrus.WithFields(logrus.Fields{
		"method":     "grpc_Login",
		"user_email": user.Email,
	})
	newUser, err := u.UseCase.Login(newUser)
	if err != nil {
		logger.Errorf("can not login: %s", err)
		return nil, err
	}

	logger.Info("successful")
	return &grpc_utils.User{
		ID:       newUser.ID,
		Username: newUser.Username,
		Email:    newUser.Email,
		Name:     newUser.Name,
		Password: newUser.Password,
		Status:   newUser.Status,
		Phone:    newUser.Phone,
	}, nil
}

func (u *UsersServiceImpl) ChangeUser(ctx context.Context, user *grpc_utils.User) (*grpc_utils.Empty, error) {
	edituser := &models.User{
		Username: user.Username,
		Email:    user.Email,
	}
	logger := logrus.WithFields(logrus.Fields{
		"method":     "grpc_ChangeUser",
		"user_email": user.Email,
	})
	err := u.UseCase.ChangeUser(edituser)
	if err != nil {
		logger.Errorf("can not change user: %s", err)
		return nil, err
	}

	logger.Info("successful")
	return &grpc_utils.Empty{}, nil
}

func (u *UsersServiceImpl) FindUsers(ctx context.Context, name *grpc_utils.UserName) (*grpc_utils.Users, error) {
	logger := logrus.WithFields(logrus.Fields{
		"method":   "grpc_FindUsers",
		"username": name.Name,
	})
	users, err := u.UseCase.FindUsers(name.Name)
	if err != nil {
		logger.Errorf("can not find resultUser: %s", err)
		return nil, err
	}

	logger.Info("successful")
	result := &grpc_utils.Users{
		Users: make([]*grpc_utils.User, 0, len(users.Users)),
	}

	for _, resultUser := range users.Users {
		result.Users = append(result.Users, &grpc_utils.User{
			ID:       resultUser.ID,
			Username: resultUser.Username,
			Email:    resultUser.Email,
			Name:     resultUser.Name,
			Password: resultUser.Password,
			Status:   resultUser.Status,
			Phone:    resultUser.Phone,
		})
	}

	return result, nil
}

func NewGRPCUsersService(useCase useCase.UsersUseCase) grpc_utils.UsersServiceServer {
	return &UsersServiceImpl{UseCase: useCase}
}

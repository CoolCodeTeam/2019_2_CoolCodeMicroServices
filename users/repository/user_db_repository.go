package repository

import (
	"database/sql"

	"github.com/CoolCodeTeam/2019_2_CoolCodeMicroServices/utils/models"
)

type DBUserStore struct {
	DB *sql.DB
}

func (D DBUserStore) GetUserByEmail(email string) (models.User, error) {
	panic("implement me")
}

func (D DBUserStore) GetUserByID(ID uint64) (models.User, error) {
	panic("implement me")
}

func (D DBUserStore) PutUser(newUser *models.User) (uint64, error) {
	panic("implement me")
}

func (D DBUserStore) Replace(ID uint64, newUser *models.User) error {
	panic("implement me")
}

func (D DBUserStore) Contains(user models.User) bool {
	panic("implement me")
}

func (D DBUserStore) GetUsers() (models.Users, error) {
	panic("implement me")
}

func (D DBUserStore) GetUserStickers(userID uint64) ([]uint64, error) {
	panic("implement me")
}

func (D DBUserStore) AddStickerpack(userID uint64, stickerpackID uint64) error {
	panic("implement me")
}

func NewUserDBStore(db *sql.DB) UserRepo {
	return &DBUserStore{
		db,
	}
}

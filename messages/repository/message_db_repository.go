package repository

import (
	"database/sql"

	"github.com/CoolCodeTeam/2019_2_CoolCodeMicroServices/utils/models"
)

type MessageDBRepository struct {
	DB *sql.DB
}

func (m *MessageDBRepository) PutMessage(message *models.Message) (uint64, error) {
	panic("implement me")
}

func (m *MessageDBRepository) GetMessageByID(messageID uint64) (*models.Message, error) {
	panic("implement me")
}

func (m *MessageDBRepository) GetMessagesByChatID(chatID uint64) (models.Messages, error) {
	panic("implement me")
}

func (m *MessageDBRepository) RemoveMessage(messageID uint64) error {
	panic("implement me")
}

func (m *MessageDBRepository) UpdateMessage(message *models.Message) error {
	panic("implement me")
}

func (m *MessageDBRepository) HideMessageForAuthor(userID uint64) error {
	panic("implement me")
}

func (m *MessageDBRepository) FindMessages(s string) (models.Messages, error) {
	panic("implement me")
}

func (m *MessageDBRepository) Like(messageID uint64) error {
	panic("implement me")
}

func NewMessageDbRepository(db *sql.DB) MessageRepository {
	return &MessageDBRepository{DB: db}
}

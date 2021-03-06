package repository

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/CoolCodeTeam/2019_2_CoolCodeMicroServices/utils/models"
)

type MessageDBRepository struct {
	DB *sql.DB
}

func (m *MessageDBRepository) Like(messageID uint64) error {
	_, err := m.DB.Exec("UPDATE messages SET likes = likes+1 where id=$1", messageID)
	if err != nil {
		return models.NewServerError(err, http.StatusInternalServerError, "Can not like message: "+err.Error())
	}
	return nil
}
func (m *MessageDBRepository) PutMessage(message *models.Message) (uint64, error) {
	var chatID uint64
	time, err := time.Parse("02.01.2006 15:04", message.MessageTime)
	if err != nil {
		return 0, models.NewClientError(err, http.StatusBadRequest, "Wrong date format")
	}
	row := m.DB.QueryRow(`INSERT into messages (type, body, fileid,fileExtension, chatid
		, authorid,messagetime,likes,stickerid) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9) returning id`,
		message.MessageType, message.Text, message.FileID, message.FileType, message.ChatID,
		message.AuthorID, time, message.Likes, message.StickerID)
	err = row.Scan(&chatID)

	if err != nil {
		return chatID, models.NewServerError(err, http.StatusInternalServerError, "Can not INSERT message in PutMessage "+err.Error())
	}
	return chatID, err

}

func (m *MessageDBRepository) GetMessageByID(messageID uint64) (*models.Message, error) {
	var returningMessage models.Message
	var messageTime time.Time
	row := m.DB.QueryRow("SELECT id,type,body,fileid,fileExtension,chatid,authorid,messagetime,likes,stickerid FROM messages where id=$1", messageID)
	if err := row.Scan(&returningMessage.ID, &returningMessage.MessageType, &returningMessage.Text,
		&returningMessage.FileID, &returningMessage.FileType,
		&returningMessage.ChatID, &returningMessage.AuthorID, &messageTime,
		&returningMessage.Likes, &returningMessage.StickerID); err != nil {
		return &returningMessage,
			models.NewServerError(err, http.StatusBadRequest, "Message not exists:(")
	}
	timeString := messageTime.Format("02.01.2006 15:04")
	returningMessage.MessageTime = timeString
	return &returningMessage, nil
}

func (m *MessageDBRepository) GetMessagesByChatID(chatID uint64) (models.Messages, error) {
	returningMessages := make([]*models.Message, 0)
	rows, err := m.DB.Query("SELECT id,type,body,fileid,fileextension,chatid,authorid,hideforauthor,messagetime,likes,stickerid FROM messages where chatid=$1 order by id asc ", chatID)
	if err != nil {
		return models.Messages{}, models.NewServerError(err, http.StatusInternalServerError,
			"Can not get messages in GetMessagesByChatId "+err.Error())
	}
	for rows.Next() {
		var message models.Message
		var messageTime time.Time
		err := rows.Scan(&message.ID, &message.MessageType, &message.Text, &message.FileID, &message.FileType,
			&message.ChatID, &message.AuthorID, &message.HideForAuthor, &messageTime, &message.Likes, &message.StickerID)

		timeString := messageTime.Format("02.01.2006 15:04")
		message.MessageTime = timeString
		if err != nil {
			return models.Messages{}, models.NewServerError(err, http.StatusInternalServerError,
				"Can not read message in GetMessagesByChatId "+err.Error())
		}
		returningMessages = append(returningMessages, &message)
	}
	return models.Messages{Messages: returningMessages}, nil
}

func (m *MessageDBRepository) RemoveMessage(messageID uint64) error {
	_, err := m.DB.Exec("DELETE from messages where id=$1", messageID)
	if err != nil {
		return models.NewServerError(err, http.StatusInternalServerError,
			"Can not delete message in RemoveMessage "+err.Error())
	}
	return nil
}

func (m *MessageDBRepository) UpdateMessage(message *models.Message) error {
	_, err := m.DB.Exec("UPDATE messages SET body=$1 WHERE id=$2", message.Text, message.ID)
	if err != nil {
		return models.NewServerError(err, http.StatusInternalServerError,
			"Can not update message in UpdateMessage "+err.Error())
	}
	return nil
}

func (m *MessageDBRepository) HideMessageForAuthor(messageID uint64) error {
	_, err := m.DB.Exec("UPDATE messages SET hideforauthor=$1 WHERE id=$2", true, messageID)
	if err != nil {
		return models.NewServerError(err, http.StatusInternalServerError,
			"Can not update message in HideMessageForAuthor "+err.Error())
	}
	return nil
}

func (m *MessageDBRepository) FindMessages(s string) (models.Messages, error) {
	returningMessages := make([]*models.Message, 0)
	rows, err := m.DB.Query("SELECT id,type,body,fileid,fileextension,chatid,authorid,hideforauthor,messagetime,likes FROM messages where position($1 in body)>0 ", s)
	if err != nil {
		return models.Messages{}, models.NewServerError(err, http.StatusInternalServerError,
			"Can not get messages in FindMessages "+err.Error())
	}
	for rows.Next() {
		var message models.Message
		var messageTime time.Time
		err := rows.Scan(&message.ID, &message.MessageType, &message.Text, &message.FileID, &message.FileType, &message.ChatID,
			&message.AuthorID, &message.HideForAuthor, &messageTime, &message.Likes)

		timeString := messageTime.Format("02.01.2006 15:04")
		message.MessageTime = timeString
		if err != nil {
			return models.Messages{}, models.NewServerError(err, http.StatusInternalServerError,
				"Can not scan messages in FindMessages "+err.Error())
		}
		returningMessages = append(returningMessages, &message)
	}
	return models.Messages{returningMessages}, nil

}

func NewMessageDbRepository(db *sql.DB) MessageRepository {
	return &MessageDBRepository{DB: db}
}

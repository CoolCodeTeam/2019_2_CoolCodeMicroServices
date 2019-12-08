package useCase

import (
	chats "github.com/CoolCodeTeam/2019_2_CoolCodeMicroServices/chats/usecase"
	"github.com/CoolCodeTeam/2019_2_CoolCodeMicroServices/messages/repository"
	"github.com/CoolCodeTeam/2019_2_CoolCodeMicroServices/utils/models"
	"net/http"
	"os"
)

//go:generate moq -out messages_ucase_mock.go . MessagesUseCase
type MessagesUseCase interface {
	SaveChatMessage(message *models.Message) (uint64, error)
	EditMessage(message *models.Message, userID uint64) error
	DeleteMessage(messageID uint64, userID uint64) error
	GetChatMessages(chatID uint64, userID uint64) (models.Messages, error)
	GetMessageByID(messageID uint64) (*models.Message, error)
	HideMessageForAuthor(messageID uint64, userID uint64) error
	SaveChannelMessage(message *models.Message) (uint64, error)
	GetChannelMessages(channelID uint64, userID uint64) (models.Messages, error)
	FindMessages(findString string, ID uint64) (models.Messages, error)
	Like(ID uint64) error
	SaveFile(userID, chatID uint64, file models.File) (string, error)
	GetFile(userID, chatID uint64, photoUID string) (*os.File, error)
}

type MessageUseCaseImpl struct {
	messageRepository repository.MessageRepository
	photos            repository.FileRepository
	chats             chats.ChatsUseCase
}

func NewMessageUseCase(messageRepository repository.MessageRepository, chats chats.ChatsUseCase) MessagesUseCase {
	return &MessageUseCaseImpl{
		messageRepository: messageRepository,
		chats:             chats,
		photos:            repository.NewFilesArrayRepository("files"),
	}
}

func (m *MessageUseCaseImpl) Like(ID uint64) error {
	return m.messageRepository.Like(ID)
}

func (m *MessageUseCaseImpl) GetChatMessages(chatID uint64, userID uint64) (models.Messages, error) {
	permissionOk, err := m.chats.CheckChatPermission(userID, chatID)
	if err != nil {
		return models.Messages{}, err
	}
	if !permissionOk {
		return models.Messages{}, models.NewClientError(nil, http.StatusForbidden, "Not enough permissions for this request:(")
	}

	return m.messageRepository.GetMessagesByChatID(chatID)
}

func (m *MessageUseCaseImpl) GetChannelMessages(chatID uint64, userID uint64) (models.Messages, error) {
	permissionOk, err := m.chats.CheckChannelPermission(userID, chatID)
	if err != nil {
		return models.Messages{}, err
	}
	if !permissionOk {
		return models.Messages{}, models.NewClientError(nil, http.StatusForbidden, "Not enough permissions for this request:(")
	}

	return m.messageRepository.GetMessagesByChatID(chatID)
}

func (m *MessageUseCaseImpl) GetMessageByID(messageID uint64) (*models.Message, error) {
	return m.messageRepository.GetMessageByID(messageID)
}

func (m *MessageUseCaseImpl) SaveFile(userID, chatID uint64, fileInfo models.File) (string, error) {
	permissionOk, err := m.chats.CheckChatPermission(userID, chatID)
	if err != nil {
		return "", err
	}
	if !permissionOk {
		return "", models.NewClientError(nil, http.StatusForbidden, "Not enough permissions for this request:(")
	}

	return m.photos.SaveFile(chatID, fileInfo)
}

func (m *MessageUseCaseImpl) GetFile(userID, chatID uint64, photoUID string) (*os.File, error) {
	permissionOk, err := m.chats.CheckChatPermission(userID, chatID)
	if err != nil {
		return &os.File{}, err
	}
	if !permissionOk {
		return &os.File{}, models.NewClientError(nil, http.StatusForbidden, "Not enough permissions for this request:(")
	}
	return m.photos.GetFile(chatID, photoUID)
}

func (m *MessageUseCaseImpl) SaveChatMessage(message *models.Message) (uint64, error) {
	permissionOk, err := m.chats.CheckChatPermission(message.AuthorID, message.ChatID)
	if err != nil {
		return 0, err
	}
	if !permissionOk {
		return 0, models.NewClientError(nil, http.StatusForbidden, "Not enough permissions for this request:(")
	}
	return m.messageRepository.PutMessage(message)
}

func (m *MessageUseCaseImpl) SaveChannelMessage(message *models.Message) (uint64, error) {
	permissionOk, err := m.chats.CheckChannelPermission(message.AuthorID, message.ChatID)
	if err != nil {
		return 0, err
	}
	if !permissionOk {
		return 0, models.NewClientError(nil, http.StatusForbidden, "Not enough permissions for this request:(")
	}
	return m.messageRepository.PutMessage(message)
}

func (m *MessageUseCaseImpl) EditMessage(message *models.Message, userID uint64) error {
	DBmessage, err := m.messageRepository.GetMessageByID(message.ID)
	if err != nil {
		return err
	}
	if userID != DBmessage.AuthorID {
		return models.NewClientError(nil, http.StatusForbidden, "Not enough permissions for this request:(")
	}
	return m.messageRepository.UpdateMessage(message)
}

func (m *MessageUseCaseImpl) DeleteMessage(messageID uint64, userID uint64) error {
	message, err := m.messageRepository.GetMessageByID(messageID)
	if err != nil {
		return err
	}
	if userID != message.AuthorID {
		return models.NewClientError(nil, http.StatusForbidden, "Not enough permissions for this request:(")
	}
	return m.messageRepository.RemoveMessage(messageID)
}

func (m *MessageUseCaseImpl) HideMessageForAuthor(messageID uint64, userID uint64) error {
	message, err := m.messageRepository.GetMessageByID(messageID)
	if err != nil {
		return err
	}
	if userID != message.AuthorID {
		return models.NewClientError(nil, http.StatusForbidden, "Not enough permissions for this request:(")
	}
	return m.messageRepository.HideMessageForAuthor(messageID)
}

func (m *MessageUseCaseImpl) FindMessages(findString string, ID uint64) (models.Messages, error) {
	messages, err := m.messageRepository.FindMessages(findString)
	if err != nil {
		return messages, err
	}
	result := models.Messages{}

	for _, message := range messages.Messages {
		ok, err := m.chats.CheckChatPermission(ID, message.ChatID)
		if err != nil {
			ok, _ := m.chats.CheckChannelPermission(ID, message.ChatID)
			if ok {
				result.Messages = append(result.Messages, message)
			}
			continue
		}
		if ok {
			result.Messages = append(result.Messages, message)
		}
	}
	return result, nil
}

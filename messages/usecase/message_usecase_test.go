package useCase

import (
	"errors"
	chats "github.com/CoolCodeTeam/2019_2_CoolCodeMicroServices/chats/usecase"
	"github.com/CoolCodeTeam/2019_2_CoolCodeMicroServices/messages/repository"
	"github.com/CoolCodeTeam/2019_2_CoolCodeMicroServices/utils/models"

	"github.com/stretchr/testify/assert"
	"testing"
)

var messageUseCase = MessageUseCaseImpl{
	messageRepository: &repository.MessageRepositoryMock{},
	chats:             &chats.ChatsUseCaseMock{},
}

func TestMessageUseCaseImpl_GetChatMessages(t *testing.T) {
	chatID := 0
	messageID := 0

	//test internal error
	messageUseCase.chats = &chats.ChatsUseCaseMock{
		CheckChatPermissionFunc: func(userID uint64, chatID uint64) (b bool, e error) {
			return false, errors.New("Internal error")
		},
	}

	_, err := messageUseCase.GetChatMessages(uint64(chatID), uint64(messageID))

	assert.NotNil(t, err)

	//test not permission
	messageUseCase.chats = &chats.ChatsUseCaseMock{
		CheckChatPermissionFunc: func(userID uint64, chatID uint64) (b bool, e error) {
			return false, nil
		},
	}

	_, err = messageUseCase.GetChatMessages(uint64(chatID), uint64(messageID))

	assert.NotNil(t, err)

	//test success
	messageUseCase.chats = &chats.ChatsUseCaseMock{
		CheckChatPermissionFunc: func(userID uint64, chatID uint64) (b bool, e error) {
			return true, nil
		},
	}
	messageUseCase.messageRepository = &repository.MessageRepositoryMock{
		GetMessagesByChatIDFunc: func(chatID uint64) (messages models.Messages, e error) {
			return models.Messages{}, nil
		},
	}

	_, err = messageUseCase.GetChatMessages(uint64(chatID), uint64(messageID))

	assert.Nil(t, err)
}

func TestMessageUseCaseImpl_GetChannelMessages(t *testing.T) {
	chatID := 0
	messageID := 0

	//test internal error
	messageUseCase.chats = &chats.ChatsUseCaseMock{
		CheckChannelPermissionFunc: func(userID uint64, chatID uint64) (b bool, e error) {
			return false, errors.New("Internal error")
		},
	}

	_, err := messageUseCase.GetChannelMessages(uint64(chatID), uint64(messageID))

	assert.NotNil(t, err)

	//test not permission
	messageUseCase.chats = &chats.ChatsUseCaseMock{
		CheckChannelPermissionFunc: func(userID uint64, chatID uint64) (b bool, e error) {
			return false, nil
		},
	}

	_, err = messageUseCase.GetChannelMessages(uint64(chatID), uint64(messageID))

	assert.NotNil(t, err)

	//test success
	messageUseCase.chats = &chats.ChatsUseCaseMock{
		CheckChannelPermissionFunc: func(userID uint64, chatID uint64) (b bool, e error) {
			return true, nil
		},
	}
	messageUseCase.messageRepository = &repository.MessageRepositoryMock{
		GetMessagesByChatIDFunc: func(chatID uint64) (messages models.Messages, e error) {
			return models.Messages{}, nil
		},
	}

	_, err = messageUseCase.GetChannelMessages(uint64(chatID), uint64(messageID))

	assert.Nil(t, err)
}

func TestMessageUseCaseImpl_SaveChatMessage(t *testing.T) {
	testMessage := &models.Message{
		ID:     0,
		ChatID: 0,
	}

	//test internal error
	messageUseCase.chats = &chats.ChatsUseCaseMock{
		CheckChatPermissionFunc: func(userID uint64, chatID uint64) (b bool, e error) {
			return false, errors.New("Internal error")
		},
	}

	_, err := messageUseCase.SaveChatMessage(testMessage)

	assert.NotNil(t, err)

	//test not permission
	messageUseCase.chats = &chats.ChatsUseCaseMock{
		CheckChatPermissionFunc: func(userID uint64, chatID uint64) (b bool, e error) {
			return false, nil
		},
	}

	_, err = messageUseCase.SaveChatMessage(testMessage)

	assert.NotNil(t, err)

	//test success
	messageUseCase.chats = &chats.ChatsUseCaseMock{
		CheckChatPermissionFunc: func(userID uint64, chatID uint64) (b bool, e error) {
			return true, nil
		},
	}
	messageUseCase.messageRepository = &repository.MessageRepositoryMock{
		PutMessageFunc: func(message *models.Message) (u uint64, e error) {
			return 0, nil
		},
	}

	_, err = messageUseCase.SaveChatMessage(testMessage)

	assert.Nil(t, err)
}

func TestMessageUseCaseImpl_SaveChannelMessage(t *testing.T) {
	testMessage := &models.Message{
		ID:     0,
		ChatID: 0,
	}

	//test internal error
	messageUseCase.chats = &chats.ChatsUseCaseMock{
		CheckChannelPermissionFunc: func(userID uint64, chatID uint64) (b bool, e error) {
			return false, errors.New("Internal error")
		},
	}

	_, err := messageUseCase.SaveChannelMessage(testMessage)

	assert.NotNil(t, err)

	//test not permission
	messageUseCase.chats = &chats.ChatsUseCaseMock{
		CheckChannelPermissionFunc: func(userID uint64, chatID uint64) (b bool, e error) {
			return false, nil
		},
	}

	_, err = messageUseCase.SaveChannelMessage(testMessage)

	assert.NotNil(t, err)

	//test success
	messageUseCase.chats = &chats.ChatsUseCaseMock{
		CheckChannelPermissionFunc: func(userID uint64, chatID uint64) (b bool, e error) {
			return true, nil
		},
	}
	messageUseCase.messageRepository = &repository.MessageRepositoryMock{
		PutMessageFunc: func(message *models.Message) (u uint64, e error) {
			return 0, nil
		},
	}

	_, err = messageUseCase.SaveChannelMessage(testMessage)

	assert.Nil(t, err)
}

func TestMessageUseCaseImpl_EditMessage(t *testing.T) {
	userID := 0
	testMessage := &models.Message{
		ID:     0,
		ChatID: 0,
	}

	//test internal error
	messageUseCase.messageRepository = &repository.MessageRepositoryMock{
		GetMessageByIDFunc: func(messageID uint64) (message *models.Message, e error) {
			return &models.Message{}, errors.New("Internal error")
		},
	}

	err := messageUseCase.EditMessage(testMessage, uint64(userID))

	assert.NotNil(t, err)

	//test not permission
	messageUseCase.messageRepository = &repository.MessageRepositoryMock{
		GetMessageByIDFunc: func(messageID uint64) (message *models.Message, e error) {
			return &models.Message{AuthorID: 1}, errors.New("Internal error")
		},
	}

	err = messageUseCase.EditMessage(testMessage, uint64(userID))

	assert.NotNil(t, err)

}

func TestMessageUseCaseImpl_DeleteMessage(t *testing.T) {
	authorID := 0
	messageID := 0

	//test internal error
	messageUseCase.messageRepository = &repository.MessageRepositoryMock{
		GetMessageByIDFunc: func(messageID uint64) (message *models.Message, e error) {
			return &models.Message{}, errors.New("Internal error")
		},
	}

	err := messageUseCase.DeleteMessage(uint64(messageID), uint64(authorID))

	assert.NotNil(t, err)

	//test not permission
	messageUseCase.messageRepository = &repository.MessageRepositoryMock{
		GetMessageByIDFunc: func(messageID uint64) (message *models.Message, e error) {
			return &models.Message{AuthorID: 1}, nil
		},
	}

	err = messageUseCase.DeleteMessage(uint64(messageID), uint64(authorID))

	assert.NotNil(t, err)

	//test success
	authorID = 1
	messageUseCase.messageRepository = &repository.MessageRepositoryMock{
		GetMessageByIDFunc: func(messageID uint64) (message *models.Message, e error) {
			return &models.Message{AuthorID: 1}, nil
		},
		RemoveMessageFunc: func(messageID uint64) error {
			return nil
		},
	}

	err = messageUseCase.DeleteMessage(uint64(messageID), uint64(authorID))

	assert.Nil(t, err)
}

func TestMessageUseCaseImpl_HideMessageForAuthor(t *testing.T) {
	authorID := 0
	messageID := 0

	//test internal error
	messageUseCase.messageRepository = &repository.MessageRepositoryMock{
		GetMessageByIDFunc: func(messageID uint64) (message *models.Message, e error) {
			return &models.Message{}, errors.New("Internal error")
		},
	}

	err := messageUseCase.HideMessageForAuthor(uint64(messageID), uint64(authorID))

	assert.NotNil(t, err)

	//test not permission
	messageUseCase.messageRepository = &repository.MessageRepositoryMock{
		GetMessageByIDFunc: func(messageID uint64) (message *models.Message, e error) {
			return &models.Message{AuthorID: 1}, nil
		},
	}

	err = messageUseCase.HideMessageForAuthor(uint64(messageID), uint64(authorID))

	assert.NotNil(t, err)

	//test success
	authorID = 1
	messageUseCase.messageRepository = &repository.MessageRepositoryMock{
		GetMessageByIDFunc: func(messageID uint64) (message *models.Message, e error) {
			return &models.Message{AuthorID: 1}, nil
		},
		HideMessageForAuthorFunc: func(userID uint64) error {
			return nil
		},
	}

	err = messageUseCase.HideMessageForAuthor(uint64(messageID), uint64(authorID))

	assert.Nil(t, err)
}

func TestMessageUseCaseImpl_FindMessages(t *testing.T) {
	authorID := 0
	//test internal error
	messageUseCase.messageRepository = &repository.MessageRepositoryMock{
		FindMessagesFunc: func(s string) (messages models.Messages, e error) {
			return models.Messages{}, errors.New("internal error")
		},
	}

	_, err := messageUseCase.FindMessages("test", uint64(authorID))

	assert.NotNil(t, err)

	//test chat not permission, channel ok
	messagesSlice := make([]*models.Message, 0)
	messagesSlice = append(messagesSlice, &models.Message{ChatID: 0})
	messageUseCase.messageRepository = &repository.MessageRepositoryMock{
		FindMessagesFunc: func(s string) (messages models.Messages, e error) {
			return models.Messages{
				Messages: messagesSlice,
			}, nil
		},
	}
	messageUseCase.chats = &chats.ChatsUseCaseMock{
		CheckChatPermissionFunc: func(userID uint64, chatID uint64) (b bool, e error) {
			return false, errors.New("not perm")
		},
		CheckChannelPermissionFunc: func(authorID uint64, channelID uint64) (b bool, e error) {
			return true, nil
		},
	}

	_, err = messageUseCase.FindMessages("test", uint64(authorID))

	assert.Nil(t, err)

}

func TestMessageUseCaseImpl_Like(t *testing.T) {
	messageID := 0
	messageUseCase.messageRepository = &repository.MessageRepositoryMock{
		LikeFunc: func(messageID uint64) error {
			return nil
		},
	}
	err := messageUseCase.Like(uint64(messageID))
	assert.Nil(t, err)
}

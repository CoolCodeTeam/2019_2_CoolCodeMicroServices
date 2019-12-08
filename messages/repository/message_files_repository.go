package repository

import (
	"github.com/CoolCodeTeam/2019_2_CoolCodeMicroServices/utils/models"
	"os"
)

//go:generate moq -out file_repository_mock.go . FileRepository
type FileRepository interface {
	SaveFile(chatID uint64, file models.File) (returnUID string, returnErr error)
	GetFile(chatID uint64, UID string) (*os.File, error)
}

package repository

import (
	"os"

	"github.com/CoolCodeTeam/2019_2_CoolCodeMicroServices/utils/models"
)

type FilesArrayRepository struct {
	dirPath string
}

func (f FilesArrayRepository) SaveFile(chatID uint64, file models.File) (returnUID string, returnErr error) {
	panic("implement me")
}

func (f FilesArrayRepository) GetFile(chatID uint64, UID string) (*os.File, error) {
	panic("implement me")
}

func NewFilesArrayRepository(path string) FileRepository {
	return &FilesArrayRepository{dirPath: path}
}

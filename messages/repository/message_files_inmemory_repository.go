package repository

import (
	"fmt"
	"github.com/CoolCodeTeam/2019_2_CoolCodeMicroServices/utils/models"
	"github.com/google/uuid"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

type FilesArrayRepository struct {
	dirPath string
}

func (p *FilesArrayRepository) SaveFile(chatID uint64, file models.File) (returnUID string, returnErr error) {
	path := p.dirPath + "/" + strconv.Itoa(int(chatID))
	defer func() {
		err := file.File.Close()
		if err != nil && returnErr == nil {
			returnErr = err
		}
	}()

	if _, err := os.Stat(p.dirPath); os.IsNotExist(err) {
		err = os.Mkdir(p.dirPath, os.ModePerm)
		if err != nil {
			return "", models.NewServerError(err, http.StatusInternalServerError, "Can not create dir in SaveMessageFile")
		}
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		err = os.Mkdir(path, os.ModePerm)
		if err != nil {
			return "", models.NewServerError(err, http.StatusInternalServerError, "Can not create dir in SaveMessageFile")
		}
	}
	tempFile, err := ioutil.TempFile(path, "upload-*"+file.Extension)
	if err != nil {
		return "", err
	}

	defer func() {
		err := tempFile.Close()

		if err != nil && returnErr == nil {
			returnErr = err
		}
	}()

	fileBytes, err := ioutil.ReadAll(file.File)
	if err != nil {
		return "", err
	}
	token := uuid.New()
	err = os.Rename(tempFile.Name(), path+"/"+token.String()+file.Extension)

	if err != nil {
		return token.String(), err
	}

	_, err = tempFile.Write(fileBytes)
	if err != nil {
		return token.String(), err
	}

	return token.String(), nil
}

func (p *FilesArrayRepository) GetFile(chatID uint64, UID string) (*os.File, error) {
	path := p.dirPath + "/" + strconv.Itoa(int(chatID))
	fileNames, err := filepath.Glob(path + "/" + UID + "*")
	if err != nil || len(fileNames) > 1 || len(fileNames) == 0 {
		return &os.File{}, models.NewServerError(err, http.StatusInternalServerError,
			fmt.Sprintf("Can not find file with uuid: %s", UID))
	}

	file, err := os.Open(fileNames[0])
	if err != nil {
		return file, err
	}
	return file, nil
}

func NewFilesArrayRepository(path string) FileRepository {
	return &FilesArrayRepository{dirPath: path}
}

package repository

import (
	"github.com/google/uuid"
	"io/ioutil"
	"mime/multipart"
	"os"
	"strconv"
)

type PhotosArrayRepository struct {
	dirPath string
}

func (p *PhotosArrayRepository) SavePhoto(chatID uint64, file multipart.File) (returnUID string, returnErr error) {
	defer func() {
		err := file.Close()

		if err != nil && returnErr == nil {
			returnErr = err
		}
	}()

	if _, err := os.Stat(p.dirPath + "/" + strconv.Itoa(int(chatID))); os.IsNotExist(err) {
		err = os.Mkdir(p.dirPath+"/"+strconv.Itoa(int(chatID)), os.ModePerm)
	}
	tempFile, err := ioutil.TempFile(p.dirPath+"/"+strconv.Itoa(int(chatID)), "upload-*.png")
	if err != nil {
		return "", err
	}

	defer func() {
		err := tempFile.Close()

		if err != nil && returnErr == nil {
			returnErr = err
		}
	}()

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		return "", err
	}
	token := uuid.New()
	err = os.Rename(tempFile.Name(), p.dirPath+"/"+strconv.Itoa(int(chatID))+"/"+token.String()+".png")

	if err != nil {
		return token.String(), err
	}

	_, err = tempFile.Write(fileBytes)
	if err != nil {
		return token.String(), err
	}

	return token.String(), nil
}

func (p *PhotosArrayRepository) GetPhoto(chatID uint64, UID string) (*os.File, error) {

	file, err := os.Open(p.dirPath + "/" + strconv.Itoa(int(chatID)) + "/" + UID + ".png")
	if err != nil {
		return file, err
	}
	return file, nil
}

func NewPhotosArrayRepository(path string) PhotoRepository {
	return &PhotosArrayRepository{dirPath: path}
}

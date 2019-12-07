package repository

import (
	"mime/multipart"
	"os"
)

//go:generate moq -out photo_repository_mock.go . PhotoRepository
type PhotoRepository interface {
	SavePhoto(chatID uint64,file multipart.File) (returnUID string,returnErr error)
	GetPhoto(chatID uint64,UID string) (*os.File, error)
}


package models

import "mime/multipart"

type File struct {
	File      multipart.File
	Extension string
}

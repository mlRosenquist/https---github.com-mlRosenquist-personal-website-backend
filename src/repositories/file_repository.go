package repositories

import (
	"bytes"
	"io"
	"personal-website-backend/models"
)

type FileRepository interface {
	StoreFile(fileName string, filetContent io.Reader) (*models.Image, error)
	GetAllFiles() ([]models.Image, error)
	GetFileContent(fileName string) (bytes.Buffer, error)
	DeleteFile(fileName string) error
}

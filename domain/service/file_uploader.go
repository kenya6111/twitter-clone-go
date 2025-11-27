package service

import (
	"io"
)

type FileInput struct {
	Filename string
	Size     int64
	Content  io.Reader
}

type FileUploader interface {
	UploadFile(file []FileInput) ([]string, error)
}

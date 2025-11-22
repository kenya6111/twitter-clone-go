package service

import "mime/multipart"

type FileUploader interface {
	UploadFile(file *multipart.Form) (string, error)
}

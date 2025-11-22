package file

import (
	"io"
	"log"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"

	"github.com/google/uuid"
)

type localFileUploader struct {
}

func NewLocalFileUploader() *localFileUploader {
	return &localFileUploader{}
}

func (lr *localFileUploader) UploadFile(form *multipart.Form) (string, error) {
	files := form.File["files"]
	var savepath string = ""

	for _, file := range files {
		uuid, err := uuid.NewRandom()
		if err != nil {
			log.Println(err)
			return "", err
		}

		dirName := "images/" + time.Now().Format("2006/01/02")
		err = os.MkdirAll(dirName, 0755)
		if err != nil {
			log.Println(err)
			return "", err
		}

		newFileName := uuid.String() + file.Filename
		savepath = filepath.Join(dirName, newFileName)

		src, err := file.Open()
		if err != nil {
			log.Println(err)
			return "", err
		}
		defer src.Close()

		var mode os.FileMode = 0o750

		dir := filepath.Dir(savepath)
		if err = os.MkdirAll(dir, mode); err != nil {
			log.Println(err)
			return "", err
		}
		if err = os.Chmod(dir, mode); err != nil {
			log.Println(err)
			return "", err
		}

		out, err := os.Create(savepath)
		if err != nil {
			log.Println(err)
			return "", err
		}
		defer out.Close()

		_, err = io.Copy(out, src)
		if err != nil {
			log.Println(err)
			return "", err
		}
	}

	return savepath, nil
}

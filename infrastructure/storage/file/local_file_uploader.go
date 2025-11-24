package file

import (
	"io"
	"log"
	"os"
	"path/filepath"
	"time"
	"twitter-clone-go/domain/service"

	"github.com/google/uuid"
)

type localFileUploader struct {
}

func NewLocalFileUploader() *localFileUploader {
	return &localFileUploader{}
}

func (lr *localFileUploader) UploadFile(files []service.FileInput) ([]string, error) {
	var savepathList []string

	for _, file := range files {
		uuid, err := uuid.NewRandom()
		if err != nil {
			log.Println(err)
			return nil, err
		}

		dirName := "images/" + time.Now().Format("2006/01/02")
		err = os.MkdirAll(dirName, 0755)
		if err != nil {
			log.Println(err)
			return nil, err
		}

		newFileName := uuid.String() + file.Filename
		savepath := filepath.Join(dirName, newFileName)
		savepathList = append(savepathList, savepath)

		src := file.Content

		var mode os.FileMode = 0o750

		dir := filepath.Dir(savepath)
		if err = os.MkdirAll(dir, mode); err != nil {
			log.Println(err)
			return nil, err
		}
		if err = os.Chmod(dir, mode); err != nil {
			log.Println(err)
			return nil, err
		}

		out, err := os.Create(savepath)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		defer out.Close()

		_, err = io.Copy(out, src)
		if err != nil {
			log.Println(err)
			return nil, err
		}
	}

	return savepathList, nil
}

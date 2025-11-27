package http

import (
	"io"
	"twitter-clone-go/application"
	"twitter-clone-go/domain/service"

	"github.com/gin-gonic/gin"
)

type TweetHandler struct {
	usecase application.TweetUsecase
}

func NewTweetHandler(u application.TweetUsecase) *TweetHandler {
	return &TweetHandler{usecase: u}
}

func (h *TweetHandler) CreateTweet(c *gin.Context) {
	var request application.TweetInfo

	form, err := c.MultipartForm()
	if err != nil {
		ErrorHandler(c, err)
		return
	}
	files := form.File["files"]
	var fileInputList []service.FileInput
	for _, fileHeader := range files {
		var file io.Reader
		if fileHeader != nil {
			file, err = fileHeader.Open()
			if err != nil {
				ErrorHandler(c, err)
				return
			}
		}
		fileInput := &service.FileInput{
			Filename: fileHeader.Filename,
			Size:     fileHeader.Size,
			Content:  file,
		}
		fileInputList = append(fileInputList, *fileInput)
	}

	request.UserId = c.PostForm("userId")
	request.Content = c.PostForm("content")
	request.ReplyToId = c.PostForm("replyToId")
	request.ImgFile = fileInputList

	tweet, err := h.usecase.CreateTweet(c.Request.Context(), request)
	if err != nil {
		ErrorHandler(c, err)
		return
	}

	SuccessResponse(c, ToCreateTweetResponse(tweet))
}

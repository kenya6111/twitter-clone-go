package http

import (
	"fmt"
	"twitter-clone-go/application"

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
	fmt.Println("^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^")

	request.UserId = c.PostForm("userId")
	request.Content = c.PostForm("content")
	request.ReplyToId = c.PostForm("replyToId")

	fmt.Println(request)
	form, err := c.MultipartForm()
	if err != nil {
		ErrorHandler(c, err)
	}

	tweet, err := h.usecase.CreateTweet(c.Request.Context(), request, form)
	if err != nil {
		ErrorHandler(c, err)
		return
	}

	fmt.Println("----???")
	SuccessResponse(c, ToCreateTweetResponse(tweet))
}

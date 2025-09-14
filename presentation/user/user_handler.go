package presentation

import (
	"twitter-clone-go/apperrors"
	"twitter-clone-go/interface/http"
	"twitter-clone-go/presentation/user/services"
	"twitter-clone-go/request"
	"twitter-clone-go/response"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service services.UserServicer
}

func NewUserHandler(s services.UserServicer) *UserHandler {
	return &UserHandler{service: s}
}

func (sc *UserHandler) Home(c *gin.Context) {
	c.IndentedJSON(200, gin.H{"message": "Hello World!"})
}

func (sc *UserHandler) HealthCheck(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "ok",
	})
}

func (sc *UserHandler) GetUserListHandler(c *gin.Context) {
	users, err := sc.service.GetUserList()
	if err != nil {
		http.ErrorHandler(c, err)
		return
	}
	response.SuccessResponse(c, users)
}

func (sc *UserHandler) SignUpHandler(c *gin.Context) {
	var request request.SignUpInfo
	if err := c.BindJSON(&request); err != nil {
		err = apperrors.ReqBodyDecodeFailed.Wrap(err, "bad request body")
		http.ErrorHandler(c, err)
		return
	}
	if err := sc.service.SignUp(c.Request.Context(), request); err != nil {
		http.ErrorHandler(c, err)
		return
	}
	response.SuccessResponse(c, nil)
}

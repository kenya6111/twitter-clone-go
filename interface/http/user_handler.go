package http

import (
	"twitter-clone-go/apperrors"
	application "twitter-clone-go/application/user"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	usecase application.UserUsecase
}

func NewUserHandler(u application.UserUsecase) *UserHandler {
	return &UserHandler{usecase: u}
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
	users, err := sc.usecase.GetUserList()
	if err != nil {
		ErrorHandler(c, err)
		return
	}
	SuccessResponse(c, users)
}

func (sc *UserHandler) SignUpHandler(c *gin.Context) {
	var request application.SignUpInfo
	if err := c.BindJSON(&request); err != nil {
		err = apperrors.ReqBodyDecodeFailed.Wrap(err, "bad request body")
		ErrorHandler(c, err)
		return
	}
	if err := sc.usecase.SignUp(c.Request.Context(), request); err != nil {
		ErrorHandler(c, err)
		return
	}
	SuccessResponse(c, nil)
}

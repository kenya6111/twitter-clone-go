package controllers

import (
	"net/http"
	"twitter-clone-go/apperrors"
	"twitter-clone-go/controllers/services"
	"twitter-clone-go/request"
	"twitter-clone-go/response"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	service services.UserServicer
}

func NewUserController(s services.UserServicer) *UserController {
	return &UserController{service: s}
}

func (sc *UserController) Home(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Hello World!"})
}

func (sc *UserController) HealthCheck(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "ok",
	})
}

func (sc *UserController) GetUserListHandler(c *gin.Context) {
	users, err := sc.service.GetUserList()
	if err != nil {
		apperrors.ErrorHandler(c, err)
		return
	}
	response.SuccessResponse(c, users)
}

func (sc *UserController) SignUpHandler(c *gin.Context) {
	var request request.SignUpInfo
	if err := c.BindJSON(&request); err != nil {
		err = apperrors.ReqBodyDecodeFailed.Wrap(err, "bad request body")
		apperrors.ErrorHandler(c, err)
		return
	}
	if err := sc.service.SignUp(c.Request.Context(), request); err != nil {
		apperrors.ErrorHandler(c, err)
		return
	}
	response.SuccessResponse(c, nil)
}

package controllers

import (
	"net/http"
	"twitter-clone-go/apperrors"
	"twitter-clone-go/controllers/services"
	"twitter-clone-go/request"
	"twitter-clone-go/response"
	"twitter-clone-go/usecase/dto"
	"twitter-clone-go/validations"

	"github.com/gin-gonic/gin"
)

type SessionController struct {
	service services.SessionServicer
}

func NewSessionController(s services.SessionServicer) *SessionController {
	return &SessionController{service: s}
}

func (sc *SessionController) Home(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Hello World!"})
}

func (sc *SessionController) HealthCheck(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "ok",
	})
}

func (sc *SessionController) GetUserListHandler(c *gin.Context) {
	users, err := sc.service.GetUserList()
	if err != nil {
		apperrors.ErrorHandler(c, err)
		return
	}
	response.SuccessResponse(c, users)
}

func (sc *SessionController) SignUpHandler(c *gin.Context) {
	var signUpInfo request.SignUpInfo
	if err := c.BindJSON(&signUpInfo); err != nil {
		err = apperrors.ReqBodyDecodeFailed.Wrap(err, "bad request body")
		apperrors.ErrorHandler(c, err)
		return
	}

	if err := validations.ValidateSignUpInfo(signUpInfo); err != nil {
		err = apperrors.ReqBodyDecodeFailed.Wrap(err, "bad request body")
		apperrors.ErrorHandler(c, err)
		return
	}

	if err := sc.service.SignUp(c, sc.toSignUpDto(&signUpInfo)); err != nil {
		apperrors.ErrorHandler(c, err)
		return
	}
	response.SuccessResponse(c, nil)
}

func (u *SessionController) toSignUpDto(signUpInfo *request.SignUpInfo) dto.SignUpInfo {
	return dto.SignUpInfo{
		Email:           signUpInfo.Email,
		Password:        signUpInfo.Password,
		ConfirmPassword: signUpInfo.ConfirmPassword,
	}
}

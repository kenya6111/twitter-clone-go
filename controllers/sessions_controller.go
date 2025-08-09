package controllers

import (
	"net/http"
	"twitter-clone-go/apperrors"
	"twitter-clone-go/request"
	"twitter-clone-go/response"
	"twitter-clone-go/services"
	"twitter-clone-go/validations"

	"github.com/gin-gonic/gin"
)

func Home(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Hello World!"})
}

func HealthCheck(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "ok",
	})
}

func GetUserListHandler(c *gin.Context) {
	users, err := services.GetUserListService(c)
	if err != nil {
		apperrors.ErrorHandler(c, err)
		return
	}
	response.SuccessResponse(c, users)
}

func SignUpHandler(c *gin.Context) {
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

	if err := services.SignUpService(c, signUpInfo); err != nil {
		apperrors.ErrorHandler(c, err)
		return
	}
	response.SuccessResponse(c, nil)
}

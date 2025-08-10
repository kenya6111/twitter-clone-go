package controllers

import (
	"net/http"
	"twitter-clone-go/apperrors"
	"twitter-clone-go/request"
	"twitter-clone-go/response"
	"twitter-clone-go/validations"

	"github.com/gin-gonic/gin"
)

func (con *MyAppController) Home(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Hello World!"})
}

func (con *MyAppController) HealthCheck(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "ok",
	})
}

func (con *MyAppController) GetUserListHandler(c *gin.Context) {
	users, err := con.svc.GetUserListService(c)
	if err != nil {
		apperrors.ErrorHandler(c, err)
		return
	}
	response.SuccessResponse(c, users)
}

func (con *MyAppController) SignUpHandler(c *gin.Context) {
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

	if err := con.svc.SignUpService(c, signUpInfo); err != nil {
		apperrors.ErrorHandler(c, err)
		return
	}
	response.SuccessResponse(c, nil)
}

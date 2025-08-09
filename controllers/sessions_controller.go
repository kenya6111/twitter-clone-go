package controllers

import (
	"net/http"
	"twitter-clone-go/apperrors"
	"twitter-clone-go/request"
	"twitter-clone-go/response"
	"twitter-clone-go/services"
	validation "twitter-clone-go/validation"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

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

	validate = validator.New(validator.WithRequiredStructEnabled())
	validate.RegisterValidation("has_kigou", validation.HasKigou)
	validate.RegisterValidation("has_han_su", validation.HasHanSu)
	validate.RegisterValidation("has_lower_ei", validation.HasLowerEi)
	validate.RegisterValidation("has_upper_ei", validation.HasUpperEi)

	if err := validate.Struct(signUpInfo); err != nil {
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

func ActivateHandler(c *gin.Context) {
	token := c.Query("token")
	if token == "" {
		err := apperrors.BadParam.Wrap(services.ErrNoRequestParam, "bad request param")
		apperrors.ErrorHandler(c, err)
	}
	userId := sessions.Default(c).Get("id").(int32)

	if err := services.ActivateService(c, token, userId); err != nil {
		apperrors.ErrorHandler(c, err)
		return
	}
	response.SuccessResponse(c, nil)
}

package controllers

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"
	"twitter-clone-go/repository"
	"twitter-clone-go/request"
	"twitter-clone-go/response"
	"twitter-clone-go/services"
	validation "twitter-clone-go/validation"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

// var (
// 	hostname = "mailcatcher"
// 	port     = 1025
// 	username = "user@example.com"
// 	password = "password"
// )

// type SignUpInfo struct {
// 	Email           string `validate:"required,email"`
// 	Password        string `validate:"required,gte=8,has_kigou,has_han_su,has_lower_ei,has_upper_ei"`
// 	ConfirmPassword string `validate:"required,gte=8,has_kigou,has_han_su,has_lower_ei,has_upper_ei"`
// }

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
		log.Println("[ERROR]GetUserListService:", err)
		response.ErrorResponse(c, http.StatusInternalServerError, "ユーザ一覧の取得に失敗しました")
	}
	response.SuccessResponse(c, users)
}

func SignUpHandler(c *gin.Context) {

	var signUpInfo request.SignUpInfo
	if err := c.BindJSON(&signUpInfo); err != nil {
		log.Println(err)
		return
	}

	validate = validator.New(validator.WithRequiredStructEnabled())
	validate.RegisterValidation("has_kigou", validation.HasKigou)
	validate.RegisterValidation("has_han_su", validation.HasHanSu)
	validate.RegisterValidation("has_lower_ei", validation.HasLowerEi)
	validate.RegisterValidation("has_upper_ei", validation.HasUpperEi)

	if err := validate.Struct(signUpInfo); err != nil {
		log.Println(err)
		response.ErrorResponse(c, http.StatusBadRequest, "バリデーションエラー")
		return
	}

	if err := services.SignUpService(c, signUpInfo); err != nil {
		switch {
		case errors.Is(err, services.ErrEmailExists):
			response.ErrorResponse(c, http.StatusBadRequest, "既に登録済のメールアドレスです")
		case errors.Is(err, services.ErrPasswordMismatch):
			response.ErrorResponse(c, http.StatusBadRequest, "パスワードが一致していません")
		default:
			response.ErrorResponse(c, http.StatusInternalServerError, "登録処理中にエラーが発生しました")
		}
		return
	}
	response.SuccessResponse(c, nil)
}

func ActivateHandler(c *gin.Context) {
	token := c.Query("token")
	if token == "" {
		response.ErrorResponse(c, http.StatusBadRequest, "不正なアクセス")
		return
	}
	userId := sessions.Default(c).Get("id").(int32)

	result, err := repository.GetEmailVerifyToken(c, userId, token, time.Now().Add(time.Hour*24))
	if err != nil {
		fmt.Println(err)
		response.ErrorResponse(c, http.StatusBadRequest, "認証失敗")
		return
	}

	if result == nil {
		response.ErrorResponse(c, http.StatusBadRequest, "認証メール未発行・もしくは認証済みのユーザ")
		return
	}

	ok := repository.UpdateUser(c, userId)
	if !ok {
		response.ErrorResponse(c, http.StatusBadRequest, "認証失敗")
		return
	}

	ok = repository.DeleteEmailVerifyToken(c, token)
	if !ok {
		response.ErrorResponse(c, http.StatusBadRequest, "認証失敗")
		return
	}

	fmt.Println("完了")
}

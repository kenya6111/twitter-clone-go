package controllers

import (
	"fmt"
	"net/http"
	"twitter-clone-go/repository"
	validation "twitter-clone-go/validation"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

var validate *validator.Validate

type signUpInfo struct {
	Email           string `validate:"required,email"`
	Password        string `validate:"required,gte=8,has_kigou,has_han_su,has_lower_ei,has_upper_ei"`
	ConfirmPassword string `validate:"required,gte=8,has_kigou,has_han_su,has_lower_ei,has_upper_ei"`
}

func Home(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Hello World!"})
}
func HealthCheck(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "ok",
	})
}
func SelectUsers(c *gin.Context) {
	user, err := repository.SelectUsers(c)
	fmt.Println("⭐️")
	fmt.Println(user, err)
}

func SignUp(c *gin.Context) {

	var newSignUpInfo signUpInfo
	if err := c.BindJSON(&newSignUpInfo); err != nil {
		fmt.Println(err)
		return
	}

	validate = validator.New(validator.WithRequiredStructEnabled())
	validate.RegisterValidation("has_kigou", validation.HasKigou)
	validate.RegisterValidation("has_han_su", validation.HasHanSu)
	validate.RegisterValidation("has_lower_ei", validation.HasLowerEi)
	validate.RegisterValidation("has_upper_ei", validation.HasUpperEi)

	signUpInfo := &signUpInfo{
		Email:           newSignUpInfo.Email,
		Password:        newSignUpInfo.Password,
		ConfirmPassword: newSignUpInfo.ConfirmPassword,
	}

	errors := validate.Struct(signUpInfo)
	if errors != nil {
		fmt.Println(errors)
		return
	}

	user, _ := repository.GetUserByEmail(c, signUpInfo.Email)
	if user == nil {
		fmt.Println("ユニーク！")
		fmt.Println(user)
	}

	if signUpInfo.Password != signUpInfo.ConfirmPassword {
		fmt.Println("パスワードが異なります。同じパスワードを入力してください")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(signUpInfo.Password), bcrypt.DefaultCost)
	fmt.Println(string(hash))
	fmt.Println(err)
	if err != nil {
		return
	}

	if repository.CreateUser(c, signUpInfo.Email, hash) {
		fmt.Println("サインアップ成功！")
	} else {
		fmt.Println("サインアップ失敗。。。")
	}
}

package controllers

import (
	"fmt"
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)
var validate *validator.Validate

type signUpInfo struct {
	Email string     `validate:"required,email"`
    Password  string `validate:"required,gte=8,has_kigou,has_han_su,has_lower_ei,has_upper_ei"`
}

func Home(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Hello World!!!!!"})
}
func HealthCheck(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "ok",
	})
}

func SignUp(c *gin.Context) {

	var newSignUpInfo signUpInfo;

	if err := c.BindJSON(&newSignUpInfo); err != nil {
		fmt.Println(err)
            return
    }

	validate = validator.New(validator.WithRequiredStructEnabled())
	validate.RegisterValidation("has_kigou", hasKigou)
	validate.RegisterValidation("has_han_su", hasHanSu)
	validate.RegisterValidation("has_lower_ei", hasLowerEi)
	validate.RegisterValidation("has_upper_ei", hasUpperEi)

	signUpInfo := &signUpInfo{
		Email: newSignUpInfo.Email,
		Password: newSignUpInfo.Password,
	}

	errors := validate.Struct(signUpInfo)
	if errors != nil {
		fmt.Println(errors)
		return
	}

	
}

func hasKigou(fl validator.FieldLevel) bool {  //引数の型、返り値は固定
	pw := fl.Field().String()
	hasKigou := regexp.MustCompile(`[-_!?]`).MatchString(pw)
	return hasKigou
}

func hasHanSu(fl validator.FieldLevel) bool {  //引数の型、返り値は固定
	pw := fl.Field().String()
	hasSu := regexp.MustCompile(`[0-9]`).MatchString(pw)
	return hasSu
}

func hasLowerEi(fl validator.FieldLevel) bool {  //引数の型、返り値は固定
	pw := fl.Field().String()
	hasLowerEi := regexp.MustCompile(`[a-z]`).MatchString(pw)
	return hasLowerEi
}

func hasUpperEi(fl validator.FieldLevel) bool {  //引数の型、返り値は固定
	pw := fl.Field().String()
	hasUpperEi := regexp.MustCompile(`[A-Z]`).MatchString(pw)

	return hasUpperEi
}
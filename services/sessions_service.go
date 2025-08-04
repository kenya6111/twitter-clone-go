package services

import (
	"errors"
	"fmt"
	"log"
	"time"
	"twitter-clone-go/common"
	"twitter-clone-go/repository"
	"twitter-clone-go/request"
	db "twitter-clone-go/tutorial"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrEmailExists      = errors.New("email already exists")
	ErrPasswordMismatch = errors.New("passwords do not match")
)

func GetUserListService(c *gin.Context) (*[]db.User, error) {
	users, err := repository.SelectUsers(c)
	if err != nil {
		return nil, fmt.Errorf("ユーザ一覧取得失敗 %w", err)
	}
	return users, nil
}

func SignUpService(c *gin.Context, signUpInfo request.SignUpInfo) error {

	user, err := repository.GetUserByEmail(c, signUpInfo.Email)
	if err != nil {
		log.Println(err)
		return err
	}
	if user != nil {
		return fmt.Errorf("duplicate error:%s is already exist", user.Email)
	}

	if signUpInfo.Password != signUpInfo.ConfirmPassword {
		return fmt.Errorf("mismatch password: %s,%s", signUpInfo.Password, signUpInfo.ConfirmPassword)
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(signUpInfo.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
		return err
	}

	createdUser, err := repository.CreateUser(c, signUpInfo.Email, hash)
	if err != nil {
		log.Println(err)
		return err
	}

	token, _ := common.GenerateSecureToken(32)

	if err := common.SendMail(token); err != nil {
		log.Println(err)
		return err
	}

	verified, err := repository.CreateEmailVerifyToken(c, createdUser.ID, token, time.Now())
	if err != nil {
		log.Println(err)
		return err
	}

	fmt.Println(verified)
	session := sessions.Default(c)
	session.Set("id", createdUser.ID)
	session.Save()
	return nil
}

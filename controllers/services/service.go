package services

import (
	"twitter-clone-go/application/user/dto"
	domain "twitter-clone-go/domain/user"

	"github.com/gin-gonic/gin"
)

type SessionServicer interface {
	GetUserList() ([]domain.User, error)
	SignUp(c *gin.Context, signUpInfo dto.SignUpInfo) error
}

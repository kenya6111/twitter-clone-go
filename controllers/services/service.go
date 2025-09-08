package services

import (
	domain "twitter-clone-go/domain/user"
	"twitter-clone-go/usecase/dto"

	"github.com/gin-gonic/gin"
)

type SessionServicer interface {
	GetUserList() ([]domain.User, error)
	SignUp(c *gin.Context, signUpInfo dto.SignUpInfo) error
}

package services

import (
	domain "twitter-clone-go/domain/user"
	"twitter-clone-go/request"

	"github.com/gin-gonic/gin"
)

type SessionServicer interface {
	GetUserListService() ([]domain.User, error)
	SignUpService(c *gin.Context, signUpInfo request.SignUpInfo) error
}

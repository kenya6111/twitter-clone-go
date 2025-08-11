package services

import (
	"twitter-clone-go/request"
	"twitter-clone-go/tutorial"

	"github.com/gin-gonic/gin"
)

type SessionServicer interface {
	GetUserListService() ([]tutorial.User, error)
	SignUpService(c *gin.Context, signUpInfo request.SignUpInfo) error
}

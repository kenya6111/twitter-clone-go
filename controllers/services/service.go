package services

import (
	"twitter-clone-go/request"
	db "twitter-clone-go/tutorial"

	"github.com/gin-gonic/gin"
)

type SessionServicer interface {
	GetUserListService() ([]db.User, error)
	SignUpService(c *gin.Context, signUpInfo request.SignUpInfo) error
}

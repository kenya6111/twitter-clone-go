package services

import (
	"context"
	domain "twitter-clone-go/domain/user"
	"twitter-clone-go/request"
)

type UserServicer interface {
	GetUserList() ([]domain.User, error)
	SignUp(c context.Context, signUpInfo request.SignUpInfo) error
}

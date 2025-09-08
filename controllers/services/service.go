package services

import (
	"context"
	"twitter-clone-go/application/user/dto"
	domain "twitter-clone-go/domain/user"
)

type UserServicer interface {
	GetUserList() ([]domain.User, error)
	SignUp(c context.Context, signUpInfo dto.SignUpInfo) error
}

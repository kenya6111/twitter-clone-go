package testdata

import (
	"twitter-clone-go/request"
	"twitter-clone-go/tutorial"

	"github.com/gin-gonic/gin"
)

type serviceMock struct{}

func NewServiceMock() *serviceMock {
	return &serviceMock{}
}

func (s *serviceMock) GetUserListService() ([]tutorial.User, error) {
	return UserTestData, nil
}

func (s *serviceMock) SignUpService(c *gin.Context, signUpInfo request.SignUpInfo) error {
	return nil
}

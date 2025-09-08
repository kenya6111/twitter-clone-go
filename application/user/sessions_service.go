package usecase

import (
	"context"
	"twitter-clone-go/apperrors"
	"twitter-clone-go/application/user/dto"
	"twitter-clone-go/common"
	domain "twitter-clone-go/domain/user"
	"twitter-clone-go/infrasctructure/postgres"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repository *postgres.UserRepository
	tx         *postgres.Transaction
}

func NewUserService(r *postgres.UserRepository, tx *postgres.Transaction) *UserService {
	return &UserService{repository: r, tx: tx}
}

func (ss *UserService) GetUserList() ([]domain.User, error) {
	users, err := ss.repository.FindAll()
	if err != nil {
		err = apperrors.GetDataFailed.Wrap(err, "fail to get users data")
		return nil, err
	}

	if len(users) == 0 {
		err = apperrors.NAData.Wrap(ErrNoData, "no data")
		return nil, err
	}
	return users, nil
}

func (ss *UserService) SignUp(c *gin.Context, signUpInfo dto.SignUpInfo) error {
	ctx := c.Request.Context()
	var token string
	var createdUser *domain.User

	user, err := ss.repository.CountByEmail(signUpInfo.Email)
	if err != nil {
		return apperrors.GetDataFailed.Wrap(ErrNoData, "fail to get user by email")
	}

	if user > 0 {
		return apperrors.DuplicateData.Wrap(ErrDuplicateData, "already exist user data")
	}

	if signUpInfo.Password != signUpInfo.ConfirmPassword {
		return apperrors.GetDataFailed.Wrap(ErrMismatchData, "mismatch password and confirmPassword")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(signUpInfo.Password), bcrypt.DefaultCost)
	if err != nil {
		return apperrors.GetDataFailed.Wrap(err, "fail to generate has value from password")
	}

	err = ss.tx.Do(ctx, func(ctx context.Context) error {
		createdUser, err = ss.repository.CreateUser(ctx, signUpInfo.Email, hash)
		if err != nil {
			return apperrors.InsertDataFailed.Wrap(err, "fail to insert user ")
		}

		token, _ = common.GenerateSecureToken(32)

		_, err := ss.repository.CreateEmailVerifyToken(ctx, createdUser.ID, token)
		if err != nil {
			return apperrors.InsertDataFailed.Wrap(err, "fail to insert emailVerifyToken")
		}
		return nil
	})
	if err != nil {
		return err
	}

	if err := common.SendMail(token, signUpInfo.Email); err != nil {
		return apperrors.GenerateTokenFailed.Wrap(err, "fail to send invitation mail")
	}

	return nil
}

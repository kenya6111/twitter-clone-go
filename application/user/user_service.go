package usecase

import (
	"context"
	"twitter-clone-go/apperrors"
	"twitter-clone-go/common"
	domain "twitter-clone-go/domain/user"
	"twitter-clone-go/infrasctructure/postgres"
	"twitter-clone-go/request"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo *postgres.UserRepository
	tx   *postgres.Transaction
	dSer *postgres.UserDomainService
}

func NewUserService(r *postgres.UserRepository, tx *postgres.Transaction, dSer domain.UserDomainService) *UserService {
	return &UserService{repo: r, tx: tx}
}

func (ss *UserService) GetUserList() ([]domain.User, error) {
	users, err := ss.repo.FindAll()
	if err != nil {
		err = apperrors.GetDataFailed.Wrap(err, "fail to get users data")
		return nil, err
	}

	if len(users) == 0 {
		err = apperrors.NAData.Wrap(apperrors.ErrNoData, "no data")
		return nil, err
	}
	return users, nil
}

func (us *UserService) SignUp(ctx context.Context, request request.SignUpInfo) error {
	var token string
	var createdUser *domain.User

	user, err := domain.NewUser(request.Email, request.Password, request.ConfirmPassword)
	if err != nil {
		return err
	}
	if err := us.dSer.IsDuplicatedEmail(user.Email); err != nil {
		return err
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password.Value()), bcrypt.DefaultCost)
	if err != nil {
		return apperrors.GetDataFailed.Wrap(err, "fail to generate has value from password")
	}

	err = us.tx.Do(ctx, func(ctx context.Context) error {
		createdUser, err = us.repo.CreateUser(ctx, user.Email, hash)
		if err != nil {
			return apperrors.InsertDataFailed.Wrap(err, "fail to insert user ")
		}

		token, _ = common.GenerateSecureToken(32)

		_, err := us.repo.CreateEmailVerifyToken(ctx, createdUser.ID, token)
		if err != nil {
			return apperrors.InsertDataFailed.Wrap(err, "fail to insert emailVerifyToken")
		}
		return nil
	})
	if err != nil {
		return err
	}

	if err := common.SendMail(token, user.Email); err != nil {
		return apperrors.GenerateTokenFailed.Wrap(err, "fail to send invitation mail")
	}

	return nil
}

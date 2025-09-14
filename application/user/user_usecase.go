package application

import (
	"context"
	"twitter-clone-go/apperrors"
	domain "twitter-clone-go/domain"
	"twitter-clone-go/domain/service"
	userDomain "twitter-clone-go/domain/user"
	"twitter-clone-go/pkg/crypt"
	"twitter-clone-go/request"

	"golang.org/x/crypto/bcrypt"
)

type UserUsecaseImpl struct {
	repo         userDomain.UserRepository
	tx           domain.Transaction
	dSer         userDomain.UserDomainService
	emailService service.EmailService
}
type UserUsecase interface {
	GetUserList() ([]userDomain.User, error)
	SignUp(c context.Context, signUpInfo request.SignUpInfo) error
}

func NewUserUsecase(r userDomain.UserRepository, tx domain.Transaction, dSer userDomain.UserDomainService, emailService service.EmailService) *UserUsecaseImpl {
	return &UserUsecaseImpl{repo: r, tx: tx, dSer: dSer, emailService: emailService}
}

func (ss *UserUsecaseImpl) GetUserList() ([]userDomain.User, error) {
	users, err := ss.repo.FindAll()
	if err != nil {
		err = apperrors.GetDataFailed.Wrap(err, "fail to get users data")
		return nil, err
	}
	return users, nil
}

func (us *UserUsecaseImpl) SignUp(ctx context.Context, request request.SignUpInfo) error {
	var token string
	var createdUser *userDomain.User

	user, err := userDomain.NewUser(request.Name, request.Email, request.Password, request.ConfirmPassword)
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

		token, _ = crypt.GenerateSecureToken(32)

		_, err := us.repo.CreateEmailVerifyToken(ctx, createdUser.ID, token)
		if err != nil {
			return apperrors.InsertDataFailed.Wrap(err, "fail to insert emailVerifyToken")
		}
		return nil
	})
	if err != nil {
		return err
	}

	if err := us.emailService.SendInvitationEmail(token, user.Email); err != nil {
		return apperrors.GenerateTokenFailed.Wrap(err, "fail to send invitation mail")
	}

	return nil
}

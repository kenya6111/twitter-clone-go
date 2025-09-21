package application

import (
	"context"
	"twitter-clone-go/apperrors"
	"twitter-clone-go/domain"
	"twitter-clone-go/domain/service"
	"twitter-clone-go/pkg/crypt"

	"golang.org/x/crypto/bcrypt"
)

type SignUpInfo struct {
	Name            string
	Email           string
	Password        string
	ConfirmPassword string
}

type UserUsecaseImpl struct {
	repo         domain.UserRepository
	tx           domain.Transaction
	dSer         domain.UserDomainService
	emailService service.EmailService
}
type UserUsecase interface {
	GetUserList() ([]domain.User, error)
	SignUp(c context.Context, signUpInfo SignUpInfo) error
}

func NewUserUsecase(r domain.UserRepository, tx domain.Transaction, dSer domain.UserDomainService, emailService service.EmailService) *UserUsecaseImpl {
	return &UserUsecaseImpl{repo: r, tx: tx, dSer: dSer, emailService: emailService}
}

func (u *UserUsecaseImpl) GetUserList() ([]domain.User, error) {
	users, err := u.repo.FindAll()
	if err != nil {
		err = apperrors.GetDataFailed.Wrap(err, "fail to get users data")
		return nil, err
	}
	return users, nil
}

func (u *UserUsecaseImpl) SignUp(ctx context.Context, request SignUpInfo) error {
	var token string
	var createdUser *domain.User

	user, err := domain.NewUser(request.Name, request.Email, request.Password, request.ConfirmPassword)
	if err != nil {
		return err
	}
	if err := u.dSer.IsDuplicatedEmail(user.Email); err != nil {
		return err
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password.Value()), bcrypt.DefaultCost)
	if err != nil {
		return apperrors.GetDataFailed.Wrap(err, "fail to generate has value from password")
	}

	err = u.tx.Do(ctx, func(ctx context.Context) error {
		createdUser, err = u.repo.CreateUser(ctx, user.Email, hash)
		if err != nil {
			return apperrors.InsertDataFailed.Wrap(err, "fail to insert user ")
		}

		token, _ = crypt.GenerateSecureToken(32)

		_, err := u.repo.CreateEmailVerifyToken(ctx, createdUser.ID, token)
		if err != nil {
			return apperrors.InsertDataFailed.Wrap(err, "fail to insert emailVerifyToken")
		}
		return nil
	})
	if err != nil {
		return err
	}

	if err := u.emailService.SendInvitationEmail(token, user.Email); err != nil {
		return apperrors.GenerateTokenFailed.Wrap(err, "fail to send invitation mail")
	}

	return nil
}

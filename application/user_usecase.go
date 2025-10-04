package application

import (
	"context"
	"time"
	"twitter-clone-go/apperrors"
	"twitter-clone-go/domain"
	"twitter-clone-go/domain/service"
)

type SignUpInfo struct {
	Name            string `json:"name" binding:"required"`
	Email           string `json:"email" binding:"required,email"`
	Password        string `json:"password" binding:"required"`
	ConfirmPassword string `json:"confirmPassword" binding:"required"`
}

type ActivateInfo struct {
	Token string `json:"token" binding:"required"`
}

type UserUsecaseImpl struct {
	userRepo          domain.UserRepository
	emailVerifyRepo   domain.EmailVerifyTokenRepository
	transaction       domain.Transaction
	userDomainService domain.UserDomainService
	emailService      service.EmailService
	passwordHasher    service.PasswordHasher
}
type UserUsecase interface {
	GetUserList() ([]domain.User, error)
	SignUp(c context.Context, signUpInfo SignUpInfo) error
	Activate(ctx context.Context, request string) error
}

func NewUserUsecase(
	userRepo domain.UserRepository,
	emailVerifyRepo domain.EmailVerifyTokenRepository,
	transaction domain.Transaction,
	userDomainService domain.UserDomainService,
	emailService service.EmailService,
	passwordHasher service.PasswordHasher,
) *UserUsecaseImpl {
	return &UserUsecaseImpl{
		userRepo:          userRepo,
		emailVerifyRepo:   emailVerifyRepo,
		transaction:       transaction,
		userDomainService: userDomainService,
		emailService:      emailService,
		passwordHasher:    passwordHasher,
	}
}

func (u *UserUsecaseImpl) GetUserList() ([]domain.User, error) {
	users, err := u.userRepo.FindAll()
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
	if err := u.userDomainService.IsDuplicatedEmail(ctx, user.Email); err != nil {
		return err
	}

	hash, err := u.passwordHasher.HashPassword(user.Password.Value())
	if err != nil {
		return apperrors.GenerateHashFailed.Wrap(err, "fail to generate hash value from password")
	}

	err = u.transaction.Do(ctx, func(ctx context.Context) error {
		createdUser, err = u.userRepo.CreateUser(ctx, user.Name, user.Email, hash)
		if err != nil {
			return apperrors.InsertDataFailed.Wrap(err, "fail to insert user ")
		}

		token, err = u.passwordHasher.GenerateSecureToken(32)
		if err != nil {
			return apperrors.GenerateTokenFailed.Wrap(err, "fail to generate secure token ")
		}

		_, err := u.emailVerifyRepo.CreateEmailVerifyToken(ctx, createdUser.ID, token)
		if err != nil {
			return apperrors.InsertDataFailed.Wrap(err, "fail to insert emailVerifyToken")
		}
		return nil
	})
	if err != nil {
		return err
	}

	if err := u.emailService.SendInvitationEmail(user.Email, token); err != nil {
		return apperrors.SendEmailFailed.Wrap(err, "fail to send invitation mail")
	}

	return nil
}

func (u *UserUsecaseImpl) Activate(ctx context.Context, token string) error {
	err := u.transaction.Do(ctx, func(ctx context.Context) error {
		result, err := u.emailVerifyRepo.GetEmailVerifyToken(ctx, token, time.Now().Add(time.Hour*24))
		if err != nil {
			err = apperrors.GetDataFailed.Wrap(err, "fail to get emailVerifyToken")
			return err
		}

		if result == nil {
			err = apperrors.BadParam.Wrap(apperrors.ErrNoData, "User with no verification email sent, or already verified")
			return err
		}

		if domain.IsExpired(result.ExpiresAt) {
			err = apperrors.BadParam.Wrap(apperrors.ErrEmailVerifyTokenExpired, "email verify token is already expired")
			return err
		}

		userId := result.UserID

		err = u.userRepo.UpdateUser(ctx, userId)
		if err != nil {
			err = apperrors.UpdateDataFailed.Wrap(err, "fail to activate user")
			return err
		}

		err = u.emailVerifyRepo.DeleteEmailVerifyToken(ctx, token)
		if err != nil {
			err = apperrors.DeleteDataFailed.Wrap(err, "fail to delete email verify token")
			return err
		}
		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

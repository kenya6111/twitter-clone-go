package application

import (
	"context"
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

type UserUsecaseImpl struct {
	userRepo          domain.UserRepository
	transaction       domain.Transaction
	userDomainService domain.UserDomainService
	emailService      service.EmailService
	passwordHasher    service.PasswordHasher
}
type UserUsecase interface {
	GetUserList() ([]domain.User, error)
	SignUp(c context.Context, signUpInfo SignUpInfo) error
}

func NewUserUsecase(
	userRepo domain.UserRepository,
	transaction domain.Transaction,
	userDomainService domain.UserDomainService,
	emailService service.EmailService,
	passwordHasher service.PasswordHasher,
) *UserUsecaseImpl {
	return &UserUsecaseImpl{
		userRepo:          userRepo,
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
		createdUser, err = u.userRepo.CreateUser(ctx, user.Email, hash)
		if err != nil {
			return apperrors.InsertDataFailed.Wrap(err, "fail to insert user ")
		}

		token, err = u.passwordHasher.GenerateSecureToken(32)
		if err != nil {
			return apperrors.GenerateTokenFailed.Wrap(err, "fail to generate secure token ")
		}

		_, err := u.userRepo.CreateEmailVerifyToken(ctx, createdUser.ID, token)
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

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
type LoginInfo struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}
type UserUsecaseImpl struct {
	userRepo          domain.UserRepository
	emailVerifyRepo   domain.EmailVerifyTokenRepository
	transaction       domain.Transaction
	userDomainService domain.UserDomainService
	emailService      service.EmailService
	passwordHasher    service.PasswordHasher
	sessionStore      service.SessionStore
}
type UserUsecase interface {
	GetUserList() ([]domain.User, error)
	SignUp(c context.Context, signUpInfo SignUpInfo) error
	Activate(ctx context.Context, request string) error
	Login(ctx context.Context, request LoginInfo) (*domain.User, error)
	Logout(ctx context.Context) error
}

func NewUserUsecase(
	userRepo domain.UserRepository,
	emailVerifyRepo domain.EmailVerifyTokenRepository,
	transaction domain.Transaction,
	userDomainService domain.UserDomainService,
	emailService service.EmailService,
	passwordHasher service.PasswordHasher,
	sessionStore service.SessionStore,
) *UserUsecaseImpl {
	return &UserUsecaseImpl{
		userRepo:          userRepo,
		emailVerifyRepo:   emailVerifyRepo,
		transaction:       transaction,
		userDomainService: userDomainService,
		emailService:      emailService,
		passwordHasher:    passwordHasher,
		sessionStore:      sessionStore,
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

		_, err := u.emailVerifyRepo.Save(ctx, createdUser.ID, token)
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
		result, err := u.emailVerifyRepo.FindByToken(ctx, token)
		if err != nil {
			err = apperrors.GetDataFailed.Wrap(err, "fail to get emailVerifyToken")
			return err
		}

		if result == nil {
			err = apperrors.BadParam.Wrap(apperrors.ErrNoData, "User with no verification email sent, or already verified")
			return err
		}

		emailVerifyToken, err := domain.ReconstructEmailVerifyToken(result.ID, result.UserID, result.Token, result.ExpiresAt, result.CreatedAt)
		if err != nil {
			return err
		}

		userId := emailVerifyToken.UserID
		token := emailVerifyToken.Token

		_, err = u.userRepo.ActivateUser(ctx, userId)
		if err != nil {
			err = apperrors.UpdateDataFailed.Wrap(err, "fail to activate user")
			return err
		}

		err = u.emailVerifyRepo.DeleteByToken(ctx, token)
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

func (u *UserUsecaseImpl) Login(ctx context.Context, request LoginInfo) (*domain.User, error) {
	user, err := u.userRepo.FindByEmail(ctx, request.Email)
	if err != nil {
		return nil, apperrors.Unauthorized.Wrap(err, "fail to get user data")
	}

	if !user.IsActive {
		return nil, apperrors.Unauthorized.Wrap(err, "user is not be activated")
	}

	err = u.passwordHasher.CompareHashAndPassword(user.Password.Value(), request.Password)
	if err != nil {
		return nil, apperrors.Unauthorized.Wrap(err, "password is invalid")
	}

	err = u.sessionStore.Set(ctx, user.ID)
	if err != nil {
		return nil, apperrors.Unauthorized.Wrap(err, "")
	}

	return user, nil

}

func (u *UserUsecaseImpl) Logout(ctx context.Context) error {
	err := u.sessionStore.Delete(ctx)
	if err != nil {
		return apperrors.LogoutFailed.Wrap(err, "")
	}
	return nil
}

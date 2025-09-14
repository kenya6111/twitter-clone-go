package usecase

import (
	"context"
	"twitter-clone-go/apperrors"
	"twitter-clone-go/common"
	domain "twitter-clone-go/domain"
	"twitter-clone-go/domain/service"
	userDomain "twitter-clone-go/domain/user"
	"twitter-clone-go/request"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo         userDomain.UserRepository
	tx           domain.Transaction
	dSer         userDomain.UserDomainService
	emailService service.EmailService
}

func NewUserService(r userDomain.UserRepository, tx domain.Transaction, dSer userDomain.UserDomainService, emailService service.EmailService) *UserService {
	return &UserService{repo: r, tx: tx, dSer: dSer, emailService: emailService}
}

func (ss *UserService) GetUserList() ([]userDomain.User, error) {
	users, err := ss.repo.FindAll()
	if err != nil {
		err = apperrors.GetDataFailed.Wrap(err, "fail to get users data")
		return nil, err
	}
	return users, nil
}

func (us *UserService) SignUp(ctx context.Context, request request.SignUpInfo) error {
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

	if err := us.emailService.SendInvitationEmail(token, user.Email); err != nil {
		return apperrors.GenerateTokenFailed.Wrap(err, "fail to send invitation mail")
	}

	return nil
}

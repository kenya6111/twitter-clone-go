package usecase

import (
	"context"
	"time"
	"twitter-clone-go/apperrors"
	"twitter-clone-go/common"
	domain "twitter-clone-go/domain/user"
	"twitter-clone-go/infrasctructure/postgres"
	"twitter-clone-go/usecase/dto"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
	"golang.org/x/crypto/bcrypt"
)

type SessionService struct {
	repository domain.UserRepository
	tx         *postgres.Transaction
}

func NewSessionService(r domain.UserRepository, tx *postgres.Transaction) *SessionService {
	return &SessionService{repository: r, tx: tx}
}

func (ss *SessionService) GetUserListService() ([]domain.User, error) {
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

func (ss *SessionService) SignUpService(c *gin.Context, signUpInfo dto.SignUpInfo) error {
	ctx := c.Request.Context()
	var token string
	var createdUser *domain.User

	user, err := ss.repository.CountByEmail(signUpInfo.Email)
	if err != nil {
		err = apperrors.GetDataFailed.Wrap(ErrNoData, "fail to get user by email")
		return err
	}
	if user > 0 {
		err = apperrors.DuplicateData.Wrap(ErrDuplicateData, "already exist user data")
		return err
	}

	if signUpInfo.Password != signUpInfo.ConfirmPassword {
		err = apperrors.GetDataFailed.Wrap(ErrMismatchData, "mismatch password and confirmPassword")
		return err
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(signUpInfo.Password), bcrypt.DefaultCost)
	if err != nil {
		err = apperrors.GetDataFailed.Wrap(err, "fail to generate has value from password")
		return err
	}

	err = ss.tx.Do(ctx, func(ctx context.Context) error {
		// transaction適用
		createdUser, err = ss.repository.CreateUser(ctx, signUpInfo.Email, hash)
		if err != nil {
			err = apperrors.InsertDataFailed.Wrap(err, "fail to insert user ")
			return err
		}

		token, _ = common.GenerateSecureToken(32)

		expiredAt := pgtype.Timestamp{}
		_ = expiredAt.Scan(time.Now())

		// transaction適用
		_, err := ss.repository.CreateEmailVerifyToken(ctx, createdUser.ID, token, expiredAt)
		if err != nil {
			err = apperrors.InsertDataFailed.Wrap(err, "fail to insert emailVerifyToke")
			return err
		}
		return nil

	})
	if err != nil {
		return err
	}

	// --- トランザクション成功後にメール送信 ---
	if err := common.SendMail(token, signUpInfo.Email); err != nil {
		err = apperrors.GenerateTokenFailed.Wrap(err, "fail to send invitation mail")
		return err
	}

	session := sessions.Default(c)
	session.Set("id", createdUser.ID)
	session.Save()
	return nil
}

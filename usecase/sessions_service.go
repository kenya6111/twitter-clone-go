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
		// transaction適用
		createdUser, err = ss.repository.CreateUser(ctx, signUpInfo.Email, hash)
		if err != nil {
			return apperrors.InsertDataFailed.Wrap(err, "fail to insert user ")
		}

		token, _ = common.GenerateSecureToken(32)

		expiredAt := pgtype.Timestamp{}
		_ = expiredAt.Scan(time.Now())

		// transaction適用
		_, err := ss.repository.CreateEmailVerifyToken(ctx, createdUser.ID, token, expiredAt)
		if err != nil {
			return apperrors.InsertDataFailed.Wrap(err, "fail to insert emailVerifyToke")
		}
		return nil

	})
	if err != nil {
		return err
	}

	// --- トランザクション成功後にメール送信 ---
	if err := common.SendMail(token, signUpInfo.Email); err != nil {
		return apperrors.GenerateTokenFailed.Wrap(err, "fail to send invitation mail")
	}

	session := sessions.Default(c)
	session.Set("id", createdUser.ID)
	if err := session.Save(); err != nil {
		return apperrors.SessionSaveFailed.Wrap(err, "failed to save session")
	}

	return nil
}

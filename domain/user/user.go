package domain

import (
	"context"
	"net/mail"
	"regexp"
	"twitter-clone-go/apperrors"
	"twitter-clone-go/domain"
	"unicode/utf8"

	ulid "github.com/oklog/ulid/v2"
)

type User struct {
	ID       string
	Name     string
	Email    string
	Password password
	IsActive bool
}

type UserRepository interface {
	FindAll() ([]User, error)
	FindByEmail(email string) (*User, error)
	CountByEmail(email string) (int64, error)
	CreateUser(c context.Context, email string, hash []byte) (*User, error)
	CreateEmailVerifyToken(ctx context.Context, userId string, token string) (*EmailVerifyToken, error)
}

func NewUser(name string, email string, password string, confirmPassword string) (*User, error) {
	// name のドメインルール
	if utf8.RuneCountInString(name) < nameLengthMin || utf8.RuneCountInString(name) > nameLengthMax {
		return nil, apperrors.BadParam.Wrap(apperrors.ErrMismatchData, "Name must be between 1 and 255 characters")
	}
	// email のドメインルール
	if _, err := mail.ParseAddress(email); err != nil {
		return nil, apperrors.BadParam.Wrap(err, "Please enter a valid email address")
	}

	// password のドメインルール
	if password != confirmPassword {
		return nil, apperrors.BadParam.Wrap(apperrors.ErrMismatchData, "mismatch password and confirmPassword")
	}
	p, err := NewPassword(password)
	if err != nil {
		return nil, err
	}

	return &User{
		ID:       ulid.Make().String(),
		Name:     name,
		Email:    email,
		Password: p,
		IsActive: false,
	}, nil
}

type password struct {
	value string
}

func NewPassword(pass string) (password, error) {
	if len(pass) < passwordLengthMin {
		return password{}, apperrors.BadParam.Wrap(domain.ErrTooShort, "password must be at least 8 characters")
	}
	if !HasKigou(pass) {
		return password{}, apperrors.BadParam.Wrap(domain.ErrNoHasKigou, "password must not contain symbols (-_!?)")
	}
	if !HasHanSu(pass) {
		return password{}, apperrors.BadParam.Wrap(domain.ErrNoHasHanSu, "password must contain at least one number")
	}
	if !HasLowerEi(pass) {
		return password{}, apperrors.BadParam.Wrap(domain.ErrNoHasLowerEi, "password must contain at least one lowercase letter")
	}
	if !HasUpperEi(pass) {
		return password{}, apperrors.BadParam.Wrap(domain.ErrNoHasUpperEi, "password must contain at least one uppercase letter")
	}
	return password{
		value: pass,
	}, nil

}

func (p password) Value() string {
	return p.value
}

func HasKigou(password string) bool {
	hasKigou := regexp.MustCompile(`[-_!?]`).MatchString(password)
	return hasKigou
}

func HasHanSu(password string) bool {
	hasSu := regexp.MustCompile(`[0-9]`).MatchString(password)
	return hasSu
}

func HasLowerEi(password string) bool {
	hasLowerEi := regexp.MustCompile(`[a-z]`).MatchString(password)
	return hasLowerEi
}

func HasUpperEi(password string) bool {
	hasUpperEi := regexp.MustCompile(`[A-Z]`).MatchString(password)
	return hasUpperEi
}

const (
	nameLengthMax     = 255
	nameLengthMin     = 1
	passwordLengthMin = 8
)

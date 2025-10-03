package domain

import (
	"context"
	"net/mail"
	"regexp"
	"twitter-clone-go/apperrors"
	"unicode/utf8"

	ulid "github.com/oklog/ulid/v2"
)

type User struct {
	ID       string
	Name     string
	Email    string
	Password Password
	IsActive bool
}

type UserRepository interface {
	FindAll() ([]User, error)
	FindByEmail(c context.Context, email string) (*User, error)
	CountByEmail(c context.Context, email string) (int64, error)
	CreateUser(c context.Context, name string, email string, hash string) (*User, error)
	ActivateUser(ctx context.Context, userId string) (*User, error)
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

func ReconstructUser(id, name, email, hashedPassword string, isActive bool) (*User, error) {
	if id == "" {
		return nil, apperrors.BadParam.Wrap(apperrors.ErrMismatchData, "id must not be empty")
	}
	if utf8.RuneCountInString(name) < nameLengthMin || utf8.RuneCountInString(name) > nameLengthMax {
		return nil, apperrors.BadParam.Wrap(apperrors.ErrMismatchData, "Name must be between 1 and 255 characters")
	}
	if _, err := mail.ParseAddress(email); err != nil {
		return nil, apperrors.BadParam.Wrap(err, "invalid email")
	}

	password, err := NewPasswordHash(hashedPassword)
	if err != nil {
		return nil, apperrors.BadParam.Wrap(err, "invalid password")
	}

	return &User{
		ID:       id,
		Name:     name,
		Email:    email,
		Password: password,
		IsActive: isActive,
	}, nil
}

type Password struct {
	value string
}

func NewPassword(pass string) (Password, error) {
	if len(pass) < passwordLengthMin {
		return Password{}, apperrors.BadParam.Wrap(ErrTooShort, "password must be at least 8 characters")
	}
	if !HasKigou(pass) {
		return Password{}, apperrors.BadParam.Wrap(ErrNoHasKigou, "password must not contain symbols (-_!?)")
	}
	if !HasHanSu(pass) {
		return Password{}, apperrors.BadParam.Wrap(ErrNoHasHanSu, "password must contain at least one number")
	}
	if !HasLowerEi(pass) {
		return Password{}, apperrors.BadParam.Wrap(ErrNoHasLowerEi, "password must contain at least one lowercase letter")
	}
	if !HasUpperEi(pass) {
		return Password{}, apperrors.BadParam.Wrap(ErrNoHasUpperEi, "password must contain at least one uppercase letter")
	}
	return Password{
		value: pass,
	}, nil
}

func NewPasswordHash(hash string) (Password, error) {
	return Password{
		value: hash,
	}, nil
}

func (p Password) Value() string {
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

const (
	UserStatusActive   = true
	UserStatusInactive = false
)

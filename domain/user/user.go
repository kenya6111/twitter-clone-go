package domain

import (
	"context"
	"twitter-clone-go/tutorial"

	"github.com/jackc/pgx/v5/pgtype"
)

type User struct {
	ID       int32
	Name     string
	Email    string
	Password string
	IsActive pgtype.Bool
}

type UserRepository interface {
	FindAll() ([]User, error)
	FindByEmail(email string) (*User, error)
	CountByEmail(email string) (int64, error)
	CreateUser(c context.Context, email string, hash []byte) (*User, error)
	CreateEmailVerifyToken(ctx context.Context, userId int32, token string, expiredAt pgtype.Timestamp) (*tutorial.EmailVerifyToken, error)
}

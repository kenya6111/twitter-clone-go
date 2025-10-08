package domain

import (
	"context"
	"time"

	ulid "github.com/oklog/ulid/v2"
)

type EmailVerifyToken struct {
	ID        string
	UserID    string
	Token     string
	ExpiresAt time.Time
	CreatedAt time.Time
}

type EmailVerifyTokenRepository interface {
	Save(ctx context.Context, userId string, token string) (*EmailVerifyToken, error)
	FindByToken(ctx context.Context, token string) (*EmailVerifyToken, error)
	DeleteByToken(ctx context.Context, token string) error
}

func ReconstructEmailVerifyToken(id string, userId string, token string, expiredAt time.Time, createdAt time.Time) (*EmailVerifyToken, error) {
	return &EmailVerifyToken{
		ID:        ulid.Make().String(),
		UserID:    userId,
		Token:     token,
		ExpiresAt: expiredAt,
		CreatedAt: createdAt,
	}, nil
}

func IsExpired(expiredAt time.Time) bool {
	return expiredAt.Before(time.Now())
}

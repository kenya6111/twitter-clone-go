package domain

import (
	"context"
	"time"
)

type EmailVerifyToken struct {
	ID        int32
	UserID    string
	Token     string
	ExpiresAt time.Time
	CreatedAt time.Time
}

type EmailVerifyTokenRepository interface {
	CreateEmailVerifyToken(ctx context.Context, userId string, token string) (*EmailVerifyToken, error)
	GetEmailVerifyToken(ctx context.Context, token string, expiredAt time.Time) (*EmailVerifyToken, error)
	DeleteEmailVerifyToken(ctx context.Context, token string) error
}

func IsExpired(expiredAt time.Time) bool {
	return expiredAt.After(time.Now())
}

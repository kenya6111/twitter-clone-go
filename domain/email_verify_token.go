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
	Save(ctx context.Context, userId string, token string) (*EmailVerifyToken, error)
	FindByToken(ctx context.Context, token string, expiredAt time.Time) (*EmailVerifyToken, error)
	DeleteByToken(ctx context.Context, token string) error
}

func IsExpired(expiredAt time.Time) bool {
	return expiredAt.Before(time.Now())
}

package domain

import (
	"time"
)

type EmailVerifyToken struct {
	ID        int32
	UserID    string
	Token     string
	ExpiresAt time.Time
	CreatedAt time.Time
}

// type EmailVerifyTokenRepository interface {
// 	CreateEmailVerifyToken(ctx context.Context, userId int32, token string, expiredAt time.Time) (*EmailVerifyToken, error)
// }

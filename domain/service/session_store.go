package service

import (
	"context"
)

// セッション名の定数
const (
	DefaultSessionName = "go-twitter-session"
)

// セッション管理のインターフェース
type SessionStore interface {
	Set(ctx context.Context, value interface{}) error
	Get(ctx context.Context, key string) (interface{}, error)
	Delete(ctx context.Context) error
	Clear(ctx context.Context) error
}

// セッションミドルウェアのインターフェース
type SessionMiddleware interface {
	GetMiddleware(sessionName string) interface{}
}

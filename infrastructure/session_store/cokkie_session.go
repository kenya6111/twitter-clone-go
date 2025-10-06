package session_store

import (
	"context"
	"fmt"
	"time"
	"twitter-clone-go/apperrors"

	"github.com/gin-gonic/gin"
	"github.com/oklog/ulid/v2"
	"github.com/redis/go-redis/v9"
)

// Redisベースのセッションストア
type SessionStore struct {
	Client *redis.Client
}

func NewSessionStore() *SessionStore {
	sessionStore := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})
	return &SessionStore{Client: sessionStore}
}

// // 格納される gin.Context のキー
type ginContextKey struct{}

var GinContextKey = ginContextKey{}

// // セッションに値を設定
func (s *SessionStore) Set(ctx context.Context, value interface{}) error {
	redisKey := ulid.Make().String()

	cmd := s.Client.Set(ctx, "session:"+redisKey, value, time.Minute*30)
	if err := cmd.Err(); err != nil {
		return err
	}
	ginCtx, ok := ctx.Value(GinContextKey).(*gin.Context)
	if !ok || ginCtx == nil {
		return apperrors.ErrInvalidContext
	}
	fmt.Println("redisKey:!!! " + redisKey)
	fmt.Println(value)
	ginCtx.SetCookie("sid", redisKey, 3600, "/", "localhost", false, false)
	return nil
}

// セッションから値を取得
func (s *SessionStore) Get(ctx context.Context, key string) (interface{}, error) {
	ginCtx, ok := ctx.Value(GinContextKey).(*gin.Context)
	if !ok || ginCtx == nil {
		return nil, apperrors.ErrInvalidContext
	}

	redisKey, _ := ginCtx.Cookie("sid")
	redisValue, err := s.Client.Get(ctx, "session:"+redisKey).Result()
	switch {
	case err == redis.Nil:
		fmt.Println("SessionKeyが登録されていません。")
		return nil, err
	case err != nil:
		fmt.Println("Session取得時にエラー発生：" + err.Error())
		return nil, err
	}
	return redisValue, nil
}

// セッションから値を削除
func (s *SessionStore) Delete(ctx context.Context) error {
	ginCtx, ok := ctx.Value(GinContextKey).(*gin.Context)
	if !ok || ginCtx == nil {
		return apperrors.ErrInvalidContext
	}
	redisKey, _ := ginCtx.Cookie("sid")
	fmt.Println("redisKey: !!!!!!!!!!!!!" + redisKey)
	s.Client.Del(ginCtx, "session:"+redisKey)
	ginCtx.SetCookie("sid", "", -1, "/", "localhost", false, false)
	return nil
}

// セッションをクリア
func (s *SessionStore) Clear(ctx context.Context) error {
	return nil
}

// GetStore は内部のstoreフィールドを取得（Gin用）
func (s *SessionStore) GetStore() *redis.Client {
	return s.Client
}

package repository

import (
	"context"
	"fmt"
	"log"
	"time"

	"twitter-clone-go/tutorial"
	db "twitter-clone-go/tutorial"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

var pool *pgxpool.Pool

func InitDB(p *pgxpool.Pool) {
	pool = p
}
func SelectUsers(c *gin.Context) (*[]db.User, error) {
	q := db.New(pool)
	resultSet, err := q.ListUsers(context.Background())
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	// c.JSON(http.StatusOK, resultSet)
	return &resultSet, nil
}

func GetUserByEmail(c *gin.Context, email string) (*db.User, error) {
	q := db.New(pool)
	resultSet, err := q.GetUserByEmail(context.Background(), email)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &resultSet, nil
}

func CreateUser(ctx context.Context, email string, hash []byte) (*tutorial.User, error) {
	q := db.New(pool)
	userInfo := db.CreateUserParams{
		Name:     email,
		Email:    email,
		Password: string(hash),
	}
	resultSet, err := q.CreateUser(context.Background(), userInfo)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return &resultSet, nil
}

func CreateEmailVerifyToken(ctx context.Context, userId int32, token string, expiredAt time.Time) (*tutorial.EmailVerifyToken, error) {
	q := db.New(pool)

	verifyInfo := db.CreateEmailVerifyTokenParams{
		UserID:    userId,
		Token:     token,
		ExpiresAt: expiredAt,
	}
	resultSet, err := q.CreateEmailVerifyToken(context.Background(), verifyInfo)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	fmt.Println(resultSet)
	return &resultSet, nil
}

func GetEmailVerifyToken(ctx context.Context, userId int32, token string, expiredAt time.Time) (*tutorial.EmailVerifyToken, error) {
	q := db.New(pool)
	verifyInfo := db.GetEmailVerifyTokenParams{
		Token:  token,
		UserID: userId,
	}
	resultSet, err := q.GetEmailVerifyToken(ctx, verifyInfo)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &resultSet, nil
}
func UpdateUser(ctx context.Context, userId int32) bool {
	q := db.New(pool)
	activateInfo := db.UpdateUserParams{
		ID:       userId,
		IsActive: true,
	}
	err := q.UpdateUser(ctx, activateInfo)
	if err != nil {
		fmt.Println(err)
		return false
	}
	fmt.Println("アクティベート完了")

	return true
}

func DeleteEmailVerifyToken(ctx context.Context, token string) bool {
	q := db.New(pool)
	err := q.DeleteEmailVerifyToken(ctx, token)
	if err != nil {
		fmt.Println(err)
		return false
	}
	fmt.Println("トークン削除完了")

	return true
}

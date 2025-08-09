package repository

import (
	"context"
	"fmt"
	"log"
	"time"

	"twitter-clone-go/tutorial"
	db "twitter-clone-go/tutorial"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

var pool *pgxpool.Pool

func InitDB(p *pgxpool.Pool) {
	pool = p
}
func SelectUsers(c *gin.Context) ([]db.User, error) {
	q := db.New(pool)
	resultSet, err := q.ListUsers(context.Background())
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return resultSet, nil
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
func CountUsersByEmail(c *gin.Context, email string) (int64, error) {
	q := db.New(pool)
	resultNum, err := q.CountUsersByEmail(context.Background(), email)
	if err != nil {
		log.Println(err)
		return 99, err
	}

	return resultNum, nil
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
		log.Println(err)
		return nil, err
	}
	return &resultSet, nil
}

func CreateEmailVerifyToken(ctx context.Context, userId int32, token string, expiredAt pgtype.Timestamp) (*tutorial.EmailVerifyToken, error) {
	q := db.New(pool)

	verifyInfo := db.CreateEmailVerifyTokenParams{
		UserID:    userId,
		Token:     token,
		ExpiresAt: expiredAt,
	}
	resultSet, err := q.CreateEmailVerifyToken(context.Background(), verifyInfo)
	if err != nil {
		log.Println(err)
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
		log.Println(err)
		return nil, err
	}

	return &resultSet, nil
}
func UpdateUser(ctx context.Context, userId int32) error {
	q := db.New(pool)
	activateInfo := db.UpdateUserParams{
		ID:       userId,
		IsActive: pgtype.Bool{Bool: true, Valid: true},
	}
	err := q.UpdateUser(ctx, activateInfo)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func DeleteEmailVerifyToken(ctx context.Context, token string) error {
	q := db.New(pool)
	err := q.DeleteEmailVerifyToken(ctx, token)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

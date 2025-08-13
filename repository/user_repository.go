package repository

import (
	"context"
	"fmt"
	"log"

	"twitter-clone-go/tutorial"
	db "twitter-clone-go/tutorial"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	pool *pgxpool.Pool
}

func NewUserRepository(pool *pgxpool.Pool) *UserRepository {
	return &UserRepository{pool: pool}
}

func (ur *UserRepository) SelectUsers() ([]db.User, error) {
	q := db.New(ur.pool)
	resultSet, err := q.ListUsers(context.Background())
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return resultSet, nil
}

func (ur *UserRepository) GetUserByEmail(c *gin.Context, email string) (*db.User, error) {
	q := db.New(ur.pool)
	resultSet, err := q.GetUserByEmail(context.Background(), email)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &resultSet, nil

}
func (ur *UserRepository) CountUsersByEmail(c *gin.Context, email string) (int64, error) {
	q := db.New(ur.pool)
	resultNum, err := q.CountUsersByEmail(context.Background(), email)
	if err != nil {
		log.Println(err)
		return 99, err
	}

	return resultNum, nil
}

func (ur *UserRepository) CreateUser(ctx context.Context, email string, hash []byte) (*tutorial.User, error) {
	q := db.New(ur.pool)
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

func (ur *UserRepository) CreateEmailVerifyToken(ctx context.Context, userId int32, token string, expiredAt pgtype.Timestamp) (*tutorial.EmailVerifyToken, error) {
	q := db.New(ur.pool)

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

package postgres

import (
	"context"
	"log"
	domain "twitter-clone-go/domain/user"
	"twitter-clone-go/tutorial"

	db "twitter-clone-go/tutorial"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	pool *pgxpool.Pool
}

func NewUserRepository(pool *pgxpool.Pool) *UserRepository {
	return &UserRepository{pool: pool}
}

func (ur *UserRepository) FindAll() ([]domain.User, error) {
	q := db.New(ur.pool)
	users, err := q.ListUsers(context.Background())
	if err != nil {
		log.Println(err)
		return nil, err
	}

	resultSets := make([]domain.User, 0, len(users))
	for _, u := range users {
		resultSets = append(resultSets, ur.adapt(&u))
	}
	return resultSets, nil
}

func (ur *UserRepository) FindByEmail(email string) (*domain.User, error) {
	q := db.New(ur.pool)
	user, err := q.GetUserByEmail(context.Background(), email)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	resultSet := ur.adapt(&user)
	return &resultSet, nil
}

func (ur *UserRepository) CountByEmail(email string) (int64, error) {
	q := db.New(ur.pool)
	resultNum, err := q.CountUsersByEmail(context.Background(), email)
	if err != nil {
		log.Println(err)
		return 99, err
	}
	return resultNum, nil
}

func (ur *UserRepository) CreateUser(email string, hash []byte) (*domain.User, error) {
	q := db.New(ur.pool)
	userInfo := db.CreateUserParams{
		Name:     email,
		Email:    email,
		Password: string(hash),
	}
	user, err := q.CreateUser(context.Background(), userInfo)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	resultSet := ur.adapt(&user)
	return &resultSet, nil
}

func (ur *UserRepository) CreateEmailVerifyToken(userId int32, token string, expiredAt pgtype.Timestamp) (*tutorial.EmailVerifyToken, error) {
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
	return &resultSet, nil
}

func (u *UserRepository) adapt(in *db.User) domain.User {
	return domain.User{
		ID:       in.ID,
		Name:     in.Name,
		Email:    in.Email,
		Password: in.Password,
		IsActive: in.IsActive,
	}
}

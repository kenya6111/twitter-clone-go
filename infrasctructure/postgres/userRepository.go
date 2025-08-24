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
	client *client
}

func NewUserRepository(pool *pgxpool.Pool) *UserRepository {
	return &UserRepository{&client{pool: pool}}
}

func (ur *UserRepository) FindAll() ([]domain.User, error) {
	q := db.New(ur.client.pool)
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
	q := db.New(ur.client.pool)
	user, err := q.GetUserByEmail(context.Background(), email)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	resultSet := ur.adapt(&user)
	return &resultSet, nil
}

func (ur *UserRepository) CountByEmail(email string) (int64, error) {
	q := db.New(ur.client.pool)
	resultNum, err := q.CountUsersByEmail(context.Background(), email)
	if err != nil {
		log.Println(err)
		return 99, err
	}
	return resultNum, nil
}

func (ur *UserRepository) CreateUser(c context.Context, email string, hash []byte) (*domain.User, error) {
	q := ur.client.Querier(c)
	userInfo := db.CreateUserParams{
		Name:     email,
		Email:    email,
		Password: string(hash),
	}
	user, err := q.CreateUser(c, userInfo)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	resultSet := ur.adapt(&user)
	return &resultSet, nil
}

func (ur *UserRepository) CreateEmailVerifyToken(ctx context.Context, userId int32, token string, expiredAt pgtype.Timestamp) (*tutorial.EmailVerifyToken, error) {
	q := ur.client.Querier(ctx)

	verifyInfo := db.CreateEmailVerifyTokenParams{
		UserID:    userId,
		Token:     token,
		ExpiresAt: expiredAt,
	}
	resultSet, err := q.CreateEmailVerifyToken(ctx, verifyInfo)
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

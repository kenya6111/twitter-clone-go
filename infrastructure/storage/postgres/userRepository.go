package postgres

import (
	"context"
	"log"
	"time"

	"twitter-clone-go/apperrors"
	domain "twitter-clone-go/domain/user"
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

	if len(users) == 0 {
		err = apperrors.NAData.Wrap(apperrors.ErrNoData, "no data")
		return nil, err
	}

	resultSets := make([]domain.User, 0, len(users))
	for _, u := range users {
		resultSets = append(resultSets, toUserDomain(&u))
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
	resultSet := toUserDomain(&user)
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
	resultSet := toUserDomain(&user)
	return &resultSet, nil
}

func (ur *UserRepository) CreateEmailVerifyToken(ctx context.Context, userId string, token string) (*domain.EmailVerifyToken, error) {
	q := ur.client.Querier(ctx)
	expiredAt := pgtype.Timestamp{}
	_ = expiredAt.Scan(time.Now().Add(24 * time.Hour * 7))

	args := db.CreateEmailVerifyTokenParams{
		UserID:    userId,
		Token:     token,
		ExpiresAt: expiredAt,
	}
	resultSet, err := q.CreateEmailVerifyToken(ctx, args)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	emailVerifyToken := toEmailVerifyTokenDomain(&resultSet)
	return &emailVerifyToken, nil
}

func toUserDomain(in *db.User) domain.User {
	p, _ := domain.NewPassword(in.Password)
	return domain.User{
		ID:       in.ID,
		Name:     in.Name,
		Email:    in.Email,
		Password: p,
		IsActive: in.IsActive.Bool,
	}
}
func toEmailVerifyTokenDomain(in *db.EmailVerifyToken) domain.EmailVerifyToken {
	return domain.EmailVerifyToken{
		ID:        in.ID,
		UserID:    in.UserID,
		Token:     in.Token,
		ExpiresAt: in.ExpiresAt.Time,
		CreatedAt: in.CreatedAt.Time,
	}
}

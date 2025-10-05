package postgres

import (
	"context"
	"log"
	"time"
	"twitter-clone-go/domain"
	"twitter-clone-go/tutorial"
	db "twitter-clone-go/tutorial"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

type EmailVerifyRepository struct {
	client *client
}

func NewEmailVerifyRepository(pool *pgxpool.Pool) *EmailVerifyRepository {
	return &EmailVerifyRepository{&client{pool: pool}}
}

func (ur *EmailVerifyRepository) Save(ctx context.Context, userId string, token string) (*domain.EmailVerifyToken, error) {
	q := ur.client.Querier(ctx)
	expiredAt := pgtype.Timestamp{}
	_ = expiredAt.Scan(time.Now().Add(24 * time.Hour * 7))

	args := tutorial.CreateEmailVerifyTokenParams{
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

func (ur *EmailVerifyRepository) FindByToken(ctx context.Context, token string) (*domain.EmailVerifyToken, error) {
	q := ur.client.Querier(ctx)
	args := tutorial.GetEmailVerifyTokenParams{
		Token: token,
	}
	resultSet, err := q.GetEmailVerifyToken(ctx, args)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	emailVerifyToken := toEmailVerifyTokenDomain(&resultSet)
	return &emailVerifyToken, nil
}

func (ur *EmailVerifyRepository) DeleteByToken(ctx context.Context, token string) error {
	q := ur.client.Querier(ctx)
	err := q.DeleteEmailVerifyToken(ctx, token)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
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

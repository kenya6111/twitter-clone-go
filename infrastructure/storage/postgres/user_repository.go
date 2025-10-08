package postgres

import (
	"context"
	"log"

	"twitter-clone-go/apperrors"
	"twitter-clone-go/domain"
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

func (ur *UserRepository) FindByEmail(ctx context.Context, email string) (*domain.User, error) {
	q := ur.client.Querier(ctx)
	user, err := q.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	resultSet := toUserDomainForHash(&user)
	return &resultSet, nil
}

func (ur *UserRepository) CountByEmail(ctx context.Context, email string) (int64, error) {
	q := ur.client.Querier(ctx)
	resultNum, err := q.CountUsersByEmail(ctx, email)
	if err != nil {
		return 1, err
	}
	return resultNum, nil
}

func (ur *UserRepository) CreateUser(ctx context.Context, name string, email string, hash string) (*domain.User, error) {
	q := ur.client.Querier(ctx)
	userInfo := db.CreateUserParams{
		Name:     name,
		Email:    email,
		Password: hash,
	}
	user, err := q.CreateUser(ctx, userInfo)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	resultSet := toUserDomain(&user)
	return &resultSet, nil
}

func (ur *UserRepository) ActivateUser(ctx context.Context, userId string) (*domain.User, error) {
	q := ur.client.Querier(ctx)
	activateInfo := db.UpdateUserParams{
		ID:       userId,
		IsActive: pgBool(true),
	}
	user, err := q.UpdateUser(ctx, activateInfo)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	resultSet := toUserDomain(&user)
	return &resultSet, nil
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
func toUserDomainForHash(in *db.User) domain.User {
	p, _ := domain.NewPasswordHash(in.Password)
	return domain.User{
		ID:       in.ID,
		Name:     in.Name,
		Email:    in.Email,
		Password: p,
		IsActive: in.IsActive.Bool,
	}
}

func pgBool(b bool) pgtype.Bool { return pgtype.Bool{Bool: b, Valid: true} }

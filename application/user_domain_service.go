package application

import (
	"context"
	"twitter-clone-go/apperrors"
	"twitter-clone-go/domain"
)

type UserDomainService struct {
	repo domain.UserRepository
}

func NewUserDomainService(r domain.UserRepository) *UserDomainService {
	return &UserDomainService{repo: r}
}

func (s *UserDomainService) IsDuplicatedEmail(ctx context.Context, email string) error {
	user, err := s.repo.CountByEmail(ctx, email)
	if err != nil {
		return apperrors.GetDataFailed.Wrap(apperrors.ErrNoData, "fail to get user by email")
	}

	if user > 0 {
		return apperrors.DuplicateData.Wrap(apperrors.ErrDuplicateData, "already exist user data")
	}
	return nil
}

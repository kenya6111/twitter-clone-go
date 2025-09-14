package postgres

import (
	"twitter-clone-go/apperrors"
)

type UserDomainService struct {
	repo *UserRepository
}

func NewUserDomainService(r *UserRepository) *UserDomainService {
	return &UserDomainService{repo: r}
}

func (s *UserDomainService) IsDuplicatedEmail(email string) error {
	user, err := s.repo.CountByEmail(email)
	if err != nil {
		return apperrors.GetDataFailed.Wrap(apperrors.ErrNoData, "fail to get user by email")
	}

	if user > 0 {
		return apperrors.DuplicateData.Wrap(apperrors.ErrDuplicateData, "already exist user data")
	}
	return nil
}

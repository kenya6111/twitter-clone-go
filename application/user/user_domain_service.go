package application

import (
	"twitter-clone-go/apperrors"
	domain "twitter-clone-go/domain/user"
)

type UserDomainService struct {
	repo domain.UserRepository
}

func NewUserDomainService(r domain.UserRepository) *UserDomainService {
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

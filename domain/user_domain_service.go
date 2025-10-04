package domain

import "context"

type UserDomainService interface {
	IsDuplicatedEmail(ctx context.Context, email string) error
}

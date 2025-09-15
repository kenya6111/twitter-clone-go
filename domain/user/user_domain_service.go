package domain

type UserDomainService interface {
	IsDuplicatedEmail(email string) error
}

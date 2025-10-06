package apperrors

import "errors"

type MyAppError struct {
	ErrCode
	Message string
	Err     error `json:"-"`
}

func (myErr *MyAppError) Error() string {
	return myErr.Err.Error()
}

func (myErr *MyAppError) Unwrap() error {
	return myErr.Err
}

var (
	ErrNoData                  = errors.New("get 0 record from db.Query")
	ErrDuplicateData           = errors.New("already exist user by email")
	ErrMismatchData            = errors.New("mismatch password and confirmPassword")
	ErrNoRequestParam          = errors.New("get no value from request")
	ErrEmailVerifyTokenExpired = errors.New("email verify token is already expired")
)

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

var ErrNoData = errors.New("get 0 record from db.Query")
var ErrDuplicateData = errors.New("already exist user by email")
var ErrMismatchData = errors.New("mismatch password and confirmPassword")
var ErrNoRequestParam = errors.New("get no value from request")

var ErrNoInvalidName = errors.New("name validation failed: missing uppercase letter")
var ErrNoInvalidEmail = errors.New("email validation failed: missing uppercase letter")

var ErrTooShort = errors.New("password validation failed: too short")
var ErrNoHasKigou = errors.New("password validation failed: contains symbol")
var ErrNoHasHanSu = errors.New("password validation failed: missing number")
var ErrNoHasLowerEi = errors.New("password validation failed: missing lowercase letter")
var ErrNoHasUpperEi = errors.New("password validation failed: missing uppercase letter")

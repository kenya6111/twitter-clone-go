package domain

import "errors"

var ErrNoInvalidName = errors.New("name validation failed: missing uppercase letter")
var ErrNoInvalidEmail = errors.New("email validation failed: missing uppercase letter")

var ErrTooShort = errors.New("password validation failed: too short")
var ErrNoHasKigou = errors.New("password validation failed: contains symbol")
var ErrNoHasHanSu = errors.New("password validation failed: missing number")
var ErrNoHasLowerEi = errors.New("password validation failed: missing lowercase letter")
var ErrNoHasUpperEi = errors.New("password validation failed: missing uppercase letter")

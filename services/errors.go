package services

import "errors"

var ErrNoData = errors.New("get 0 record from db.Query")
var ErrDuplicateData = errors.New("already exist user by email")
var ErrMismatchData = errors.New("mismatch password and confirmPassword")
var ErrNoRequestParam = errors.New("get no value from request")

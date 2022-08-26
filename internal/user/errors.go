package user

import "errors"

var (
	ErrNotFound         = errors.New("user not found")
	ErrWrongCredentials = errors.New("wrong credentials")
)

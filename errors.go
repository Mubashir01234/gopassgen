package gopassgen

import "errors"

var (
	errInvalidLength = errors.New("password length must be greater than 0")
	errMissingOption = errors.New("password must choose one of the options")
)

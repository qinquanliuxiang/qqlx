package reason

import "errors"

var (
	ErrPermission  = errors.New("permission denied")
	ErrHeaderEmpty = errors.New("auth in the request header is empty")
	ErrTokenMode   = errors.New("token mode error")
)

var (
	ErrUserNotFound    = errors.New("user does not exist")
	ErrInvalidPassword = errors.New("password is invalid")
)

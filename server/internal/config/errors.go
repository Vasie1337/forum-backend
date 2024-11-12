package config

import "errors"

var (
	ErrKeyAlreadyRedeemed = errors.New("key has already been redeemed")
	ErrKeyNotFound        = errors.New("key not found")
	ErrAdminNotFound      = errors.New("admin not found")
	ErrUserNotFound       = errors.New("user not found")
	ErrInvalidPassword    = errors.New("invalid password")
)

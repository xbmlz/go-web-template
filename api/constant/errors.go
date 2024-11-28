package constant

import "errors"

var (
	ErrInvalidRequest = errors.New("error.invalid_request")

	// auth errors
	ErrInvalidToken      = errors.New("error.auth.invalid_token")
	ErrTokenExpired      = errors.New("error.auth.token_expired")
	ErrUserNotFound      = errors.New("error.auth.user_not_found")
	ErrUsernameExists    = errors.New("error.auth.username_exists")
	ErrPasswordIncorrect = errors.New("error.auth.password_incorrect")

	// user errors
	ErrUserNotActive = errors.New("error.user.not_active")
)

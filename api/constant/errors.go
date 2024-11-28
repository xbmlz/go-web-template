package constant

import "errors"

var (
	ErrInvalidRequest = errors.New("error.invalid_request")
	ErrInvalidToken   = errors.New("error.invalid_token")

	// auth errors
	ErrUserNotFound   = errors.New("error.auth.user_not_found")
	ErrUsernameExists = errors.New("error.auth.username_exists")
)

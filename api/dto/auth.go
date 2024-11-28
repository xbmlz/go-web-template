package dto

import (
	"time"
)

type RegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
}

type RegisterResponse struct {
	ID uint `json:"id"`
}

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
type LoginResponse struct {
	Token       string    `json:"token"`
	ExpireAt    time.Time `json:"expire_at"`
	TokenPrefix string    `json:"token_prefix"` // "Bearer"
}

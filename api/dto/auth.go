package dto

import (
	"time"

	"github.com/xbmlz/go-web-template/api/model"
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

type UserInfoResponse struct {
	ID        uint         `json:"id"`
	Username  string       `json:"username"`
	Email     string       `json:"email"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
	Roles     []model.Role `json:"roles"`
}

type UserPermissionsResponse struct {
	Menus []model.Menu `json:"menus"`
}

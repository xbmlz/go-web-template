package dto

import "github.com/xbmlz/go-web-template/api/model"

type UserPageRequest = PageRequest

type UserPageResponse struct {
	PageResponse[*model.User]
}

type UserCreateRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required"`
}

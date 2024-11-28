package dto

import "github.com/xbmlz/go-web-template/api/model"

type UserPageRequest = PageRequest

type UserPageResponse struct {
	PageResponse[*model.User]
}

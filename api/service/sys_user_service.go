package service

import (
	"github.com/xbmlz/go-web-template/api/dto"
	"github.com/xbmlz/go-web-template/api/model"
	"github.com/xbmlz/go-web-template/api/query"
)

type SysUserService interface {
	GetUsers(req dto.UserPageRequest) (resp *dto.UserPageResponse, err error)
}

type sysUserService struct{}

func NewSysUserService() SysUserService {
	return &sysUserService{}
}

func (s *sysUserService) GetUsers(req dto.UserPageRequest) (resp *dto.UserPageResponse, err error) {
	q := query.User
	users, count, err := q.Preload(q.Roles).FindByPage((req.Page-1)*req.Size, req.Size)
	if err != nil {
		return nil, err
	}
	resp = &dto.UserPageResponse{
		PageResponse: dto.PageResponse[*model.User]{
			List:  users,
			Page:  req.Page,
			Size:  req.Size,
			Total: count,
		},
	}

	return resp, nil
}

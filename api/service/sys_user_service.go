package service

import (
	"github.com/jinzhu/copier"
	"github.com/xbmlz/go-web-template/api/dto"
	"github.com/xbmlz/go-web-template/api/model"
	"github.com/xbmlz/go-web-template/api/query"
)

type SysUserService interface {
	GetUserByID(userID uint) (resp *dto.UserInfoResponse, err error)
	GetUsers(req dto.UserPageRequest) (resp *dto.UserPageResponse, err error)
}

type sysUserService struct{}

func NewSysUserService() SysUserService {
	return &sysUserService{}
}

func (s *sysUserService) GetUserByID(userID uint) (resp *dto.UserInfoResponse, err error) {
	// get user by id
	q := query.User
	user, err := q.Preload(q.Roles).FindByID(userID)
	if err != nil {
		return resp, err
	}
	// convert user to dto
	resp = &dto.UserInfoResponse{}
	err = copier.Copy(resp, user)
	if err != nil {
		return resp, err
	}
	return resp, nil
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

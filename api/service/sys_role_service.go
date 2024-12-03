package service

import (
	"github.com/xbmlz/go-web-template/api/model"
)

type SysRoleService interface {
	FindMenusByID(id uint) (menus []model.Menu, err error)
}

type sysRoleService struct{}

func NewSysRoleService() SysRoleService {
	return &sysRoleService{}
}

func (s *sysRoleService) FindMenusByID(id uint) (menus []model.Menu, err error) {
	// q := query.Role

	// role, err := q.Preload(q.Menus).FindByID(id)
	// if err != nil {
	// 	return nil, err
	// }

	// menus = buildMenuTree(role.Menus)

	return menus, nil
}

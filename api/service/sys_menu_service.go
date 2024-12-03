package service

import (
	"github.com/xbmlz/go-web-template/api/model"
	"github.com/xbmlz/go-web-template/api/query"
)

type SysMenuService interface {
	GetMenuTreeByRoleIds(ids []uint) (menus []model.Menu, err error)
}

type sysMenuService struct{}

func NewSysMenuService() SysMenuService {
	return &sysMenuService{}
}

func (s *sysMenuService) GetMenuTreeByRoleIds(ids []uint) (menuTree []model.Menu, err error) {
	// get menus by role ids
	roleMenus, err := query.RoleMenu.Where(query.RoleMenu.RoleID.In(ids...)).Find()
	if err != nil {
		return nil, err
	}
	menuIDs := make([]uint, 0)
	for _, roleMenu := range roleMenus {
		menuIDs = append(menuIDs, roleMenu.MenuID)
	}
	// get menus by ids
	menus, err := query.Menu.Where(query.Menu.ID.In(menuIDs...)).Find()
	if err != nil {
		return nil, err
	}
	menuTree = model.BuildMenuTree(menus, 0)

	return menuTree, nil
}

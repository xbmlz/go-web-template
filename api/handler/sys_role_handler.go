package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/xbmlz/go-web-template/api/service"
)

type RoleHandler struct {
	BaseHandler
	roleService service.SysRoleService
	menuService service.SysMenuService
}

func NewRoleHandler(roleService service.SysRoleService, menuService service.SysMenuService) *RoleHandler {
	return &RoleHandler{roleService: roleService, menuService: menuService}
}

func (h *RoleHandler) RegisterRoutes(router *gin.RouterGroup) {

}

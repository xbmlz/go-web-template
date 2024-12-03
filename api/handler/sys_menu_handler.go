package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/xbmlz/go-web-template/api/service"
)

type MenuHandler struct {
	BaseHandler
	menuService service.SysMenuService
}

func NewMenuHandler(menuService service.SysMenuService) *MenuHandler {
	return &MenuHandler{menuService: menuService}
}

func (h *MenuHandler) RegisterRoutes(group *gin.RouterGroup) {}

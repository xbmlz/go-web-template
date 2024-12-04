package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xbmlz/go-web-template/api/constant"
	"github.com/xbmlz/go-web-template/api/dto"
	"github.com/xbmlz/go-web-template/api/service"
	"github.com/xbmlz/go-web-template/internal/middleware"
)

type AuthHandler struct {
	BaseHandler
	authService service.SysAuthService
	userService service.SysUserService
	menuService service.SysMenuService
}

func NewAuthHandler(authService service.SysAuthService, userService service.SysUserService, menuService service.SysMenuService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
		userService: userService,
		menuService: menuService,
	}
}

func (h *AuthHandler) RegisterRoutes(router *gin.RouterGroup) {
	group := router.Group("/auth")
	{
		group.POST("/register", h.register)
		group.POST("/login", h.login)

	}
	auth := group.Group("")
	auth.Use(middleware.AuthRequired())
	{
		auth.GET("/user", h.user)
		auth.GET("/permissions", h.permissions)
	}

}

// @tags auth
// @summary 注册
// @description 注册
// @accept application/json
// @produce application/json
// @param body body dto.RegisterRequest true "注册请求"
// @success 200  {object} dto.RegisterResponse "注册响应"
// @router /auth/register [post]
func (h *AuthHandler) register(c *gin.Context) {
	var req dto.RegisterRequest
	if !h.BindAndValidateJSON(c, &req) {
		return
	}
	resp, err := h.authService.Register(c, req)
	h.Response(c, err, resp)
}

// @tags auth
// @summary 登录
// @description 登录
// @accept application/json
// @produce application/json
// @param body body dto.LoginRequest true "登录请求"
// @success 200  {object} dto.LoginResponse "登录响应"
// @router /auth/login [post]
func (h *AuthHandler) login(c *gin.Context) {
	var req dto.LoginRequest
	if !h.BindAndValidateJSON(c, &req) {
		return
	}
	resp, err := h.authService.Login(c, req)
	h.Response(c, err, resp)
}

// @tags auth
// @summary 获取用户信息
// @description 获取用户信息
// @Security ApiKeyAuth
// @accept application/json
// @produce application/json
// @success 200  {object} dto.UserInfoResponse "用户信息响应"
// @router /auth/user [get]
func (h *AuthHandler) user(c *gin.Context) {
	userID := h.GetCurrentUserID(c)
	if userID == 0 {
		h.ErrorWithCode(c, constant.ErrUnauthorized, http.StatusUnauthorized)
		return
	}
	info, err := h.userService.GetUserByID(userID)
	h.Response(c, err, info)
}

// @tags auth
// @summary 获取用户权限
// @description 获取用户权限
// @Security ApiKeyAuth
// @accept application/json
// @produce application/json
// @success 200  {object} dto.UserPermissionsResponse "用户权限响应"
// @router /auth/permissions [get]
func (h *AuthHandler) permissions(c *gin.Context) {
	userID := h.GetCurrentUserID(c)
	if userID == 0 {
		h.ErrorWithCode(c, constant.ErrUnauthorized, http.StatusUnauthorized)
		return
	}

	user, err := h.userService.GetUserByID(userID)
	if err != nil {
		h.ErrorWithCode(c, constant.ErrUnauthorized, http.StatusUnauthorized)
		return
	}

	roleIDs := make([]uint, 0)
	for _, role := range user.Roles {
		roleIDs = append(roleIDs, role.ID)
	}

	// get user menus
	menus, err := h.menuService.GetMenuTreeByRoleIds(roleIDs)

	h.Response(c, err, dto.UserPermissionsResponse{
		Menus: menus,
	})
}

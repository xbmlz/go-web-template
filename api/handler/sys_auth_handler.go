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
}

func NewAuthHandler(authService service.SysAuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

func (h *AuthHandler) RegisterRoutes(router *gin.RouterGroup) {
	group := router.Group("/auth")
	{
		group.POST("/register", h.register)
		group.POST("/login", h.login)
		group.Use(middleware.AuthRequired()).GET("/user", h.user)
		group.Use(middleware.AuthRequired()).GET("/permissions", h.permissions)
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
	info, err := h.authService.GetUserInfo(userID)
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

	permissions, err := h.authService.GetUserPermissions(userID)
	h.Response(c, err, permissions)
}

package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/xbmlz/go-web-template/api/dto"
	"github.com/xbmlz/go-web-template/api/service"
)

type AuthHandler struct {
	BaseHandler
	authService service.AuthService
}

func NewAuthHandler(authService service.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

func (h *AuthHandler) Register(router *gin.RouterGroup) {
	group := router.Group("/auth")
	{
		group.POST("/register", h.register)
		group.POST("/login", h.login)
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

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
	if err := h.BindAndValidate(c, &req); err != nil {
		return
	}
	resp, err := h.authService.Register(c, req)
	h.Response(c, resp, err)
}

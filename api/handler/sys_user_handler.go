package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/xbmlz/go-web-template/api/dto"
	"github.com/xbmlz/go-web-template/api/service"
	"github.com/xbmlz/go-web-template/internal/middleware"
)

type SysUserHandler struct {
	BaseHandler
	userService service.SysUserService
}

func NewUserHandler(userService service.SysUserService) *SysUserHandler {
	return &SysUserHandler{userService: userService}
}

func (h *SysUserHandler) RegisterRoutes(router *gin.RouterGroup) {
	group := router.Group("/users").Use(middleware.AuthRequired())
	{
		group.GET("", h.getUsers)
		group.POST("", h.createUser)
	}
}

// @Summary Get users
// @Description Get users
// @Tags users
// @Produce json
// @Security ApiKeyAuth
// @Param page query int false "page number"
// @Param size query int false "page size"
// @Success 200 {object} dto.UserPageResponse
// @Router /users [get]
func (h *SysUserHandler) getUsers(c *gin.Context) {
	var req dto.UserPageRequest
	if !h.BindAndValidateQuery(c, &req) {
		return
	}
	users, err := h.userService.GetUsers(req)
	h.Response(c, err, users)
}

// @Summary Create user
// @Description Create user
// @Tags users
// @Produce json
// @Security ApiKeyAuth
// @Param user body dto.UserCreateRequest true "user"
// @Router /users [post]
func (h *SysUserHandler) createUser(c *gin.Context) {

}

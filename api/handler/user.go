package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/xbmlz/go-web-template/api/dto"
	"github.com/xbmlz/go-web-template/api/service"
	"github.com/xbmlz/go-web-template/internal/middleware"
)

type UserHandler struct {
	BaseHandler
	userService service.UserService
}

func NewUserHandler(userService service.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (h *UserHandler) Register(router *gin.RouterGroup) {
	group := router.Group("/users").Use(middleware.AuthRequired())
	{
		group.GET("", h.getUsers)
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
func (h *UserHandler) getUsers(c *gin.Context) {
	var req dto.UserPageRequest
	if !h.BindAndValidateQuery(c, &req) {
		return
	}
	users, err := h.userService.GetUsers(req)
	h.Response(c, err, users)
}

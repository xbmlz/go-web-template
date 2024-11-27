package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/xbmlz/go-web-template/api/dto"
	"github.com/xbmlz/go-web-template/api/service"
)

type AuthHandler struct {
	authService service.AuthService
}

func NewAuthHandler(authService service.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

func (a *AuthHandler) Setup(router *gin.RouterGroup) {
	group := router.Group("/auth")
	{
		group.POST("/login", a.register)
	}
}

func (a *AuthHandler) register(c *gin.Context) {
	// TODO: implement register logic
	var req dto.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"message": "invalid request",
		})
		return
	}

	a.authService.Register(c, req)

	c.JSON(200, gin.H{
		"message": "register success",
	})
}

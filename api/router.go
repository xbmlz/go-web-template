package api

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/xbmlz/go-web-template/api/handler"
	"github.com/xbmlz/go-web-template/api/service"
	"github.com/xbmlz/go-web-template/docs"
	"github.com/xbmlz/go-web-template/internal/config"
	"github.com/xbmlz/go-web-template/internal/logger"
	"github.com/xbmlz/go-web-template/internal/middleware"
	"github.com/xbmlz/go-web-template/internal/token"
)

func InitRouter(c *config.Config) *gin.Engine {
	if c.Server.IsDev() {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(cors.Default())
	r.Use(static.Serve("/", middleware.MustFS("")))

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "not found",
		})
	})

	// swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	docs.SwaggerInfo.BasePath = c.Server.BasePath

	// health check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, "ok")
	})

	// init token provider
	token.InitProvider(c.Token)

	logger.Infof("config: %v", c)

	root := r.Group(c.Server.BasePath)
	{
		// app server DI

		// service
		authService := service.NewSysAuthService(c)
		userService := service.NewSysUserService()
		roleService := service.NewSysRoleService()
		menuService := service.NewSysMenuService()

		// handler
		handler.NewAuthHandler(authService, userService, menuService).RegisterRoutes(root)
		handler.NewUserHandler(userService).RegisterRoutes(root)
		handler.NewRoleHandler(roleService, menuService).RegisterRoutes(root)
		handler.NewMenuHandler(menuService).RegisterRoutes(root)
	}

	return r
}

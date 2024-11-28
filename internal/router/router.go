package router

import (
	"net/http"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/xbmlz/go-web-template/api/handler"
	"github.com/xbmlz/go-web-template/api/service"
	"github.com/xbmlz/go-web-template/docs"
	"github.com/xbmlz/go-web-template/internal/config"
	"github.com/xbmlz/go-web-template/internal/middleware"
)

func Init(c *config.Config) *gin.Engine {
	if c.Server.IsDev() {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
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

	root := r.Group(c.Server.BasePath)
	{

		// service
		authService := service.NewAuthService()

		// handler
		handler.NewAuthHandler(authService).Register(root)
	}

	return r
}

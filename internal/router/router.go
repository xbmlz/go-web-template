package router

import (
	"net/http"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/xbmlz/go-web-template/api/handler"
	"github.com/xbmlz/go-web-template/api/service"
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

	root := r.Group("/api")
	{
		// ping
		root.GET("/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})

		// service
		authService := service.NewAuthService()

		handler.NewAuthHandler(authService).Setup(root)

	}

	return r
}

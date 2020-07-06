package core

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"skr-shop-cms-api/api"
	"skr-shop-cms-api/middleware"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.Use(middleware.Logger(), gin.Recovery())
	// 添加 Get 请求路由
	router.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, "hello gin")
	})
	router.Group("cms")
	brand := router.Group("brand")
	{
		brand.POST("/", middleware.Auth(), api.Brand)
	}

	return router
}

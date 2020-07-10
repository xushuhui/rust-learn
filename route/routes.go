package route

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"skr-shop-cms-api/api"
	"skr-shop-cms-api/middleware"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	//router.Use(middleware.Logger(), gin.Recovery())
	// 添加 Get 请求路由
	router.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, "hello gin")
	})
	authMiddle := middleware.AuthMiddleware()
	router.POST("/login", authMiddle.LoginHandler)
	router.Use(authMiddle.MiddlewareFunc())

	brand := router.Group("brand")

	brand.POST("/", middleware.Auth(), api.StoreBrand)

	brand.PUT("/:id/disable", middleware.Auth(), api.DisableBrand)
	brand.PUT("/:id", middleware.Auth(), api.UpdateBrand)
	brand.GET("/:id", middleware.Auth(), api.ShowBrand)
	brand.GET("/", middleware.Auth(), api.IndexBrand)
	brand.DELETE("/:id", middleware.Auth(), api.DestroyBrand)
	//
	//staff := router.Group("staff")
	//staff.POST("/", middleware.Auth().LoginHandler, api.StoreStaff)
	//staff.PUT("/:id", middleware.Auth().LoginHandler, api.UpdateStaff)
	//staff.GET("/:id", middleware.Auth().LoginHandler, api.ShowStaff)
	//staff.GET("/", middleware.Auth().LoginHandler, api.IndexStaff)
	return router
}

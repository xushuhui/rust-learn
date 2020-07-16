package route

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"skrshop-cms-api/api"
	"skrshop-cms-api/middleware"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.Use(middleware.ErrorHandle())
	//router.Use(middleware.Logger(), gin.Recovery())
	// 添加 Get 请求路由
	router.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, "hello gin")
	})

	router.POST("/login", api.Login)

	brand := router.Group("brand")
	{
		brand.POST("/", middleware.Auth(), api.StoreBrand)

		brand.PUT("/:id/disable", middleware.Auth(), api.DisableBrand)
		brand.PUT("/:id/enable", middleware.Auth(), api.EnableBrand)
		brand.PUT("/:id", middleware.Auth(), api.UpdateBrand)
		brand.GET("/:id", api.ShowBrand)
		brand.GET("/", middleware.Auth(), api.IndexBrand)
		brand.DELETE("/:id", middleware.Auth(), api.DestroyBrand)
	}

	//
	staff := router.Group("staff")
	{
		staff.POST("/", middleware.Auth(), api.StoreStaff)
		staff.PUT("/:id", middleware.Auth(), api.UpdateStaff)
		staff.GET("/:id", middleware.Auth(), api.ShowStaff)
		staff.GET("/", middleware.Auth(), api.IndexStaff)
	}

	return router
}

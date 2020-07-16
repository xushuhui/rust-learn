package api

import (
	"skrshop-cms-api/core"
	"skrshop-cms-api/model"
	"skrshop-cms-api/request"

	"github.com/gin-gonic/gin"
)

func StoreCategory(c *gin.Context) {
	var req request.Category
	if err := core.ParseRequest(c, &req); err != nil {
		return
	}

	return
}
func IndexCategory(c *gin.Context) {

	return
}
func ShowCategory(c *gin.Context) {
	id := c.Param("id")
	data, err := model.GetProductCategoryById(id)
	if err != nil {
		c.Error(err)
		return
	}
	core.SetData(c, data)
	return
}
func DestroyCategory(c *gin.Context) {
	return
}
func UpdateCategory(c *gin.Context) {
	return
}

package api

import (
	"github.com/gin-gonic/gin"
	"skrshop-cms-api/core"
	"skrshop-cms-api/model"
	"skrshop-cms-api/request"
)

func StoreSku(c *gin.Context) {
	var req request.Sku
	if err := core.ParseRequest(c, &req); err != nil {
		return
	}

	return
}
func IndexSku(c *gin.Context) {
	return
}
func ShowSku(c *gin.Context) {
	id := c.Param("id")
	data, err := model.GetProductSkuById(id)
	if err != nil {
		c.Error(err)
		return
	}
	core.SetData(c, data)
	return
}
func DestroySku(c *gin.Context) {
	return
}
func UpdateSku(c *gin.Context) {
	return
}
func IncrementStock(c *gin.Context) {
	return
}

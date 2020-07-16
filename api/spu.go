package api

import (
	"github.com/gin-gonic/gin"
	"skrshop-cms-api/core"
	"skrshop-cms-api/model"
	"skrshop-cms-api/request"
)

func StoreSpu(c *gin.Context) {
	var req request.Spu
	if err := core.ParseRequest(c, &req); err != nil {
		return
	}

	return
}
func IndexSpu(c *gin.Context) {
	return
}
func ShowSpu(c *gin.Context) {
	id := c.Param("id")
	data, err := model.GetProductSpuById(id)
	if err != nil {
		return
	}
	core.SetData(c, data)

	return
}
func DestroySpu(c *gin.Context) {
	return
}
func UpdateSpu(c *gin.Context) {
	return
}

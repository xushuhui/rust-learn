package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"skrshop-cms-api/core"
	"skrshop-cms-api/model"
	"skrshop-cms-api/request"
	"skrshop-cms-api/utils"
)

func StoreBrand(c *gin.Context) {
	var req request.Brand

	if err := core.ParseRequest(c, &req); err != nil {
		c.Error(err)
		return
	}
	uid := uint(1)
	brandModel := model.ProductBrands{Name: req.Name, Desc: req.Desc, LogoUrl: req.LogoUrl, CreateBy: uid}
	err := brandModel.Create()
	if err != nil {
		return
	}
	core.SuccessResp(c)
	return
}
func IndexBrand(c *gin.Context) {
	data := model.GetProductBrandsPage(1, 10)
	core.SetData(c, data)
	return
}
func ShowBrand(c *gin.Context) {
	id := c.Param("id")
	//data, err := model.GetProductBrandsById(id)
	data, err := model.GetProductBrandsOne("name=?", id)
	if err != nil {
		fmt.Println("ttt", err)
		c.Error(err)
		return
	}

	core.SetData(c, data)
	return
}
func DisableBrand(c *gin.Context) {
	id := c.Param("id")
	err := model.UpdateStatus(id, 0)
	if err != nil {
		c.Error(err)
		return
	}
	core.SuccessResp(c)
	return
}
func EnableBrand(c *gin.Context) {
	id := c.Param("id")
	err := model.UpdateStatus(id, 1)
	if err != nil {
		c.Error(err)
		return
	}
	core.SuccessResp(c)
	return
}
func DestroyBrand(c *gin.Context) {
	return
}
func UpdateBrand(c *gin.Context) {
	var req request.Brand

	if err := core.ParseRequest(c, &req); err != nil {
		c.Error(err)
		return
	}
	uid := uint(1)
	input := c.Param("id")
	id, err := utils.StringToUint(input)
	if err != nil {
		return
	}
	brandModel := &model.ProductBrands{Id: id, Name: req.Name, Desc: req.Desc, LogoUrl: req.LogoUrl, UpdateBy: uid}
	err = brandModel.Update()
	if err != nil {
		c.Error(err)
		return
	}
	core.SuccessResp(c)
	return
}

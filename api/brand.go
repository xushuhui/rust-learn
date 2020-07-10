package api

import (
	"github.com/gin-gonic/gin"
	"skr-shop-cms-api/core"
	"skr-shop-cms-api/model"
	"skr-shop-cms-api/request"
	"skr-shop-cms-api/utils"
)

func StoreBrand(context *gin.Context) {
	var req request.Brand

	if err := core.ParseRequest(context, &req); err != nil {
		return
	}
	uid := uint(1)
	brandModel := model.ProductBrands{Name: req.Name, Desc: req.Desc, LogoUrl: req.LogoUrl, CreateBy: uid}
	err := brandModel.Create()
	if err != nil {
		return
	}
	core.SuccessResp(context)
	return
}
func IndexBrand(context *gin.Context) {
	data := model.GetProductBrandsPage(1, 10)
	core.SetData(context, data)
	return
}
func ShowBrand(context *gin.Context) {
	id := context.Param("id")
	data, err := model.GetProductBrandsById(id)
	if err != nil {
		return
	}
	core.SetData(context, data)
	return
}
func DisableBrand(context *gin.Context) {
	id := context.Param("id")
	err := model.UpdateStatus(id, 0)
	if err != nil {
		return
	}
	core.SuccessResp(context)
	return
}
func EnableBrand(context *gin.Context) {
	id := context.Param("id")
	err := model.UpdateStatus(id, 1)
	if err != nil {
		return
	}
	core.SuccessResp(context)
	return
}
func DestroyBrand(context *gin.Context) {
	return
}
func UpdateBrand(context *gin.Context) {
	var req request.Brand

	if err := core.ParseRequest(context, &req); err != nil {
		return
	}
	uid := uint(1)
	input := context.Param("id")
	id, err := utils.StringToUint(input)
	if err != nil {
		return
	}
	brandModel := &model.ProductBrands{Id: id, Name: req.Name, Desc: req.Desc, LogoUrl: req.LogoUrl, UpdateBy: uid}
	brandModel.Update()
	core.SuccessResp(context)
	return
}
